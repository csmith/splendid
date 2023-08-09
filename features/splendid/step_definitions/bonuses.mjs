import { Given, Then, When } from "@cucumber/cucumber";
import assert from "assert";
import _ from "lodash";

Given(/^(.*?) had the following splendid bonuses:$/, function (playerName, dataTable) {
  let bonuses = _.mapValues(this.engine.state.tokens, () => 0);
  dataTable.hashes().forEach((row) => {
    bonuses[row.type] = parseInt(row.amount);
  });

  const playerState = this.playerState(playerName);

  this.setState({
    ...this.engine.state,
    players: {
      ...this.engine.state.players,
      [playerState.details.id]: {
        ...playerState,
        bonuses,
      },
    },
  });
});

When(/^(.*?) draws the following splendid tokens:$/, function (playerName, dataTable) {
  this.perform("take-tokens", playerName, {
    tokens: Object.fromEntries(dataTable.hashes().map((row) => [row.type, parseInt(row.amount)])),
  });
});

When(/^(.*?) discards the following splendid tokens:$/, function (playerName, dataTable) {
  this.perform("discard-tokens", playerName, {
    tokens: Object.fromEntries(dataTable.hashes().map((row) => [row.type, parseInt(row.amount)])),
  });
});

Then(/^(.*?) will have the following splendid bonuses:$/, function (playerName, dataTable) {
  const bonuses = this.playerState(playerName).bonuses;

  dataTable.hashes().forEach((row) => {
    assert.equal(bonuses[row.type], parseInt(row.amount));
  });
});
