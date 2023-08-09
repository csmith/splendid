import { Given, Then, When } from "@cucumber/cucumber";
import assert from "assert";

Given(/the following splendid nobles were available:$/, function (dataTable) {
  this.setState({
    ...this.engine.state,
    nobles: dataTable.raw().map((r) => this.parseSplendidNoble(r[0])),
  });
});

When(/^(.*?) receives the splendid noble (.*?)$/, function (playerName, noble) {
  this.perform("receive-noble", playerName, { noble: this.parseSplendidNoble(noble) });
});

Then(/^there will be (\d+) splendid nobles available?$/, function (count) {
  assert.equal(this.engine.state.nobles.length, count);
});
