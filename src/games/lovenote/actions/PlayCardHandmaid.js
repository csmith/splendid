import SetProtection from "../events/SetProtection.js";

export default {
  name: "play-card-handmaid",

  available: () => false,

  perform: function* (state, { playerData }) {
    yield SetProtection.create(playerData.details.id, true);
  },
};
