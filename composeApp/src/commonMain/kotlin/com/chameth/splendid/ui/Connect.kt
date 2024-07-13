package com.chameth.splendid.ui

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.material3.*
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp

@Composable
fun Connect(
    state: UiState.NotConnected,
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
                text = "Connect to a server to play",
                style = MaterialTheme.typography.bodyMedium
            )

            Row(verticalAlignment = Alignment.CenterVertically) {
                Checkbox(
                    checked = state.secure,
                    onCheckedChange = { state.eventSink(UiEvent.NotConnected.SetSecure(it)) }
                )
                Text("Secure connection")
            }

            TextField(
                value = state.host,
                onValueChange = { state.eventSink(UiEvent.NotConnected.SetHost(it)) },
                label = { Text("Host") }
            )

            TextField(
                value = state.port.toString(),
                onValueChange = { state.eventSink(UiEvent.NotConnected.SetPort(it)) },
                label = { Text("Port") }
            )

            TextField(
                value = state.path,
                onValueChange = { state.eventSink(UiEvent.NotConnected.SetPath(it)) },
                label = { Text("Path") }
            )

            Button(onClick = { state.eventSink(UiEvent.NotConnected.ConnectTapped)}) {
                Text(text = "Connect")
            }
        }
    }
}