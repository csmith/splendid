package com.chameth.splendid.games.klondike.ui.components

import androidx.compose.foundation.clickable
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import com.chameth.splendid.shared.ui.CardBack
import com.chameth.splendid.shared.ui.CardPlaceholder
import com.chameth.splendid.games.klondike.ui.UiEvent
import com.chameth.splendid.games.klondike.ui.UiState

@Composable
fun Stock(state: UiState, modifier: Modifier = Modifier) {
    if (state.hasStock) {
        CardBack(
            modifier = modifier
                .clickable { state.eventSink(UiEvent.StockClicked) }
        )
    } else {
        CardPlaceholder(
            modifier = modifier
                .clickable { state.eventSink(UiEvent.StockClicked) }
        )
    }
}