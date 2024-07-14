package com.chameth.splendid.games.diceclimbing.ui

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.remember
import androidx.compose.ui.Modifier
import com.chameth.splendid.games.diceclimbing.State
import com.chameth.splendid.games.diceclimbing.ui.components.Players
import com.chameth.splendid.games.diceclimbing.ui.components.Rolls
import com.chameth.splendid.games.diceclimbing.ui.components.Tableau
import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.shared.ui.components.Dialog

@Composable
fun Board(
    gameState: State,
    actionHandler: (Action<State>) -> Unit,
    modifier: Modifier = Modifier,
) {
    val presenter = remember { Presenter(actionHandler) }

    LaunchedEffect(gameState) {
        presenter.updateGameState(gameState)
    }

    val uiState = presenter.present()

    Box(modifier = modifier) {
        Row(modifier = Modifier.fillMaxSize()) {
            Tableau(
                modifier = Modifier.weight(1f),
                columns = uiState.board
            )

            Column {
                Players(uiState)
                Rolls(uiState)
            }
        }

        uiState.winner?.let {
            Dialog {
                Text("Game Over! ${uiState.winner} wins!")
            }
        }
    }
}