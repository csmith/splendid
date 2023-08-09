import { nextPlayer, remainingPlayers } from "../../../common/state.js";

export default {
  name: "end-turn",

  available: () => false,

  perform: function* (state) {
    const players = remainingPlayers(state);

    if (players.length === 1) {
      yield {
        action: "end-round",
        args: {
          winningPlayerId: players[0].details.id,
        },
      };
      return;
    }

    if (state.deck.length === 0) {
      yield {
        event: "end-of-round-showdown",
        hands: players.map((player) => ({
          playerId: player.details.id,
          hand: player.hand,
        })),
      };

      const scores = Object.fromEntries(players.map((player) => [player.details.id, player.hand[0].closeness]));
      const bestScore = _.max(Object.values(scores));
      const bestPlayers = Object.entries(scores)
        .filter(([_, score]) => score === bestScore)
        .map(([playerId, _]) => playerId);
      const discardSums = Object.fromEntries(
        bestPlayers.map((id) => [id, _.sum(state.players[id].discards.map((card) => card.closeness))]),
      );
      const bestDiscards = _.max(Object.values(discardSums));
      const winners = bestPlayers.filter((id) => discardSums[id] === bestDiscards);

      if (winners.length > 1) {
        throw new Error("The round is fully tied...");
      }

      yield {
        action: "end-round",
        args: {
          winningPlayerId: winners[0],
        },
      };
      return;
    }

    const next = nextPlayer(state);

    yield {
      event: "change-player",
      playerId: next,
    };

    if (state.players[next].protected) {
      yield {
        event: "set-protection",
        playerId: next,
        isProtected: false,
      };
    }

    yield {
      event: "deal-card",
      playerId: next,
      card: state.deck[0],
    };
  },
};
