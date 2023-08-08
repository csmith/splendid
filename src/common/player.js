export function newPlayer(name, publicKey) {
  return {
    id: crypto.randomUUID(),
    name,
    publicKey,
  };
}
