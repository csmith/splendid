export default {
  name: "swap-hands",

  perform: (state, { players }) => {
    state.players[players[0].id].hand = players[0].hand;
    state.players[players[1].id].hand = players[1].hand;
  },
};
