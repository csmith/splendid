import ChangePhase from "../../shared/events/ChangePhase.js";
import ChangePlayer from "../../shared/events/ChangePlayer.js";

export default {
  name: "end-round",

  available: () => false,

  perform: function* (state, { winningPlayerId }) {
    yield {
      event: "round-over",
      winningPlayerId,
    };

    if (state.players[winningPlayerId].points >= state.tokensToWin - 1) {
      yield {
        event: "game-over",
        winningPlayerId,
      };

      yield ChangePhase.create("end");
      return;
    }

    yield ChangePlayer.create(winningPlayerId);

    yield {
      action: "start-round",
    };
  },
};
