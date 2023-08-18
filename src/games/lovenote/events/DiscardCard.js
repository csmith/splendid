import _ from "lodash";

export default {
  name: "discard-card",

  /**
   * @param playerId {string} The ID of the player that is discarding a card
   * @param card {object} The card that is being discarded
   */
  create: function (playerId, card) {
    return {
      event: this.name,
      playerId,
      card,
    };
  },

  perform: (state, { playerId, card }) => {
    state.players[playerId].hand = _.filter(state.players[playerId].hand, (c) => c.id !== card.id);
    state.players[playerId].discards.push(card);
  },
};
