import _ from 'lodash';

export default {
    name: 'discard-reserve',

    perform: (state, {playerId, card}) => {
        if (state.players[playerId].reserved[0].bonus) {
            // We have the full state, remove the correct card
            state.players[playerId].reserved =
                _.filter(
                    state.players[playerId].reserved,
                    (c) => !_.isEqual(c, card)
                );
        } else {
            // We have a masked state, just pop any card off.
            // TODO: Technically we should match the level, but nothing uses it at the minute..
            state.players[playerId].reserved =
                state.players[playerId].reserved.slice(1);
        }
    }
}