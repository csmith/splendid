export default {
  name: "remove-card-from-deck",

  // TODO: This should be done as part of dealing/reserving the card.
  /**
   * @param level {number} The level of the deck (1, 2 or 3) to remove a card from.
   * @param playerId {?string} The player who caused the card to be removed
   * @param reason {string} The reason for removing the card.
   */
  create: function (level, playerId, reason) {
    return {
      event: this.name,
      level,
      playerId,
      reason,
    };
  },

  perform: (state, { level }) => {
    state.decks[level - 1].shift();
  },
};
