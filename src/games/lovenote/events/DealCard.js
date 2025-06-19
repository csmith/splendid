export default {
  name: "deal-card",

  /**
   * @param playerId {string} The ID of the player that is being dealt the card
   * @param card {object} The card that is being dealt
   */
  create: function (playerId, card) {
    return {
      event: this.name,
      playerId,
      card,
    };
  },

  mask: function (playerId, data) {
    if (playerId === data.playerId) {
      return data;
    }

    return {
      ...data,
      card: { id: data.card.id },
    };
  },

  perform: (state, { playerId, card }) => {
    state.deck = state.deck.slice(1);
    state.players[playerId].hand.push(card);
  },
};
