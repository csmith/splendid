export default {
    name: 'change-phase',

    available: () => false,

    perform: function(state, {phase}) {
        return {
            ...state,
            phase
        }
    }
}