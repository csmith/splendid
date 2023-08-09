import Join from "../shared/actions/Join.js";
import BuyCard from "./actions/BuyCard.js";
import Deal from "./actions/Deal.js";
import DiscardTokens from "./actions/DiscardTokens.js";
import EndTurn from "./actions/EndTurn.js";
import ReceiveNoble from "./actions/ReceiveNoble.js";
import ReserveCard from "./actions/ReserveCard.js";
import ReserveFromDeck from "./actions/ReserveFromDeck.js";
import Start from "./actions/Start.js";
import TakeTokens from "./actions/TakeTokens.js";

export default {
  setup: {
    actions: [Join, Start, Deal],
  },
  play: {
    actions: [TakeTokens, Deal, BuyCard, ReserveCard, ReserveFromDeck, EndTurn],
  },
  noble: {
    actions: [ReceiveNoble, EndTurn],
  },
  discard: {
    actions: [DiscardTokens, EndTurn],
  },
  end: {
    actions: [],
  },
};
