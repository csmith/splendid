export default {
  name: "add-player",

  /**
   * @param details {Object} The player details, as provided by the game engine
   */
  create: function (details) {
    return {
      event: this.name,
      details,
    };
  },

  perform: (state, { details }) => {
    state.players[details.id] = {
      details,
      points: 0,
      order: undefined,
      hand: [],
      discards: [],
      protected: false,
      eliminated: false,
    };
  },
};
