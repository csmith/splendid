export default {
  name: "add-bonus",

  /**
   * @param playerId {string} The ID of the player who should receive the bonus
   * @param type {string} The bonus to add
   */
  create: function (playerId, type) {
    return {
      event: this.name,
      playerId,
      type,
    };
  },

  perform: (state, { playerId, type }) => {
    state.players[playerId].bonuses[type]++;
  },
};
