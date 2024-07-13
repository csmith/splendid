package com.chameth.splendid.ui

import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.shared.engine.GameType
import com.chameth.splendid.shared.engine.State

sealed interface UiState {
    data class NotConnected(
        val host: String,
        val port: Int,
        val path: String,
        val eventSink: (UiEvent.NotConnected) -> Unit
    ) : UiState

    data class NoGame(
        val gameId: String,
        val availableTypes: List<String>,
        val eventSink: (UiEvent.NoGame) -> Unit
    ) : UiState

    data class InGame(
        val gameType: GameType<*>,
        val gameId: String,
        val clientId: String?,
        val state: State,
        val actionSink: (Action<*>) -> Unit
    ) : UiState
}
