import { findPlayer } from "../../../common/state.js";
import AddPlayer from "../events/AddPlayer.js";

export default {
  name: "join",

  available: function (state, { player }) {
    return findPlayer(state, player) === undefined;
  },

  perform: function* (state, { player }) {
    yield AddPlayer.create(player);
  },
};
