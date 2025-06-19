import { nextPlayer, remainingPlayers } from "../../../common/state.js";
import ChangePlayer from "../../shared/events/ChangePlayer.js";
import DealCard from "../events/DealCard.js";
import SetProtection from "../events/SetProtection.js";
import EndRound from "./EndRound.js";
import _ from "lodash";

export default {
  name: "end-turn",

  available: () => false,

  perform: function* (state) {
    const players = remainingPlayers(state);

    if (players.length === 1) {
      yield* EndRound.perform(state, { winningPlayerId: players[0].details.id });
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

      yield* EndRound.perform(state, { winningPlayerId: winners[0] });
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
