export default {
  name: "end-of-round-showdown",

  /**
   * @param hands {Array} Array of objects with playerId and hand
   */
  create: function (hands) {
    return {
      event: this.name,
      hands,
    };
  },

  perform: () => {
    // Do nothing: informational only
  },
};