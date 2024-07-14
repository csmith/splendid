package com.chameth.splendid.games.diceclimbing.rules

import com.chameth.splendid.games.diceclimbing.State

fun State.canPlayInColumns(first: Int, second: Int): Boolean {
    val firstColumn = columnForRoll(first)
    val secondColumn = columnForRoll(second)

    val tokensNeeded = (if (firstColumn.hasBlackToken) 0 else 1) +
            (if (secondColumn.hasBlackToken) 0 else 1)
    val tokensLeft = 3 - blackTokenCount

    return when {
        firstColumn.completed -> false
        secondColumn.completed -> false
        tokensLeft < tokensNeeded -> false
        else -> true
    }
}