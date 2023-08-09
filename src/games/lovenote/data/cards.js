import _ from "lodash";

// TODO: These are the original card, the re-published version has more.
export default _.concat(
  _.times(5, () => ({ type: "Guard", closeness: 1 })),
  _.times(2, () => ({ type: "Priest", closeness: 2 })),
  _.times(2, () => ({ type: "Baron", closeness: 3 })),
  _.times(2, () => ({ type: "Handmaid", closeness: 4 })),
  _.times(2, () => ({ type: "Prince", closeness: 5 })),
  { type: "King", closeness: 6 },
  { type: "Countess", closeness: 7 },
  { type: "Princess", closeness: 8 },
);
