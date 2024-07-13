package com.chameth.splendid.ui

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import com.chameth.splendid.shared.ui.LocalClientId

@Composable
fun Game(
    state: UiState.InGame,
    modifier: Modifier = Modifier
) {
    CompositionLocalProvider(LocalClientId provides state.clientId.orEmpty()) {
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
}