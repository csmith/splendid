export default {
    phase: 'setup',
    turn: undefined,
    players: {},
    finalRound: false,
    tokens: {
        emerald: 7,
        diamond: 7,
        sapphire: 7,
        onyx: 7,
        ruby: 7,
        gold: 5,
    },
    decks: [
        [],
        [],
        [],
    ],
    cards: [
        Array(4).fill(undefined),
        Array(4).fill(undefined),
        Array(4).fill(undefined),
    ],
    nobles: [],
}