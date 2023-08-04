import _ from "lodash";

export function replaceNth(array, n, f) {
    return array.map((el, i) => i === n ? f(el) : el);
}

export function addObjects(obj1, obj2) {
    return _.mapValues(obj1, (v, k) => v + (obj2[k] || 0));
}

export function subtractObjects(obj1, obj2) {
    return _.mapValues(obj1, (v, k) => v - (obj2[k] || 0));
}

export function isUUID(str) {
    return /^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i.test(str);
}