export default {
    name: 'change-player',

    perform: function(state, {playerId}) {
        state.turn = playerId;
    }
}