export default {
  name: "play-card-handmaid",

  available: () => false,

  perform: function* (state, { playerData }) {
    yield {
      event: "set-protection",
      playerId: playerData.details.id,
      isProtected: true,
    };
  },
};
