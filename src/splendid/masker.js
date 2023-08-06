export default function ({event, ...args}, playerId) {
    if (event === 'setup') {
        return {
            event,
            ...args,
            decks: args.decks.map(deck => deck.map(card => ({
                id: card.id,
                level: card.level
            }))),
        }
    } else if (event === 'reserve-card' && args.playerId !== playerId) {
        return {
            event,
            ...args,
            card: {
                level: args.card.level,
                id: args.card.id,
            },
        }
    } else {
        return {event, ...args};
    }
}