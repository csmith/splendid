export default {
    name: 'setup',

    perform: (state, {tokens, nobles, decks}) => {
        state.tokens = tokens;
        state.nobles = nobles;
        state.decks = decks;
    }
}