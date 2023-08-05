export default {
    name: 'add-bonus',

    perform: (state, {type}) => {
        state.players[state.turn].bonuses[type]++;
    }
}