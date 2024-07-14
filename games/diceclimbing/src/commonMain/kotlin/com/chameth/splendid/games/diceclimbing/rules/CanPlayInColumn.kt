package com.chameth.splendid.games.diceclimbing.rules

import com.chameth.splendid.games.diceclimbing.State

fun State.canPlayInColumn(roll: Int): Boolean {
    val column = columnForRoll(roll)

    return when {
        column.completed -> false
        column.hasBlackToken -> true
        blackTokenCount == 3 -> false
        else -> true
    }
}