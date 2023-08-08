import AddBonus from "./events/AddBonus.js";
import AddPlayer from "./events/AddPlayer.js";
import AddPoints from "./events/AddPoints.js";
import ChangePhase from "./events/ChangePhase.js";
import ChangePlayer from "./events/ChangePlayer.js";
import DiscardCard from "./events/DiscardCard.js";
import DiscardReserve from "./events/DiscardReserve.js";
import FinalRound from "./events/FinalRound.js";
import PlaceCard from "./events/PlaceCard.js";
import ReceiveNoble from "./events/ReceiveNoble.js";
import RemoveCardFromDeck from "./events/RemoveCardFromDeck.js";
import ReserveCard from "./events/ReserveCard.js";
import ReturnTokens from "./events/ReturnTokens.js";
import SetPlayerOrder from "./events/SetPlayerOrder.js";
import Setup from "./events/Setup.js";
import TakeTokens from "./events/TakeTokens.js";
import masker from "./masker.js";
import phases from "./phases.js";
import state from "./state.js";

export default {
  name: "Splendid",
  phases,
  state,
  masker,
  events: [
    AddBonus,
    AddPlayer,
    AddPoints,
    ChangePhase,
    ChangePlayer,
    DiscardCard,
    DiscardReserve,
    FinalRound,
    PlaceCard,
    ReceiveNoble,
    RemoveCardFromDeck,
    ReserveCard,
    SetPlayerOrder,
    Setup,
    TakeTokens,
    ReturnTokens,
  ],
};
