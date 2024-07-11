package com.chameth.splendid.server

import com.chameth.splendid.shared.engine.Game
import com.chameth.splendid.shared.transport.Message
import io.ktor.server.websocket.*
import io.ktor.websocket.*
import kotlinx.coroutines.flow.MutableSharedFlow
import kotlinx.coroutines.flow.consumeAsFlow
import kotlinx.coroutines.flow.filterIsInstance
import kotlinx.coroutines.flow.map
import kotlinx.coroutines.launch
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json
import kotlinx.serialization.modules.SerializersModule

class ClientSession(
    private val gameManager: GameManager,
    private val webSocketSession: WebSocketServerSession,
    private val id: String
) {

    private val json by lazy {
        Json {
            serializersModule = SerializersModule {
                gameManager.types.forEach {
                    include(it.serializersModule)
                }
            }
        }
    }

    private val sendQueue = MutableSharedFlow<Message.Server>()

    private var game: Game<*>? = null

    fun start() {
        webSocketSession.launch {
            sendQueue
                .map(json::encodeToString)
                .collect(webSocketSession::send)
        }

        webSocketSession.launch {
            webSocketSession.incoming.consumeAsFlow()
                .filterIsInstance<Frame.Text>()
                .map { json.decodeFromString<Message.Client>(it.readText()) }
                .collect(::processMessage)
        }
    }

    private suspend fun send(message: Message.Server) = sendQueue.emit(message)

    private suspend fun processMessage(message: Message.Client) {
        // TODO: Split this out into sensible chunks
        when (message) {
            is Message.Client.CreateGame -> {
                val newGame = gameManager.createGame(message.type)
                game = newGame
                send(Message.Server.MessageAcknowledged(message))
                send(Message.Server.GameJoined(newGame.id, newGame.type.name))

                webSocketSession.launch {
                    newGame.eventFlow.collect {
                        send(Message.Server.EventOccurred(it.event))
                    }
                }
            }

            is Message.Client.JoinGame -> TODO()
            Message.Client.LeaveGame -> TODO()

            is Message.Client.PerformAction -> {
                val localGame = game
                if (localGame != null) {
                    // TODO: Check if action was allowed
                    localGame.invoke(message.action)
                    send(Message.Server.MessageAcknowledged(message))
                } else {
                    send(Message.Server.MessageRejected(message, "Not joined to a game"))
                }
            }

            is Message.Client.SetId -> TODO()
            is Message.Client.SetName -> TODO()
        }
    }
}