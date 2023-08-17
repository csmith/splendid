export default {
  name: "hand-revealed",

  mask: function (playerId, data) {
    if (playerId === data.playerId || playerId === data.handPlayerId) {
      return data;
    }

    return {
      ...data,
      card: { id: data.card.id },
    };
  },

  perform: () => {
    // Do nothing: informational only
  },
};
