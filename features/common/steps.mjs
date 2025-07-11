import Engine from "../../src/common/engine.js";
import { newPlayer } from "../../src/common/player.js";
import games from "../../src/games.js";
import SetOptions from "../../src/games/shared/events/SetOptions.js";
import { Given, Then } from "@cucumber/cucumber";
import assert from "assert";
import _ from "lodash";

Given(/^the game being played was (.*?)$/, function (gameName) {
  this.engine = new Engine(games[gameName]);
  this.events = [];
  this.engine.onEvent((e) => this.events.push(e));
});

Given(/^the following players joined the game:$/, function (dataTable) {
  dataTable.rows().forEach((row) => {
    this.engine.perform("join", newPlayer(row[0]));
  });
});

Given(/^the game was started by (.*?)$/, function (playerName) {
  this.perform("start", playerName);
});

Given(/^it was (.*?)'s turn$/, function (playerName) {
  this.setState({
    ...this.engine.state,
    turn: this.playerDetails(playerName).id,
  });
});

Given(/^(.*?) had (\d+) points?$/, function (playerName, points) {
  const playerState = this.playerState(playerName);
  this.setState({
    ...this.engine.state,
    players: {
      ...this.engine.state.players,
      [playerState.details.id]: {
        ...playerState,
        points,
      },
    },
  });
});

Given(/^the turn order was:$/, function (dataTable) {
  const order = dataTable.raw().map((row) => row[0]);
  this.setState({
    ...this.engine.state,
    players: _.mapValues(this.engine.state.players, (v) => ({
      ...v,
      order: order.indexOf(v.details.name),
    })),
  });
});

Given(/^this was the final round$/, function () {
  this.setState({
    ...this.engine.state,
    finalRound: true,
  });
});

Given(/^the game phase was (.*?)$/, function (phase) {
  this.setState({
    ...this.engine.state,
    phase,
  });
});

Given(/^the game had the following options set:$/, function (dataTable) {
  const options = {};
  dataTable.rows().forEach(([key, value]) => {
    options[key] = value;
  });
  this.engine.applyEvent(SetOptions.create(options));
});

Then(/^an? "(.*?)" error will occur$/, function (message) {
  assert.equal(message, this.error?.message);
});

Then(/^the game phase will be (.*?)(?: still)?$/, function (phase) {
  assert.equal(this.engine.state.phase, phase);
});

Then(/^it will be (.*?)'s turn(?: still)?$/, function (playerName) {
  const details = this.playerDetails(playerName);
  assert.equal(this.engine.currentPlayer, details.id);
});

Then(/^(.*?) will have (\d+) points?$/, function (playerName, score) {
  const playerState = this.playerState(playerName);
  assert.equal(playerState.points, score);
});

Then(/^this will be the final round$/, function () {
  assert.ok(this.engine.state.finalRound);
});

Then(/^(.*?) will be eliminated$/, function (playerName) {
  const playerState = this.playerState(playerName);
  assert.ok(playerState.eliminated);
});

Then(/^(.*?) will not be eliminated$/, function (playerName) {
  const playerState = this.playerState(playerName);
  assert.ok(!playerState.eliminated);
});

Then(/^an? "(.*?)" event will be raised$/, function (eventName) {
  assert.ok(this.events.some((e) => e.event === eventName));
});

Then(/^an? "(.*?)" event will be not raised$/, function (eventName) {
  assert.ok(this.events.every((e) => e.event !== eventName));
});
