export default {
  name: "add-player",

  perform: (state, { details }) => {
    state.players[details.id] = {
      details,
      points: 0,
      order: undefined,
      hand: [],
      discards: [],
      protected: false,
      eliminated: false,
    };
  },
};
