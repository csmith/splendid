import { Given, Then, When } from "@cucumber/cucumber";
import assert from "assert";

Given(/^(.*?) had the following reserved splendid cards:$/, function (playerName, dataTable) {
  const cards = dataTable.raw().map((r) => this.parseCard(r[0]));
  const playerState = this.playerState(playerName);
  this.setState({
    ...this.engine.state,
    players: {
      ...this.engine.state.players,
      [playerState.details.id]: {
        ...playerState,
        reserved: cards,
      },
    },
  });
});

When(/^(.*?) reserves the splendid card (.*?)$/, function (playerName, card) {
  this.perform("reserve-card", playerName, { card: this.parseCard(card) });
});

When(/^(.*?) reserves a splendid card from deck (\d+)$/, function (playerName, deck) {
  this.perform("reserve-card-from-deck", playerName, { level: deck });
});

Then(/^(.*?) will have the following reserved splendid cards:$/, function (playerName, dataTable) {
  const actual = this.playerState(playerName).reserved;
  const expected = dataTable.raw().map((row) => this.parseCard(row[0]));
  assert.equal(JSON.stringify(actual), JSON.stringify(expected));
});
