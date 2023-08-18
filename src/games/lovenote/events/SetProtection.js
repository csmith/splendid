export default {
  name: "set-protection",

  /**
   * @param playerId {string} The ID of the player whose protection is being changed.
   * @param isProtected {boolean} Whether the player is now protected or not.
   */
  create: function (playerId, isProtected) {
    return {
      event: this.name,
      playerId,
      isProtected,
    };
  },

  perform: (state, { playerId, isProtected }) => {
    state.players[playerId].protected = isProtected;
  },
};
