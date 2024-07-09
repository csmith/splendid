package com.chameth.splendid.games.klondike.ui.model

import com.chameth.splendid.shared.playingcards.Card

data class Selection(
    val source: SelectionSource,
    val card: Card,
    val tableau: Int = -1
) {
    enum class SelectionSource {
        Tableau,
        Foundation,
        Waste,
    }
}