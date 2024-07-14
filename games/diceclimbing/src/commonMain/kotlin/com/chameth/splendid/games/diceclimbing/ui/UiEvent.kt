package com.chameth.splendid.games.diceclimbing.ui

sealed interface UiEvent {
    data object JoinGameClicked : UiEvent
    data object StartGameClicked : UiEvent
    data class AdvanceColumnsClicked(val columns: List<Int>) : UiEvent
    data object GoBustClicked : UiEvent
    data object RollClicked : UiEvent
    data object StopClicked : UiEvent
}