export default {
    name: 'add-player',

    perform: (state, {details}) => {
        state.players[details.id] = {
            details,
            reserved: [],
            nobles: [],
            order: undefined,
            points: 0,
            tokens: {
                emerald: 0,
                diamond: 0,
                sapphire: 0,
                onyx: 0,
                ruby: 0,
                gold: 0,
            },
            bonuses: {
                emerald: 0,
                diamond: 0,
                sapphire: 0,
                onyx: 0,
                ruby: 0,
            },
        };
    }
}