import { Given, Then } from "@cucumber/cucumber";
import assert from "assert";
import _ from "lodash";

Given(/^the following splendid tokens were available:$/, function (dataTable) {
  let tokens = _.mapValues(this.engine.state.tokens, () => 0);
  dataTable.hashes().forEach((row) => {
    tokens[row.type] = parseInt(row.amount);
  });

  this.setState({
    ...this.engine.state,
    tokens: tokens,
  });
});

Given(/^(.*?) had the following splendid tokens:$/, function (playerName, dataTable) {
  let tokens = _.mapValues(this.engine.state.tokens, () => 0);
  dataTable.hashes().forEach((row) => {
    tokens[row.type] = parseInt(row.amount);
  });

  const playerState = this.playerState(playerName);

  this.setState({
    ...this.engine.state,
    players: {
      ...this.engine.state.players,
      [playerState.details.id]: {
        ...playerState,
        tokens: tokens,
      },
    },
  });
});

Then(/^(.*?) will have the following splendid tokens:$/, function (playerName, dataTable) {
  const tokens = this.playerState(playerName).tokens;

  dataTable.hashes().forEach((row) => {
    assert.equal(tokens[row.type], parseInt(row.amount));
  });
});

Then(/^the following splendid tokens will be available:$/, function (dataTable) {
  const tokens = this.engine.state.tokens;

  dataTable.hashes().forEach((row) => {
    assert.equal(tokens[row.type], parseInt(row.amount), row.type);
  });
});
