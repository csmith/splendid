package com.chameth.splendid.games.diceclimbing.rules

import com.chameth.splendid.games.diceclimbing.State

fun State.canPlayAny() = rollPermutations.any {
    canPlayInColumn(it.first) || canPlayInColumn(it.second)
}
