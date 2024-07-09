package com.chameth.splendid.games.klondike.ui.components

import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember
import androidx.compose.ui.Modifier
import com.chameth.splendid.games.klondike.ui.UiState
import com.chameth.splendid.shared.ui.components.Dialog
import dev.chrisbanes.haze.HazeState

@Composable
fun GameOverDialog(
    state: UiState,
    modifier: Modifier = Modifier,
    hazeState: HazeState = remember { HazeState() }
) {
    if (state.gameOver) {
        Dialog(modifier = modifier, hazeState = hazeState) {
            Text(
                text = "You win! Yay!",
                style = MaterialTheme.typography.headlineLarge
            )
        }
    }
}