package com.chameth.splendid.games.klondike.ui

import com.chameth.splendid.games.klondike.Variant
import com.chameth.splendid.shared.playingcards.Card

sealed interface UiEvent {
    data object StockClicked : UiEvent
    data object WasteClicked : UiEvent
    data object RestartClicked : UiEvent
    data object AutoSolveClicked : UiEvent

    data class TableauClicked(val tableau: Int, val card: Card?) : UiEvent
    data class FoundationClicked(val foundation: Int) : UiEvent
    data class StartGameClicked(val variant: Variant) : UiEvent
}