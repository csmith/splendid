import _ from "lodash";

export default {
  name: "reserve-card",

  perform: (state, { playerId, card }) => {
    state.players[playerId].reserved = _.concat(state.players[playerId].reserved, card);
  },
};
