import AddPlayer from "../shared/events/AddPlayer.js";
import ChangePhase from "../shared/events/ChangePhase.js";
import ChangePlayer from "../shared/events/ChangePlayer.js";
import SetPlayerOrder from "../shared/events/SetPlayerOrder.js";
import CardNoOp from "./events/CardNoOp.js";
import DealCard from "./events/DealCard.js";
import DiscardCard from "./events/DiscardCard.js";
import EliminatePlayer from "./events/EliminatePlayer.js";
import GameOver from "./events/GameOver.js";
import GuardFailed from "./events/GuardFailed.js";
import HandRevealed from "./events/HandRevealed.js";
import ResetHands from "./events/ResetHands.js";
import RoundOver from "./events/RoundOver.js";
import SetProtection from "./events/SetProtection.js";
import Setup from "./events/Setup.js";
import SwapHands from "./events/SwapHands.js";

export default [
  AddPlayer,
  CardNoOp,
  ChangePhase,
  ChangePlayer,
  DealCard,
  DiscardCard,
  EliminatePlayer,
  GameOver,
  GuardFailed,
  HandRevealed,
  ResetHands,
  RoundOver,
  SetPlayerOrder,
  SetProtection,
  Setup,
  SwapHands,
];
