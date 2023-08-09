import { findPlayerByName } from "../../src/common/state.js";
import { Before } from "@cucumber/cucumber";

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
});
