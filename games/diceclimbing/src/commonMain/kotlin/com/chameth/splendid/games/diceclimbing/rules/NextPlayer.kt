package com.chameth.splendid.games.diceclimbing.rules

import com.chameth.splendid.games.diceclimbing.State

fun State.nextPlayer(): String =
    turnOrder[(turnOrder.indexOf(currentPlayer) + 1) % turnOrder.size]
