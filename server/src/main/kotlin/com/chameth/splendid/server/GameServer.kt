package com.chameth.splendid.server

import io.ktor.server.websocket.*
import java.util.*

class GameServer(private val gameManager: GameManager) {

    private val clients = mutableMapOf<String, ClientSession>()

    suspend fun adoptWebSocket(session: WebSocketServerSession) {
        val id = UUID.randomUUID().toString()
        val client = ClientSession(
            gameManager = gameManager,
            webSocketSession = session,
            id = id
        )
        client.start()
        clients[id] = client
    }

}