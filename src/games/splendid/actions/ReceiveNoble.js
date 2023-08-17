import AddPoints from "../events/AddPoints.js";
import ReceiveNoble from "../events/ReceiveNoble.js";
import { canReceiveNoble } from "../util.js";
import _ from "lodash";

export default {
  name: "receive-noble",

  available: function (state, { player }) {
    return player.id === state.turn;
  },

  perform: function* (state, { player, noble }) {
    const index = _.findIndex(state.nobles, (n) => n.id === noble.id);
    if (index === -1) {
      throw new Error("Noble not found");
    }

    if (!canReceiveNoble(state.players[player.id], noble)) {
      throw new Error("Not eligible for noble");
    }

    yield* [
      ReceiveNoble.create(state.turn, noble),
      AddPoints.create(state.turn, 3),
      {
        action: "end-turn",
      },
    ];
  },
};
