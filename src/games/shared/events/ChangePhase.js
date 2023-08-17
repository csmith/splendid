export default {
  name: "change-phase",

  /**
   * @param phase {string} The new phase of the game.
   */
  create: function (phase) {
    return {
      event: this.name,
      phase,
    };
  },

  perform: function (state, { phase }) {
    state.phase = phase;
  },
};
