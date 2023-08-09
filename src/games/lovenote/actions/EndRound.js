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

      yield {
        event: "change-phase",
        phase: "end",
      };
      return;
    }

    yield {
      event: "change-player",
      playerId: winningPlayerId,
    };

    yield {
      action: "start-round",
    };
  },
};
