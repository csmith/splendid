export default {
  name: "set-options",

  /**
   * @param options {Object} Map of option keys to values
   */
  create: function (options) {
    return {
      event: this.name,
      options,
    };
  },

  perform: function (state, { options }) {
    state.options = {
      ...state.options,
      ...options
    };
  },
};