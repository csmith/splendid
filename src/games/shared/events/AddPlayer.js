import _ from "lodash";

export default {
  name: "add-player",

  /**
   * @param details {Object} The engine-provided player details.
   */
  create: function (details) {
    return {
      event: this.name,
      details,
    };
  },

  perform: function (state, { details }) {
    state.players[details.id] = {
      ..._.cloneDeep(state.playerSkeleton),
      details,
    };
  },
};
