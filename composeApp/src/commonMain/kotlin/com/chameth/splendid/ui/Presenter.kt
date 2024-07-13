package com.chameth.splendid.ui

import androidx.compose.runtime.*
import com.chameth.splendid.client.Client
import com.chameth.splendid.client.ClientState
import com.chameth.splendid.games.all.Games
import kotlinx.coroutines.launch

class Presenter(private val client: Client) {

    @Composable
    fun present(): UiState {
        val clientState by client.state.collectAsState()
        return when {
            !clientState.connected -> presentDisconnected()
            clientState.gameType == null -> presentNoGame()
            else -> presentGame(clientState)
        }
    }

    @Composable
    private fun presentDisconnected(): UiState {
        var host by remember { mutableStateOf("localhost") }
        var port by remember { mutableStateOf(8080) }
        var path by remember { mutableStateOf("/client") }
        val scope = rememberCoroutineScope()

        return UiState.NotConnected(
            host = host,
            port = port,
            path = path,
            eventSink = {
                when (it) {
                    UiEvent.NotConnected.ConnectTapped -> client.connect(scope, host, port, path)
                    is UiEvent.NotConnected.SetHost -> host = it.host
                    is UiEvent.NotConnected.SetPath -> path = it.path
                    is UiEvent.NotConnected.SetPort -> port = it.port.replace(Regex("[^0-9]"), "").toInt()
                }
            }
        )
    }

    @Composable
    private fun presentNoGame(): UiState {
        var gameId by remember { mutableStateOf("") }
        val scope = rememberCoroutineScope()

        return UiState.NoGame(
            gameId = gameId,
            availableTypes = Games.available.map { it.name },
            eventSink = {
                when (it) {
                    is UiEvent.NoGame.CreateGameTapped -> scope.launch { client.createGame(it.type) }
                    UiEvent.NoGame.JoinGameTapped -> scope.launch { client.joinGame(gameId) }
                    is UiEvent.NoGame.SetGameId -> gameId = it.gameId
                }
            }
        )
    }

    @Composable
    private fun presentGame(clientState: ClientState): UiState.InGame {
        val coroutineScope = rememberCoroutineScope()
        return UiState.InGame(
            gameType = clientState.gameType!!,
            state = clientState.state!!,
            gameId = clientState.gameId!!,
            clientId = clientState.clientId,
            actionSink = {
                coroutineScope.launch {
                    client.performAction(it)
                }
            }
        )
    }

}