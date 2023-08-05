export default {
    name: 'add-points',

    perform: (state, {points}) => {
        state.players[state.turn].points += points;
    }
}