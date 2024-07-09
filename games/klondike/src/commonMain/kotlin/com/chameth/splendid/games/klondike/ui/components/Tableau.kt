package com.chameth.splendid.games.klondike.ui.components

import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.runtime.Composable
import androidx.compose.runtime.key
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.chameth.splendid.games.klondike.ui.UiEvent
import com.chameth.splendid.games.klondike.ui.UiState
import com.chameth.splendid.games.klondike.ui.util.applySelectableState
import com.chameth.splendid.shared.ui.CardPlaceholder
import com.chameth.splendid.shared.ui.PlayingCard

@Composable
fun Tableau(state: UiState, modifier: Modifier = Modifier) {
    Row(
        modifier = modifier,
        horizontalArrangement = Arrangement.SpaceBetween
    ) {
        state.tableau.forEachIndexed { i, tableau ->
            key("tableau-$i") {
                Column(verticalArrangement = Arrangement.spacedBy((-120).dp)) {
                    tableau.forEach {
                        PlayingCard(
                            modifier = Modifier
                                .clickable { state.eventSink(UiEvent.TableauClicked(i, it.card)) }
                                .applySelectableState(it),
                            card = it.card,
                        )
                    }

                    if (tableau.isEmpty()) {
                        CardPlaceholder(
                            modifier = Modifier
                                .clickable { state.eventSink(UiEvent.TableauClicked(i, null)) }
                        )
                    }
                }
            }
        }
    }
}