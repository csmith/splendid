package com.chameth.splendid.games.klondike.ui.components

import androidx.compose.foundation.clickable
import androidx.compose.runtime.Composable
import androidx.compose.runtime.key
import androidx.compose.ui.Modifier
import com.chameth.splendid.shared.ui.PlayingCard
import com.chameth.splendid.games.klondike.ui.UiEvent
import com.chameth.splendid.games.klondike.ui.UiState
import com.chameth.splendid.games.klondike.ui.util.applySelectableState

@Composable
fun Foundations(state: UiState) {
    state.foundations.forEachIndexed { i, foundation ->
        key("foundation-$i") {
            PlayingCard(
                modifier = Modifier
                    .clickable { state.eventSink(UiEvent.FoundationClicked(i)) }
                    .applySelectableState(foundation),
                card = foundation?.card,
            )
        }
    }
}