import AddPlayer from "../shared/events/AddPlayer.js";
import ChangePhase from "../shared/events/ChangePhase.js";
import ChangePlayer from "../shared/events/ChangePlayer.js";
import SetOptions from "../shared/events/SetOptions.js";
import SetPlayerOrder from "../shared/events/SetPlayerOrder.js";
import AddBonus from "./events/AddBonus.js";
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

export default [
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
  ReturnTokens,
  SetOptions,
  SetPlayerOrder,
  Setup,
  TakeTokens,
];
