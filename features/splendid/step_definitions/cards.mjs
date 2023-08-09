import { replaceNth } from "../../../src/common/util.js";
import { Given, Then, When } from "@cucumber/cucumber";
import assert from "assert";
import _ from "lodash";

Given(/^the following splendid cards were visible:$/, function (dataTable) {
  const cards = _.reverse(dataTable.raw().map((row) => row.map((card) => this.parseCard(card))));

  this.setState({
    ...this.engine.state,
    cards: cards,
  });
});

Given(/^the top splendid card of deck (\d+) was (.*?)$/, function (deck, card) {
  this.setState({
    ...this.engine.state,
    decks: replaceNth(this.engine.state.decks, deck - 1, () => [this.parseCard(card)]),
  });
});

Given(/^there were no splendid cards in deck (\d+)$/, function (deck) {
  this.setState({
    ...this.engine.state,
    decks: replaceNth(this.engine.state.decks, deck - 1, () => []),
  });
});

When(/^(.*?) buys the splendid card (.*?)$/, function (playerName, card) {
  this.perform("buy-card", playerName, { card: this.parseCard(card) });
});

Then(/^the splendid card in row (\d+) column (\d+) will be (.*?)$/, function (row, column, card) {
  const actual = this.engine.state.cards[3 - row][column - 1];
  const expected = this.parseCard(card);
  assert.equal(JSON.stringify(actual), JSON.stringify(expected));
});

Then(/^the size of splendid deck (\d+) will be (\d+)$/, function (deck, size) {
  assert.equal(this.engine.state.decks[deck - 1].length, size);
});
