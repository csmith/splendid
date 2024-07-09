package com.chameth.splendid.games.klondike.ui.components

import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Row
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.chameth.splendid.shared.ui.CardPlaceholder
import com.chameth.splendid.shared.ui.PlayingCard
import com.chameth.splendid.games.klondike.ui.UiEvent
import com.chameth.splendid.games.klondike.ui.UiState
import com.chameth.splendid.games.klondike.ui.util.applySelectableState

@Composable
fun Waste(state: UiState, modifier: Modifier = Modifier) {
    if (state.waste.isEmpty()) {
        CardPlaceholder(
            modifier = modifier
                .clickable { state.eventSink(UiEvent.WasteClicked) }
        )
    } else {
        Row(
            modifier = modifier,
            horizontalArrangement = Arrangement.spacedBy((-50).dp)
        ) {
            state.waste.forEach {
                PlayingCard(
                    modifier = Modifier
                        .clickable { state.eventSink(UiEvent.WasteClicked) }
                        .applySelectableState(it),
                    card = it.card
                )
            }
        }
    }
}