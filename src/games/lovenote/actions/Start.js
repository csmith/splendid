import { countPlayers, findPlayer } from "../../../common/state.js";
import ChangePhase from "../../shared/events/ChangePhase.js";
import SetPlayerOrder from "../../shared/events/SetPlayerOrder.js";
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

    yield SetPlayerOrder.create(turnOrder);
    yield ChangePhase.create("play");

    yield {
      action: "start-round",
    };
  },
};
