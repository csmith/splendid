package com.chameth.splendid.ui

sealed interface UiEvent {
    sealed interface NotConnected : UiEvent {
        data object ConnectTapped : NotConnected
        data class SetHost(val host: String) : NotConnected
        data class SetPort(val port: String) : NotConnected
        data class SetPath(val path: String) : NotConnected
    }

    sealed interface NoGame : UiEvent {
        data class CreateGameTapped(val type: String) : NoGame
        data class SetGameId(val gameId: String) : NoGame
        data object JoinGameTapped : NoGame
    }
}