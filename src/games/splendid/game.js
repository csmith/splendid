import ChangePhase from "../shared/events/ChangePhase.js";
import ChangePlayer from "../shared/events/ChangePlayer.js";
import SetPlayerOrder from "../shared/events/SetPlayerOrder.js";
import AddBonus from "./events/AddBonus.js";
import AddPlayer from "./events/AddPlayer.js";
import AddPoints from "./events/AddPoints.js";
import DiscardCard from "./events/DiscardCard.js";
import DiscardReserve from "./events/DiscardReserve.js";
import FinalRound from "./events/FinalRound.js";
import PlaceCard from "./events/PlaceCard.js";
import ReceiveNoble from "./events/ReceiveNoble.js";
import RemoveCardFromDeck from "./events/RemoveCardFromDeck.js";
import ReserveCard from "./events/ReserveCard.js";
import ReturnTokens from "./events/ReturnTokens.js";
import Setup from "./events/Setup.js";
import TakeTokens from "./events/TakeTokens.js";
import phases from "./phases.js";
import state from "./state.js";

export default {
  name: "Splendid",
  players: { min: 2, max: 4 },
  based_on: {
    game: "Splendor",
    creator: "Marc André",
    link: "https://www.spacecowboys.fr/splendor-english",
    purchase: "https://www.board-game.co.uk/product/splendor/",
  },
  description:
    "You play as a merchant, buying up developments to gain prestige and attract noble visitors. Do you build up a " +
    "strong foundation to support your future purchases, or do you rush to buy the most prestigious developments " +
    "sooner?",
  phases,
  state,
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
