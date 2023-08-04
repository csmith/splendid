export function newPlayer(name) {
    return {
        id: crypto.randomUUID(),
        name,
    }
}