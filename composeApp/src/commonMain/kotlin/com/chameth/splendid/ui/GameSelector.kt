package com.chameth.splendid.ui

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.material3.Button
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.material3.TextField
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp

@Composable
fun GameSelector(
    state: UiState.NoGame,
    modifier: Modifier = Modifier,
) {
    Box(modifier = modifier) {
        Column(
            modifier = Modifier.align(Alignment.Center),
            horizontalAlignment = Alignment.CenterHorizontally,
            verticalArrangement = Arrangement.spacedBy(48.dp)
        ) {
            Text(
                text = "Welcome to Splendid!",
                style = MaterialTheme.typography.headlineMedium
            )

            Text(
                text = "Start a new game",
                style = MaterialTheme.typography.bodyMedium
            )

            state.availableTypes.forEach {
                Button(onClick = {
                    state.eventSink(UiEvent.NoGame.CreateGameTapped(it))
                }) {
                    Text(text = it)
                }
            }

            Text(
                text = "Or join an existing game",
                style = MaterialTheme.typography.bodyMedium
            )

            TextField(
                value = state.gameId,
                onValueChange = { state.eventSink(UiEvent.NoGame.SetGameId(it)) },
                label = { Text("Game ID") }
            )

            Button(onClick = {
                state.eventSink(UiEvent.NoGame.JoinGameTapped)
            }) {
                Text(text = "Join")
            }
        }
    }
}