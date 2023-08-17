import _ from "lodash";

export default {
  name: "reserve-card",

  /**
   * @param playerId {String} the player who is reserving the card
   * @param card {Object} the card to reserve
   */
  create: function (playerId, card) {
    return {
      event: this.name,
      playerId,
      card,
    };
  },

  perform: (state, { playerId, card }) => {
    state.players[playerId].reserved = _.concat(state.players[playerId].reserved, card);
  },
};
