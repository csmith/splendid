export default {
    name: 'remove-card-from-deck',

    perform: (state, {level}) => {
        state.decks[level - 1].shift();
    },
}