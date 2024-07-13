package com.chameth.splendid.server

import com.chameth.splendid.games.all.Games
import io.ktor.server.application.*
import io.ktor.server.engine.*
import io.ktor.server.http.content.*
import io.ktor.server.netty.*
import io.ktor.server.routing.*
import io.ktor.server.websocket.*
import java.io.File

val server = GameServer(GameManager(Games.available))

fun main() {
    embeddedServer(
        factory = Netty,
        port = 8080,
        host = "0.0.0.0",
        module = Application::module
    ).start(wait = true)
}

fun Application.module() {
    install(WebSockets)

    routing {
        staticFiles("/", File("static"))

        webSocket("/client") {
            server.adoptWebSocket(this)
            println("Lost websocket")
        }
    }
}