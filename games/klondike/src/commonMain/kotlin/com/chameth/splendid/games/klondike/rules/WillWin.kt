package com.chameth.splendid.games.klondike.rules

import com.chameth.splendid.games.klondike.State

fun State.willWin(): Boolean = foundations.sumOf { it.count() } == 51