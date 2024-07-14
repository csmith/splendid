package com.chameth.splendid.games.diceclimbing.ui.components

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.size
import androidx.compose.material3.Button
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.chameth.splendid.games.diceclimbing.ui.UiEvent
import com.chameth.splendid.games.diceclimbing.ui.UiState

@Composable
fun Players(state: UiState, modifier: Modifier = Modifier) {
    Column(modifier = modifier) {
        Text(
            text = "Players",
            style = MaterialTheme.typography.headlineSmall
        )

        if (state.players.isEmpty()) {
            Text("No-one has joined this game")
        }

        state.players.forEach {
            Row(verticalAlignment = Alignment.CenterVertically) {
                Token(
                    modifier = Modifier.size(16.dp),
                    type = it.first
                )

                Text(it.second)
            }
        }

        if (state.canJoin) {
            Button(onClick = { state.eventSink(UiEvent.JoinGameClicked) }) {
                Text("Join this game")
            }
        }

        if (state.canStart) {
            Button(onClick = { state.eventSink(UiEvent.StartGameClicked) }) {
                Text("Start this game")
            }
        }
    }
}