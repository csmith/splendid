import { findPlayerByName } from "../../src/common/state.js";
import { Before, After } from "@cucumber/cucumber";

Before(function () {
  this.setState = function (state) {
    this.engine.state = state;
  };

  this.playerState = function (name) {
    return findPlayerByName(this.engine.state, name);
  };

  this.playerDetails = function (name) {
    return this.playerState(name).details;
  };

  this.perform = function (name, playerName, args) {
    try {
      this.engine.perform(name, this.playerDetails(playerName), args);
      this.error = null;
    } catch (e) {
      this.error = e;
    }
  };

  this.error = null;
});

After(function (scenario) {
  if (scenario.result.status === 'FAILED' && this.engine) {
    console.log("=== DEBUG INFO FOR FAILED TEST ===");
    console.log("All events raised:", this.events?.map(e => e.event) || []);
    console.log("Engine state:", JSON.stringify(this.engine.state, null, 2));
    console.log("=====================================");
  }
});
