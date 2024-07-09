package com.chameth.splendid.games.klondike.ui.components

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Row
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.PlayArrow
import androidx.compose.material.icons.filled.Refresh
import androidx.compose.material3.Button
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.chameth.splendid.games.klondike.ui.UiEvent
import com.chameth.splendid.games.klondike.ui.UiState

@Composable
fun ButtonBar(state: UiState, modifier: Modifier = Modifier) {
    Row(modifier = modifier, horizontalArrangement = Arrangement.spacedBy(8.dp)) {
        if (state.canAutoSolve) {
            Button(onClick = { state.eventSink(UiEvent.AutoSolveClicked) }) {
                Icon(imageVector =  Icons.Filled.PlayArrow, contentDescription = "")
                Text("Auto solve")
            }
        }

        Button(onClick = { state.eventSink(UiEvent.RestartClicked) }) {
            Icon(imageVector =  Icons.Filled.Refresh, contentDescription = "")
            Text("New game")
        }
    }
}
