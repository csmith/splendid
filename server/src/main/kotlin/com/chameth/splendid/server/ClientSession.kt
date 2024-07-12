package com.chameth.splendid.server

import com.chameth.splendid.shared.transport.Message
import io.ktor.server.websocket.*
import io.ktor.websocket.*
import kotlinx.coroutines.flow.*
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

    suspend fun start() {
        webSocketSession.launch {
            sendQueue
                .map(json::encodeToString)
                .onEach { println("Sending: $it") }
                .collect(webSocketSession::send)
        }

        webSocketSession.incoming.consumeAsFlow()
            .filterIsInstance<Frame.Text>()
            .onEach { println("Sending: $it") }
            .map { json.decodeFromString<Message.Client>(it.readText()) }
            .collect(::processMessage)

        println("Finished session!")
    }

    private suspend fun send(message: Message.Server) = sendQueue.emit(message)

    private suspend fun processMessage(message: Message.Client) {
        // TODO: Split this out into sensible chunks
        when (message) {
            is Message.Client.CreateGame -> {
                val newGame = gameManager.createGame(message.gameType)
                game = newGame
                send(Message.Server.MessageAcknowledged(message))
                send(Message.Server.GameJoined(newGame.id, newGame.type.name))

                webSocketSession.launch {
                    newGame.eventFlow.collect {
                        send(Message.Server.EventOccurred(it.event))
                    }
                }
            }

            is Message.Client.JoinGame -> {
                val newGame = gameManager.getGame(message.gameId)
                if (newGame != null) {
                    // TODO: Dedupe
                    game = newGame
                    send(Message.Server.MessageAcknowledged(message))
                    send(Message.Server.GameJoined(newGame.id, newGame.type.name))

                    webSocketSession.launch {
                        newGame.eventFlow.collect {
                            send(Message.Server.EventOccurred(it.event))
                        }
                    }
                } else {
                    send(Message.Server.MessageRejected(message, "Game not found"))
                }
            }

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