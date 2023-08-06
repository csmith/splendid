import _ from "lodash";

export function findPlayer(state, player) {
    const id = player.id || player;
    return state.players[id];
}

export function findPlayerByName(state, name) {
    return Object.values(state.players).find(player => player.details.name === name);
}

export function countPlayers(state) {
    return Object.keys(state.players).length;
}

export function nextPlayer(state) {
    const order = _.map(_.sortBy(Object.values(state.players), 'order'), 'details.id');
    const currentIndex = order.indexOf(state.turn);
    return order[(currentIndex + 1) % order.length];
}

export function isFirstPlayer(state, player) {
    const minOrder = _.min(_.map(Object.values(state.players), 'order'));
    return findPlayer(state, player).order === minOrder;
}

export function isLastPlayer(state, player) {
    const maxOrder = _.max(_.map(Object.values(state.players), 'order'));
    return findPlayer(state, player).order === maxOrder;
}