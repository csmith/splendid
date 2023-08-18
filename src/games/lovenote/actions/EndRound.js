import ChangePhase from "../../shared/events/ChangePhase.js";
import ChangePlayer from "../../shared/events/ChangePlayer.js";
import GameOver from "../events/GameOver.js";
import RoundOver from "../events/RoundOver.js";

export default {
  name: "end-round",

  available: () => false,

  perform: function* (state, { winningPlayerId }) {
    yield RoundOver.create(winningPlayerId);

    if (state.players[winningPlayerId].points >= state.tokensToWin - 1) {
      yield GameOver.create(winningPlayerId);
      yield ChangePhase.create("end");
      return;
    }

    yield ChangePlayer.create(winningPlayerId);

    yield {
      action: "start-round",
    };
  },
};
