import { nextPlayer, remainingPlayers } from "../../../common/state.js";
import ChangePlayer from "../../shared/events/ChangePlayer.js";
import DealCard from "../events/DealCard.js";
import SetProtection from "../events/SetProtection.js";
import EndRound from "./EndRound.js";
import EndOfRoundShowdown from "../events/EndOfRoundShowdown.js";
import _ from "lodash";

export default {
  name: "end-turn",

  available: () => false,

  perform: function* (state) {
    const players = remainingPlayers(state);

    if (players.length === 1) {
      yield* EndRound.perform(state, { winningPlayerIds: [players[0].details.id] });
      return;
    }

    if (state.deck.length === 0) {
      yield EndOfRoundShowdown.create(
        players.map((player) => ({
          playerId: player.details.id,
          hand: player.hand,
        }))
      );

      const scores = Object.fromEntries(players.map((player) => [player.details.id, player.hand[0].closeness]));
      const bestScore = _.max(Object.values(scores));
      const bestPlayers = Object.entries(scores)
        .filter(([_, score]) => score === bestScore)
        .map(([playerId, _]) => playerId);
      
      const tiebreakBehaviour = state.options?.['tiebreak-behaviour'] || 'check-discards';
      let winners;
      
      switch (tiebreakBehaviour) {
        case 'check-discards':
          const discardSums = Object.fromEntries(
            bestPlayers.map((id) => [id, _.sum(state.players[id].discards.map((card) => card.closeness))]),
          );
          const bestDiscards = _.max(Object.values(discardSums));
          winners = bestPlayers.filter((id) => discardSums[id] === bestDiscards);
          break;
          
        case 'all-win':
          winners = bestPlayers;
          break;
          
        case 'no-winner':
          winners = bestPlayers.length > 1 ? [] : bestPlayers;
          break;
          
        case 'eliminate-ties':
          // Group players by score
          const scoreGroups = _.groupBy(Object.entries(scores), ([_, score]) => score);
          const sortedScores = Object.keys(scoreGroups).map(Number).sort((a, b) => b - a);
          
          // Find the highest score group with only one player
          winners = [];
          for (const score of sortedScores) {
            const playersAtScore = scoreGroups[score].map(([playerId, _]) => playerId);
            if (playersAtScore.length === 1) {
              winners = playersAtScore;
              break;
            }
          }
          break;
          
        default:
          winners = bestPlayers;
      }

      yield* EndRound.perform(state, { winningPlayerIds: winners });
      return;
    }

    const next = nextPlayer(state);

    yield ChangePlayer.create(next);

    if (state.players[next].protected) {
      yield SetProtection.create(next, false);
    }

    yield DealCard.create(next, state.deck[0]);
  },
};
