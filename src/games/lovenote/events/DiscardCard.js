import _ from "lodash";

export default {
  name: "discard-card",

  perform: (state, { playerId, card }) => {
    state.players[playerId].hand = _.filter(state.players[playerId].hand, (c) => c.id !== card.id);
    state.players[playerId].discards.push(card);
  },
};
