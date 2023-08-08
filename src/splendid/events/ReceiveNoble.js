import _ from "lodash";

export default {
  name: "receive-noble",

  perform: (state, { playerId, noble }) => {
    _.remove(state.nobles, (n) => n.id === noble.id);
    state.players[playerId].nobles = [...state.players[playerId].nobles, noble];
  },
};
