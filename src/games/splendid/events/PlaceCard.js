import _ from "lodash";

export default {
  name: "place-card",

  /**
   * @param card {Object} The card to be placed in the grid.
   */
  create: function (card) {
    return {
      event: this.name,
      card,
    };
  },

  perform: (state, { card }) => {
    const index = _.findIndex(state.cards[card.level - 1], (c) => c === undefined);
    state.cards[card.level - 1][index] = card;
  },
};
