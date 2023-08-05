export default {
    name: 'deal',

    available: () => false,

    perform: function (state, {level}) {
        const card = state.decks[level-1][0];

        return [
            {
                event: 'remove-card-from-deck',
                level,
            },
            {
                event: 'place-card',
                card,
            }
        ];
    }
}