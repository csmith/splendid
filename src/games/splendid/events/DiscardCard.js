export default {
  name: "discard-card",

  perform: (state, { card }) => {
    state.cards = state.cards.map((row) => row.map((c) => (c.id === card.id ? undefined : c)));
  },
};
