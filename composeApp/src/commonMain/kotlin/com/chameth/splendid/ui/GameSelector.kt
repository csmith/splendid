package com.chameth.splendid.ui

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.material3.Button
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.chameth.splendid.shared.engine.Game
import com.chameth.splendid.shared.engine.GameManager
import com.chameth.splendid.shared.engine.GameType
import kotlinx.coroutines.launch

@Composable
fun GameSelector(manager: GameManager, modifier: Modifier = Modifier) {
    var selectedType by remember { mutableStateOf<GameType<*>?>(null) }
    var game by remember { mutableStateOf<Game<*>?>(null) }

    if (game == null) {
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
                    text = "Select a game to play",
                    style = MaterialTheme.typography.bodyMedium
                )

                val scope = rememberCoroutineScope()
                manager.types.forEach {
                    Button(onClick = {
                        selectedType = it
                        scope.launch {
                            game = manager.createGame(it.name)
                        }
                    }) {
                        Text(text = it.name)
                    }
                }
            }
        }
    } else {
        game?.let {
            selectedType?.uiRoot?.invoke(it, modifier)
        }
    }
}