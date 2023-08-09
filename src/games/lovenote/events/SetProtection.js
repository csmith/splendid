export default {
  name: "set-protection",

  perform: (state, { playerId, isProtected }) => {
    state.players[playerId].protected = isProtected;
  },
};
