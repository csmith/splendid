import state from "../shared/state.js";

export default {
  ...state,
  playerSkeleton: {
    ...state.playerSkeleton,
    reserved: [],
    nobles: [],
    points: 0,
    tokens: {
      emerald: 0,
      diamond: 0,
      sapphire: 0,
      onyx: 0,
      ruby: 0,
      gold: 0,
    },
    bonuses: {
      emerald: 0,
      diamond: 0,
      sapphire: 0,
      onyx: 0,
      ruby: 0,
    },
  },
  finalRound: false,
  tokens: {
    emerald: 7,
    diamond: 7,
    sapphire: 7,
    onyx: 7,
    ruby: 7,
    gold: 5,
  },
  decks: [[], [], []],
  cards: [Array(4).fill(undefined), Array(4).fill(undefined), Array(4).fill(undefined)],
  nobles: [],
};
