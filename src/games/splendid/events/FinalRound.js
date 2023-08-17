export default {
  name: "final-round",

  create: function () {
    return {
      event: this.name,
    };
  },

  perform: function (state) {
    state.finalRound = true;
  },
};
