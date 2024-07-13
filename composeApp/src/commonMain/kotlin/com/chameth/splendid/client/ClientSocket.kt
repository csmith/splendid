package com.chameth.splendid.client

import com.chameth.splendid.shared.transport.Message
import io.ktor.client.*
import io.ktor.client.plugins.websocket.*
import io.ktor.client.request.*
import io.ktor.http.*
import io.ktor.websocket.*
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json

class ClientSocket(private val json: Json) {

    private val receiveQueue = MutableSharedFlow<Message.Server>(extraBufferCapacity = Int.MAX_VALUE)
    private val sendQueue = MutableSharedFlow<Message.Client>()

    val incoming: SharedFlow<Message.Server>
        get() = receiveQueue

    private val client = HttpClient {
        install(WebSockets)
    }

    suspend fun connect(secure: Boolean, host: String, port: Int, path: String) {
        println("Connecting to ws${if (secure) "s" else ""}://$host:$port$path")

        // TODO: Nicer way to do this?
        val method : suspend (HttpMethod, String?, Int?, String?, HttpRequestBuilder.() -> Unit, suspend DefaultClientWebSocketSession.() -> Unit) -> Unit = if (secure) client::wss else client::ws

        method(HttpMethod.Get, host, port, path, {}) {
            println("Socket session")
            launch {
                sendQueue
                    .map(json::encodeToString)
                    .onEach { println("Sent: $it") }
                    .collect(::send)
            }

            println("Consuming events")
            incoming.consumeAsFlow()
                .filterIsInstance<Frame.Text>()
                .map { it.readText() }
                .onEach { println("Received: $it") }
                .map { json.decodeFromString<Message.Server>(it) }
                .collect(receiveQueue::tryEmit)

            println("Done")
        }
    }

    suspend fun send(message: Message.Client) {
        sendQueue.emit(message)
        // TODO: block and wait for a response?
    }
}