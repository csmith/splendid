import { Then } from "@cucumber/cucumber";
import assert from "assert";

Then(/^the winner will need (\d+) love note points to win$/, function (count) {
  assert.equal(this.engine.state.tokensToWin, count);
});
