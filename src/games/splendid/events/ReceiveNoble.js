import _ from "lodash";

export default {
  name: "receive-noble",

  /**
   * @param playerId {string} The ID of the player receiving the noble
   * @param noble {Object} The noble being received
   */
  create: function (playerId, noble) {
    return {
      event: this.name,
      playerId,
      noble,
    };
  },

  perform: (state, { playerId, noble }) => {
    _.remove(state.nobles, (n) => n.id === noble.id);
    state.players[playerId].nobles = [...state.players[playerId].nobles, noble];
  },
};
