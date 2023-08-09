import { countPlayers, findPlayer } from "../../../common/state.js";
import cards from "../data/cards.js";
import _ from "lodash";

export default {
  name: "start",

  available: function (state, { player }) {
    const count = countPlayers(state);
    const isPlayer = findPlayer(state, player);
    return isPlayer && count >= 2 && count <= 4;
  },

  perform: function* (state) {
    const turnOrder = _.shuffle(Object.keys(state.players));

    yield {
      event: "set-player-order",
      order: turnOrder,
    };

    yield {
      event: "change-phase",
      phase: "play",
    };

    yield {
      action: "start-round",
    };
  },
};
