package com.chameth.splendid.client

import com.chameth.splendid.shared.transport.Message
import io.ktor.client.*
import io.ktor.client.plugins.websocket.*
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

    suspend fun connect(host: String, port: Int, path: String) {
        println("Connecting to ws://$host:$port$path")
        client.webSocket(method = HttpMethod.Get, host = host, port = port, path = path) {
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