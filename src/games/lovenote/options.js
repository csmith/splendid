export default [
  {
    name: "Tiebreak behaviour",
    key: "tiebreak-behaviour",
    description: "How to handle players with equal cards at the end of a round",
    default: "check-discards",
    options: [
      {
        name: "Check discards",
        description: "Tied players add the score of all their discards, highest wins",
        value: "check-discards"
      },
      {
        name: "All win",
        description: "Tied players all get a token",
        value: "all-win"
      },
      {
        name: "No winner",
        description: "No-one wins if the highest scoring cards are tied",
        value: "no-winner"
      },
      {
        name: "Eliminate ties",
        description: "Anyone with a tie is eliminated, a player with a lower card could still win",
        value: "eliminate-ties"
      }
    ]
  }
];