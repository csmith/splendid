import ReserveCard from "./events/ReserveCard.js";
import Setup from "./events/Setup.js";

export default function ({ event, ...args }, playerId) {
  if (event === Setup.name) {
    return {
      event,
      ...args,
      decks: args.decks.map((deck) =>
        deck.map((card) => ({
          id: card.id,
          level: card.level,
        })),
      ),
    };
  } else if (event === ReserveCard.name && args.playerId !== playerId) {
    return {
      event,
      ...args,
      card: {
        level: args.card.level,
        id: args.card.id,
      },
    };
  } else {
    return { event, ...args };
  }
}
