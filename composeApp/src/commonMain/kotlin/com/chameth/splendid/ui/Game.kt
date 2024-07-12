package com.chameth.splendid.ui

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier

@Composable
fun Game(
    state: UiState.InGame,
    modifier: Modifier = Modifier
) {
    Box(modifier = modifier) {
        state.gameType.root(
            modifier = Modifier.fillMaxSize(),
            state = state.state,
            actionSink = state.actionSink
        )

        GlobalControls(
            modifier = Modifier.align(Alignment.BottomStart),
            state = state
        )
    }
}