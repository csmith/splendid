import { Given, Then, When } from "@cucumber/cucumber";
import assert from "assert";

Given(/^(.*?) is protected by a love note handmaid$/, function (playerName) {
  const playerState = this.playerState(playerName);

  this.setState({
    ...this.engine.state,
    players: {
      ...this.engine.state.players,
      [playerState.details.id]: {
        ...playerState,
        protected: true,
      },
    },
  });
});

Then(/^(.*?) will be protected by a love note handmaid$/, function (playerName) {
  const playerState = this.playerState(playerName);

  assert.ok(playerState.protected);
});

When(/^(.*?) ends their turn$/, function (playerName) {
  this.perform("end-turn", playerName);
});

Then(/^the round winners will be:$/, function (dataTable) {
  const roundOverEvent = this.events.find((e) => e.event === "round-over");
  assert.ok(roundOverEvent, "No round-over event was raised");
  
  const expectedWinners = dataTable.rows().map((row) => this.playerDetails(row[0]).id);
  const actualWinners = roundOverEvent.winningPlayerIds.sort();
  expectedWinners.sort();
  
  assert.deepEqual(actualWinners, expectedWinners);
});

Then(/^the game winners will be:$/, function (dataTable) {
  const gameOverEvent = this.events.find((e) => e.event === "game-over");
  assert.ok(gameOverEvent, "No game-over event was raised");
  
  const expectedWinners = dataTable.rows().map((row) => this.playerDetails(row[0]).id);
  const actualWinners = gameOverEvent.winningPlayerIds.sort();
  expectedWinners.sort();
  
  assert.deepEqual(actualWinners, expectedWinners);
});
