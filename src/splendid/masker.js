export default function ({event, ...args}, playerId) {
    if (event === 'setup') {
        return {
            event,
            ...args,
            decks: args.decks.map(deck => deck.map(card => ({level: card.level}))),
        }
    } else if (event === 'reserve-card' && args.playerId !== playerId) {
        return {
            event,
            ...args,
            card: {
                level: args.card.level
            },
        }
    } else {
        return {event, ...args};
    }
}