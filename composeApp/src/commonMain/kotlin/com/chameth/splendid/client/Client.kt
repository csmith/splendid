package com.chameth.splendid.client

import com.chameth.splendid.games.all.Games
import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.shared.engine.Event
import com.chameth.splendid.shared.engine.GameType
import com.chameth.splendid.shared.engine.State
import com.chameth.splendid.shared.transport.Message
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.launch

class Client {

    private val socket = ClientSocket(Games.serializer)

    private var gameType: GameType<*>? = null
    private var gameState: State? = null
    private var clientId: String? = null

    private val stateFlow = MutableStateFlow(ClientState())

    val state: StateFlow<ClientState>
        get() = stateFlow

    // TODO: Client should have its own coroutine scope, this is horrid
    fun connect(coroutineScope: CoroutineScope, host: String, port: Int, path: String) {
        GlobalScope.launch {
            socket.incoming.collect {
                handleMessage(it)
            }
        }

        GlobalScope.launch {
            println("Connecting")
            socket.connect(host, port, path)
            println("Done connecting")
            stateFlow.emit(stateFlow.value.copy(connected = false, gameType = null, state = null, gameId = null))
        }

        coroutineScope.launch {
            stateFlow.emit(stateFlow.value.copy(connected = true))
        }
    }

    suspend fun createGame(type: String) {
        socket.send(Message.Client.CreateGame(type))
    }

    suspend fun joinGame(id: String) {
        socket.send(Message.Client.JoinGame(id))
    }

    suspend fun performAction(action: Action<*>) {
        socket.send(Message.Client.PerformAction(action))
    }

    private suspend fun handleMessage(message: Message.Server) = when (message) {
        is Message.Server.EventOccurred -> {
            @Suppress("UNCHECKED_CAST")
            gameState = (message.event as Event<State>).resolve(gameState!!)
            stateFlow.emit(stateFlow.value.copy(state = gameState))
        }

        is Message.Server.GameJoined -> {
            gameType = Games.available.firstOrNull { it.name == message.gameType }
            gameType?.let {
                gameState = it.stateFactory()
            }
            println("Game joined. Type = $gameType, State = $gameState")
            stateFlow.emit(stateFlow.value.copy(
                gameType = gameType,
                gameId = message.gameId,
                state = gameState
            ))
        }

        is Message.Server.MessageAcknowledged -> {}
        is Message.Server.MessageRejected -> {}
        is Message.Server.YourId -> {
            clientId = message.id
            stateFlow.emit(stateFlow.value.copy(clientId = clientId))
        }
    }
}