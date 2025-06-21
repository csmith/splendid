import ChangePhase from "../../shared/events/ChangePhase.js";
import ChangePlayer from "../../shared/events/ChangePlayer.js";
import GameOver from "../events/GameOver.js";
import RoundOver from "../events/RoundOver.js";
import StartRound from "./StartRound.js";
import _ from "lodash";

export default {
  name: "end-round",

  available: () => false,

  perform: function* (state, { winningPlayerIds }) {
    yield RoundOver.create(winningPlayerIds);

    const gameWinners = winningPlayerIds.filter((id) => state.players[id].points >= state.tokensToWin - 1);
    if (gameWinners.length > 0) {
      yield GameOver.create(gameWinners);
      yield ChangePhase.create("end");
      return;
    }

    const nextPlayer = winningPlayerIds.length > 0 
      ? _.sample(winningPlayerIds)
      : _.sample(Object.keys(state.players));
    yield ChangePlayer.create(nextPlayer);

    yield* StartRound.perform(state);
  },
};
