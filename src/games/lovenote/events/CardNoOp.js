export default {
  name: "card-no-op",

  /**
   * @param playerId {string} The ID of the player that played the card
   */
  create: function (playerId) {
    return {
      event: this.name,
      playerId,
    };
  },

  perform: () => {
    // Do nothing: informational only
  },
};
