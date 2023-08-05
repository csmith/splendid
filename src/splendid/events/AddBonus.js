export default {
    name: 'add-bonus',

    perform: (state, {playerId, type}) => {
        state.players[playerId].bonuses[type]++;
    }
}