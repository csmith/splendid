export function areAllProtected(state, ourPlayerId) {
  return Object.values(state.players).every((p) => p.details.id === ourPlayerId || p.protected);
}

export function mustDiscardCountess(state, playerId) {
  const hand = state.players[playerId].hand;
  return hand.some((c) => c.type === "Countess") && hand.some((c) => c.type === "Prince" || c.type === "King");
}
