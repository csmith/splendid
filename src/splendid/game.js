import phases from "./phases.js";
import state from "./state.js";
import masker from "./masker.js";
import AddPoints from "./events/AddPoints.js";
import DiscardReserve from "./events/DiscardReserve.js";
import ReturnTokens from "./events/ReturnTokens.js";
import AddBonus from "./events/AddBonus.js";
import DiscardCard from "./events/DiscardCard.js";

export default {
    name: 'Splendid',
    phases,
    state,
    masker,
    events: [
        AddBonus,
        AddPoints,
        DiscardCard,
        DiscardReserve,
        ReturnTokens,
    ]
}