export default {
  name: "change-phase",

  perform: function (state, { phase }) {
    state.phase = phase;
  },
};
