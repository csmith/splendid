import state from "../shared/state.js";

export default {
  ...state,
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
