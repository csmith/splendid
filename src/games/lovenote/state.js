import state from "../shared/state.js";

export default {
  ...state,
  options: {
    ...state.options,
    "tiebreak-behaviour": "check-discards",
  },
  playerSkeleton: {
    ...state.playerSkeleton,
    points: 0,
    hand: [],
    discards: [],
    protected: false,
    eliminated: false,
  },
  deck: [],
  unused: [],
  spare: [],
  tokensToWin: 0,
};
