package com.chameth.splendid.shared.util

fun <X : Any> List<X>.replaceNth(n: Int, replacer: (X) -> X): List<X> {
    return this.mapIndexed { i, d ->
        if (i == n) {
            replacer(d)
        } else {
            d
        }
    }
}

fun <X : Any> List<X>.repeat(n: Int): List<List<X>> =
    (0 until n).map { this.toList() }