package com.chameth.splendid.ui

import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.MaterialTheme
import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember
import androidx.compose.ui.Modifier
import com.chameth.splendid.client.Client
import com.chameth.splendid.client.ConnectionSettings

@Composable
fun Root(connectionSettings: ConnectionSettings) {
    MaterialTheme {
        val presenter = remember { Presenter(Client(), connectionSettings) }

        when (val state = presenter.present()) {
            is UiState.NotConnected -> Connect(
                modifier = Modifier.fillMaxSize(),
                state = state
            )

            is UiState.NoGame -> GameSelector(
                modifier = Modifier.fillMaxSize(),
                state = state,
            )

            is UiState.InGame -> Game(
                modifier = Modifier.fillMaxSize(),
                state = state
            )
        }
    }
}