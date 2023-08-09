export default {
  name: "reset-hands",

  perform: (state) => {
    Object.keys(state.players).forEach((p) => {
      state.players[p].hand = [];
      state.players[p].discards = [];
      state.players[p].protected = false;
      state.players[p].eliminated = false;
    });
  },
};
