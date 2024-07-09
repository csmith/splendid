package com.chameth.splendid.games.klondike.ui

import androidx.compose.foundation.layout.*
import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.ui.components.*
import com.chameth.splendid.shared.engine.Game
import dev.chrisbanes.haze.HazeState
import dev.chrisbanes.haze.haze

@Composable
fun Board(
    game: Game<State>,
    modifier: Modifier = Modifier
) {
    val presenter = remember { Presenter(game) }
    val state = presenter.present()
    val hazeState = remember { HazeState() }

    Box {
        GameBoard(
            modifier = Modifier
                .fillMaxSize()
                .haze(hazeState)
        ) {
            Column(
                modifier = modifier
                    .padding(16.dp)
                    .requiredSizeIn(minWidth = 800.dp, maxWidth = 1000.dp)
                    .align(Alignment.Center)
            ) {
                Row(horizontalArrangement = Arrangement.spacedBy(16.dp)) {
                    Stock(state)
                    Waste(state)

                    Spacer(modifier = Modifier.weight(1f))

                    Foundations(state)
                }

                Spacer(modifier = Modifier.height(40.dp))

                Tableau(
                    modifier = Modifier.fillMaxWidth(),
                    state = state
                )
            }
        }

        ButtonBar(
            modifier = Modifier
                .align(Alignment.BottomEnd)
                .padding(16.dp),
            state = state
        )

        GameOverDialog(
            modifier = Modifier.align(Alignment.Center),
            hazeState = hazeState,
            state = state
        )

        StartGameDialog(
            modifier = Modifier.align(Alignment.Center),
            hazeState = hazeState,
            state = state
        )
    }
}
