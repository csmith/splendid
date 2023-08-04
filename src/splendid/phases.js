import Join from "./actions/Join.js";
import Start from "./actions/Start.js";
import Deal from "./actions/Deal.js";
import ChangePhase from "./actions/ChangePhase.js";
import TakeTokens from "./actions/TakeTokens.js";
import BuyCard from "./actions/BuyCard.js";
import ReserveCard from "./actions/ReserveCard.js";
import EndTurn from "./actions/EndTurn.js";
import ReceiveNoble from "./actions/ReceiveNoble.js";
import ReserveFromDeck from "./actions/ReserveFromDeck.js";

export default {
    setup: {
        actions: [
            Join,
            Start,
            Deal,
            ChangePhase,
        ]
    },
    play: {
        actions: [
            TakeTokens,
            Deal,
            BuyCard,
            ReserveCard,
            ReserveFromDeck,
            EndTurn,
        ]
    },
    noble: {
        actions: [
            ReceiveNoble,
            EndTurn,
        ],
    },
    discard: {
        actions: [
            EndTurn,
        ],
    },
    end: {
        actions: [],
    },
};