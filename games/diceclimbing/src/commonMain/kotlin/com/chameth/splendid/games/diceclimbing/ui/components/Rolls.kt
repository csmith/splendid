package com.chameth.splendid.games.diceclimbing.ui.components

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.material3.Button
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.chameth.splendid.games.diceclimbing.ui.UiEvent
import com.chameth.splendid.games.diceclimbing.ui.UiState

@Composable
fun Rolls(
    uiState: UiState,
    modifier: Modifier = Modifier,
) {
    Column(modifier = modifier) {
        if (uiState.dice.isNotEmpty()) {
            Text(
                text = "Current rolls",
                style = MaterialTheme.typography.headlineSmall
            )
        }

        Row(horizontalArrangement = Arrangement.spacedBy(8.dp)) {
            uiState.dice.forEach {
                Die(it)
            }
        }

        uiState.options.forEach {
            Button(onClick = { uiState.eventSink(UiEvent.AdvanceColumnsClicked(it)) }) {
                Text(
                    text = if (it.size == 1) "Advance column ${it[0]}" else "Advance columns ${it[0]} and ${it[1]}",
                )
            }
        }

        if (uiState.goBust) {
            Button(onClick = { uiState.eventSink(UiEvent.GoBustClicked) }) {
                Text("No moves - go bust")
            }
        }

        if (uiState.canRoll) {
            Button(onClick = { uiState.eventSink(UiEvent.RollClicked) }) {
                Text("Roll the dice")
            }
        }

        if (uiState.canStop) {
            Button(onClick = { uiState.eventSink(UiEvent.StopClicked) }) {
                Text("Stop")
            }
        }
    }
}