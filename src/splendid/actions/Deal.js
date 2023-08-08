export default {
  name: "deal",

  available: () => false,

  perform: function* (state, { level }) {
    if (state.decks[level - 1].length === 0) {
      return;
    }

    const card = state.decks[level - 1][0];
    yield* [
      {
        event: "remove-card-from-deck",
        reason: "deal",
        level,
      },
      {
        event: "place-card",
        card,
      },
    ];
  },
};
