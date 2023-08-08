export default {
  name: "reserve-card-from-deck",

  available: function (state, { player }) {
    return player.id === state.turn && state.players[player.id].reserved.length < 3;
  },

  perform: function* (state, { level }) {
    if (level < 1 || level > 3) {
      throw new Error("Invalid level");
    }

    if (state.decks[level - 1].length === 0) {
      throw new Error("Deck is empty");
    }

    const card = state.decks[level - 1][0];
    const getsGold = state.tokens.gold > 0;

    yield {
      event: "remove-card-from-deck",
      playerId: state.turn,
      reason: "reserve",
      level,
    };

    yield {
      event: "reserve-card",
      playerId: state.turn,
      card,
    };

    if (getsGold) {
      yield {
        event: "take-tokens",
        playerId: state.turn,
        tokens: { gold: 1 },
      };
    }

    yield {
      action: "end-turn",
    };
  },
};
