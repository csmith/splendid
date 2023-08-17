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

  mask: function (playerId, data) {
    // The person reserving the card gets to see it...
    if (playerId === data.playerId) {
      return data;
    }

    return {
      ...data,
      card: {
        level: data.card.level,
        id: data.card.id,
      },
    };
  },

  perform: (state, { playerId, card }) => {
    state.players[playerId].reserved = _.concat(state.players[playerId].reserved, card);
  },
};
