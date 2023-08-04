import Join from "./actions/Join.js";
import Start from "./actions/Start.js";
import Deal from "./actions/Deal.js";
import ChangePhase from "./actions/ChangePhase.js";
import TakeTokens from "./actions/TakeTokens.js";
import BuyCard from "./actions/BuyCard.js";

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
        ]
    },
    noble: {
        actions: [],
    },
    discard: {
        actions: [],
    },
    end: {
        actions: [],
    },
};