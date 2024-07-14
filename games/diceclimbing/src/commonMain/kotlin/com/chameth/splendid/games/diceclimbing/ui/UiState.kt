package com.chameth.splendid.games.diceclimbing.ui

import com.chameth.splendid.games.diceclimbing.Column
import com.chameth.splendid.games.diceclimbing.Token
import com.chameth.splendid.shared.Die

data class UiState(
    val canJoin: Boolean = false,
    val canStart: Boolean = false,
    val players: List<Pair<Token, String>> = emptyList(),
    val board: List<Column> = emptyList(),
    val canRoll: Boolean = false,
    val canStop: Boolean = false,
    val dice: List<Die> = emptyList(),
    val options: List<List<Int>> = emptyList(),
    val goBust: Boolean = false,
    val eventSink: (UiEvent) -> Unit = {}
)
