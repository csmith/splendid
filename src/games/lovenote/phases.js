import Join from "../shared/actions/Join.js";
import DiscardCard from "./actions/DiscardCard.js";
import EliminatePlayer from "./actions/EliminatePlayer.js";
import EndRound from "./actions/EndRound.js";
import EndTurn from "./actions/EndTurn.js";
import PlayCard from "./actions/PlayCard.js";
import PlayCardBaron from "./actions/PlayCardBaron.js";
import PlayCardCountess from "./actions/PlayCardCountess.js";
import PlayCardGuard from "./actions/PlayCardGuard.js";
import PlayCardHandmaiden from "./actions/PlayCardHandmaid.js";
import PlayCardKing from "./actions/PlayCardKing.js";
import PlayCardPriest from "./actions/PlayCardPriest.js";
import PlayCardPrince from "./actions/PlayCardPrince.js";
import PlayCardPrincess from "./actions/PlayCardPrincess.js";
import Start from "./actions/Start.js";
import StartRound from "./actions/StartRound.js";

export default {
  setup: {
    actions: [Join, Start, StartRound],
  },
  play: {
    actions: [
      DiscardCard,
      EliminatePlayer,
      EndRound,
      EndTurn,
      PlayCard,
      PlayCardBaron,
      PlayCardCountess,
      PlayCardGuard,
      PlayCardHandmaiden,
      PlayCardKing,
      PlayCardPriest,
      PlayCardPrince,
      PlayCardPrincess,
      StartRound,
    ],
  },
  end: {
    actions: [],
  },
};
