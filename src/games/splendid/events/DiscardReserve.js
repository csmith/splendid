import _ from "lodash";

export default {
  name: "discard-reserve",

  // TODO: This shouldn't be an independent event, it should happen as part of playing a card.
  /**
   * @param playerId {string} The ID of the player whose reserve is being discarded
   * @param card {Object} The card that is being discarded
   * @returns {{event: string, card, playerId}}
   */
  create: function (playerId, card) {
    return {
      event: this.name,
      playerId,
      card,
    };
  },

  perform: (state, { playerId, card }) => {
    _.remove(state.players[playerId].reserved, (c) => c.id === card.id);
  },
};
