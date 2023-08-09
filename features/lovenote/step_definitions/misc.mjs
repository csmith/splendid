import { Given, Then } from "@cucumber/cucumber";
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
