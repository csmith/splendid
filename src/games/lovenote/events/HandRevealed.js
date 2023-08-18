export default {
  name: "hand-revealed",

  /**
   * @param playerId {string} The ID of the player that the hand is being revealed to.
   * @param handPlayerId {string} The ID of the player whose hand is being revealed.
   * @param hand {Object[]} The hand being revealed.
   */
  create: function (playerId, handPlayerId, hand) {
    return {
      event: this.name,
      playerId,
      handPlayerId,
      hand,
    };
  },

  mask: function (playerId, data) {
    if (playerId === data.playerId || playerId === data.handPlayerId) {
      return data;
    }

    return {
      ...data,
      hand: data.hand.map((c) => ({ id: c.id })),
    };
  },

  perform: () => {
    // Do nothing: informational only
  },
};
