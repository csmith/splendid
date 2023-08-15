import _ from "lodash";

export default {
  name: "reserve-card",

  available: function (state, { player }) {
    return player.id === state.turn && state.players[player.id].reserved.length < 3;
  },

  perform: function* (state, { card }) {
    const index = _.findIndex(state.cards[card.level - 1], (c) => c.id === card.id);
    if (index === -1) {
      throw new Error("Card not found");
    }

    const getsGold = state.tokens.gold > 0;

    yield* [
      {
        event: "discard-card",
        reason: "reserve",
        playerId: state.turn,
        card,
      },
      {
        event: "reserve-card",
        playerId: state.turn,
        card,
      },
    ];

    if (getsGold) {
      yield {
        event: "take-tokens",
        playerId: state.turn,
        tokens: { gold: 1 },
      };
    }

    yield* [
      {
        action: "deal",
        level: card.level,
      },
      {
        action: "end-turn",
      },
    ];
  },
};
