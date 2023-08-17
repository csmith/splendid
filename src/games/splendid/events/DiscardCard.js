export default {
  name: "discard-card",

  // TODO: This shouldn't be an independent event, it should happen as part of taking a card.
  /**
   * @param card {Object} The card to discard
   * @param playerId {string} The player who caused the card to be "discarded"
   * @param reason {string} The reason the card was discarded
   */
  create: function (card, playerId, reason) {
    return {
      event: this.name,
      card,
      playerId,
      reason,
    };
  },

  perform: (state, { card }) => {
    state.cards = state.cards.map((row) => row.map((c) => (c.id === card.id ? undefined : c)));
  },
};
