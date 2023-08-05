export default {
    name: 'change-player',

    perform: function(state, {player}) {
        state.turn = player;
    }
}