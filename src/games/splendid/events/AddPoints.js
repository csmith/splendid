export default {
  name: "add-points",

  /**
   * @param playerId {string} The ID of the player to give points to
   * @param points {number} The number of points to give
   */
  create: function (playerId, points) {
    return {
      event: this.name,
      playerId,
      points,
    };
  },

  perform: (state, { playerId, points }) => {
    state.players[playerId].points += points;
  },
};
