import { findPlayer } from "../../common/state.js";

export default {
  name: "join",

  available: function (state, { player }) {
    return findPlayer(state, player) === undefined;
  },

  perform: function* (state, { player }) {
    yield {
      event: "add-player",
      details: player,
    };
  },
};
