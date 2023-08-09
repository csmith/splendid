export default {
  name: "eliminate-player",

  perform: (state, { playerId }) => {
    state.players[playerId].eliminated = true;
  },
};
