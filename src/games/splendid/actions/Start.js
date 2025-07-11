// TODO: Maybe add a ready check for everyone.
import { countPlayers, findPlayer } from "../../../common/state.js";
import ChangePhase from "../../shared/events/ChangePhase.js";
import SetPlayerOrder from "../../shared/events/SetPlayerOrder.js";
import cards from "../data/cards.js";
import nobles from "../data/nobles.js";
import Setup from "../events/Setup.js";
import Deal from "./Deal.js";
import _ from "lodash";

const tokensToRemovePerPlayerCount = {
  2: 3,
  3: 2,
  4: 0,
};

export default {
  name: "start",

  available: function (state, { player }) {
    const count = countPlayers(state);
    const isPlayer = findPlayer(state, player);
    return isPlayer && count >= 2 && count <= 4;
  },

  perform: function* (state) {
    const players = countPlayers(state);
    const tokensToRemove = tokensToRemovePerPlayerCount[players];
    const cardsWithIds = cards.map((i) => ({ ...i, id: crypto.randomUUID() }));
    const decks = _.times(3, (level) => _.filter(cardsWithIds, (c) => c.level === level + 1));
    const turnOrder = _.shuffle(Object.keys(state.players));

    yield Setup.create(
      {
        emerald: 7 - tokensToRemove,
        diamond: 7 - tokensToRemove,
        sapphire: 7 - tokensToRemove,
        onyx: 7 - tokensToRemove,
        ruby: 7 - tokensToRemove,
        gold: 5,
      },
      _.take(_.shuffle(nobles), players + 1),
      decks.map((d) => _.shuffle(d)),
    );

    yield SetPlayerOrder.create(turnOrder);

    for (let i = 0; i < decks.length; i++) {
      for (let j = 0; j < 4; j++) {
        yield* Deal.perform(state, { level: i + 1 });
      }
    }

    yield ChangePhase.create("play");
  },
};
