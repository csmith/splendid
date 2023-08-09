export default {
  name: "round-over",

  perform: (state, { winningPlayerId }) => {
    state.players[winningPlayerId].points++;
  },
};
