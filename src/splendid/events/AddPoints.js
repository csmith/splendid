export default {
    name: 'add-points',

    perform: (state, {playerId, points}) => {
        state.players[playerId].points += points;
    }
}