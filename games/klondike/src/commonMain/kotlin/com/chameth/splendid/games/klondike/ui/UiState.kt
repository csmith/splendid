package com.chameth.splendid.games.klondike.ui

data class UiState(
    val notStarted: Boolean = false,
    val gameOver: Boolean = false,
    val hasStock: Boolean = false,
    val canAutoSolve: Boolean = false,
    val waste: List<SelectableCard> = emptyList(),
    val foundations: List<SelectableCard?> = emptyList(),
    val tableau: List<List<SelectableCard>> = emptyList(),
    val eventSink: (UiEvent) -> Unit = {}
)
