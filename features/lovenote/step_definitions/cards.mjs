import { Given, Then, When } from "@cucumber/cucumber";
import assert from "assert";

Given(/^(.*?) had the following love note cards in their hand:$/, function (playerName, dataTable) {
  const playerState = this.playerState(playerName);
  this.setState({
    ...this.engine.state,
    players: {
      ...this.engine.state.players,
      [playerState.details.id]: {
        ...playerState,
        hand: dataTable.raw().map((row) => this.parseLoveNoteCard(row[0])),
      },
    },
  });
});

Given(/^the next love note card in the deck is (.*?)$/, function (type) {
  this.engine.state.deck.unshift(this.parseLoveNoteCard(type));
});

When(
  /^(.*?) plays the love note card (.*?)(?: against (.*?)(?: with the guess "(.*?)")?)?$/,
  function (playerName, cardName, targetPlayerName, guessedType) {
    const playerState = this.playerState(playerName);
    const card = playerState.hand.find((c) => c.type === cardName);
    const targetPlayerId = this.playerState(targetPlayerName)?.details.id;
    this.perform("play-card", playerName, {
      cardId: card.id,
      targetPlayerId,
      guessedType,
    });
  },
);

Then(/^there will be (\d+) cards? in the love note unused deck$/, function (count) {
  assert.equal(this.engine.state.unused.length, count);
});

Then(/^there will be (\d+) cards? in the love note spare deck$/, function (count) {
  assert.equal(this.engine.state.spare.length, count);
});

Then(/^the current player will have (\d+) love note cards in their hand$/, function (count) {
  assert.equal(this.engine.state.players[this.engine.state.turn].hand.length, count);
});

Then(/^(.*?) will have the following love note cards in their hand:$/, function (playerName, dataTable) {
  const actual = this.playerState(playerName).hand.map((c) => c.type);
  const expected = dataTable.raw().map((row) => row[0]);
  assert.deepEqual(actual, expected);
});

Then(/^(.*?) will have the following love note cards in their discards:$/, function (playerName, dataTable) {
  const actual = this.playerState(playerName).discards.map((c) => c.type);
  const expected = dataTable.raw().map((row) => row[0]);
  assert.deepEqual(actual, expected);
});

Given(/^the love note deck is empty$/, function () {
  this.setState({
    ...this.engine.state,
    deck: [],
  });
});

Given(/^(.*?) had the following love note cards in their discard pile:$/, function (playerName, dataTable) {
  const playerState = this.playerState(playerName);
  this.setState({
    ...this.engine.state,
    players: {
      ...this.engine.state.players,
      [playerState.details.id]: {
        ...playerState,
        discards: dataTable.raw().map((row) => this.parseLoveNoteCard(row[0])),
      },
    },
  });
});
