import PlaceCard from "../events/PlaceCard.js";
import RemoveCardFromDeck from "../events/RemoveCardFromDeck.js";

export default {
  name: "deal",

  available: () => false,

  perform: function* (state, { level }) {
    if (state.decks[level - 1].length === 0) {
      return;
    }

    const card = state.decks[level - 1][0];
    yield RemoveCardFromDeck.create(level, undefined, "deal");
    yield PlaceCard.create(card);
  },
};
