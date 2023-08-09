export default function ({ event, ...args }, playerId) {
  if (event === "setup") {
    return {
      event,
      ...args,
      deck: args.deck.map((c) => ({ id: c.id })),
      unused: args.unused.map((c) => ({ id: c.id })),
    };
  } else if (event === "deal-card" && playerId !== args.playerId) {
    return {
      event,
      ...args,
      card: { id: args.card.id },
    };
  } else if (event === "hand-revealed" && playerId !== args.playerId && playerId !== args.handPlayerId) {
    return {
      event,
      ...args,
      hand: args.hand.map((c) => ({ id: c.id })),
    };
  } else if (event === "swap-hands") {
    return {
      event,
      ...args,
      players: args.players.map((p) => ({
        ...p,
        hand: p.id === playerId ? p.hand : p.hand.map((c) => ({ id: c.id })),
      })),
    };
  }
  return { event, ...args };
}
