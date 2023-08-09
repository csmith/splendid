import { countPlayers } from "../../../common/state.js";
import cards from "../data/cards.js";
import _ from "lodash";

const tokensToWin = {
  2: 7,
  3: 5,
  4: 4,
};

export default {
  name: "start-round",

  available: () => false,

  perform: function* (state) {
    const players = countPlayers(state);
    let deck = _.shuffle(cards).map((c) => ({ ...c, id: crypto.randomUUID() }));
    let spare = [deck.shift()];
    let unused = [];

    if (players === 2) {
      unused = deck.slice(0, 3);
      deck = deck.slice(3);
    }

    if (Object.values(state.players).some((p) => p.hand.length > 0 || p.discards.length > 0)) {
      yield {
        event: "reset-hands",
      };
    }

    yield {
      event: "setup",
      deck,
      unused,
      spare,
      tokensToWin: tokensToWin[players],
    };

    let turnOrder = _.map(_.sortBy(Object.values(state.players), "order"), "details.id");
    const start = turnOrder.indexOf(state.turn);
    turnOrder = turnOrder.slice(start).concat(turnOrder.slice(0, start));

    yield* turnOrder.map((player) => ({
      event: "deal-card",
      playerId: player,
      card: deck.shift(),
    }));

    yield {
      event: "deal-card",
      playerId: state.turn,
      card: deck.shift(),
    };
  },
};
