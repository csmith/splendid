import {subtractObjects} from "../common/util.js";
import _ from "lodash";

export function costForCard(player, card) {
    return _.mapValues(subtractObjects(card.cost, player.bonuses), (v) => v < 0 ? 0 : v);
}

export function canAffordCard(player, card) {
    const cost = costForCard(player, card);
    const remaining = subtractObjects(player.tokens, cost);
    const missing = Object.values(remaining).filter((a) => a < 0).map((a) => a * -1)
    return _.sum(missing) <= player.tokens.gold;
}

export function canReceiveNoble(player, noble) {
    return _.every(noble.cost, (v, k) => player.bonuses[k] >= v);
}