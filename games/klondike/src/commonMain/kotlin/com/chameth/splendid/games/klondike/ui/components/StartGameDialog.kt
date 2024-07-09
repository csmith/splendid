package com.chameth.splendid.games.klondike.ui.components

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.material3.Button
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.chameth.splendid.games.klondike.Variant
import com.chameth.splendid.games.klondike.ui.UiEvent
import com.chameth.splendid.games.klondike.ui.UiState
import com.chameth.splendid.shared.ui.components.Dialog
import dev.chrisbanes.haze.HazeState

@Composable
fun StartGameDialog(
    state: UiState,
    modifier: Modifier = Modifier,
    hazeState: HazeState = remember { HazeState() }
) {
    if (state.notStarted) {
        Dialog(modifier = modifier, hazeState = hazeState) {
            Column(
                horizontalAlignment = Alignment.CenterHorizontally,
                verticalArrangement = Arrangement.spacedBy(42.dp)
            ) {
                Text(
                    text = "Game setup",
                    style = MaterialTheme.typography.headlineMedium
                )

                Row(horizontalArrangement = Arrangement.spacedBy(16.dp)) {
                    Button(onClick = { state.eventSink(UiEvent.StartGameClicked(Variant.DrawOne)) }) {
                        Text("Draw one")
                    }

                    Button(onClick = { state.eventSink(UiEvent.StartGameClicked(Variant.DrawThree)) }) {
                        Text("Draw three")
                    }
                }
            }
        }
    }
}