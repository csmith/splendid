package com.chameth.splendid.games.diceclimbing.rules

import com.chameth.splendid.games.diceclimbing.State
import com.chameth.splendid.games.diceclimbing.Token

fun State.willWinIfStopping() = currentPlayerToken?.let { token ->
    board.count { it.tokens[token] == it.height-1 || it.tokens[Token.Black] == it.height-1 } == 3
} ?: false