package com.chameth.splendid.client

import com.chameth.splendid.shared.transport.Message
import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.plugins.websocket.*
import io.ktor.http.*
import io.ktor.websocket.*
import kotlinx.coroutines.flow.MutableSharedFlow
import kotlinx.coroutines.flow.consumeAsFlow
import kotlinx.coroutines.flow.filterIsInstance
import kotlinx.coroutines.flow.map
import kotlinx.coroutines.launch
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json

class Client(private val json: Json) {

    private val receiveQueue = MutableSharedFlow<Message.Server>(extraBufferCapacity = Int.MAX_VALUE)
    private val sendQueue = MutableSharedFlow<Message.Client>()

    private val client = HttpClient(CIO) {
        install(WebSockets)
    }

    suspend fun connect(host: String, port: Int, path: String) {
        client.webSocket(method = HttpMethod.Get, host = host, port = port, path = path) {
            send("Hello!")

            launch {
                sendQueue
                    .map(json::encodeToString)
                    .collect(::send)
            }

            launch {
                incoming.consumeAsFlow()
                    .filterIsInstance<Frame.Text>()
                    .map { json.decodeFromString<Message.Server>(it.readText()) }
                    .collect(receiveQueue::tryEmit)
            }
        }
    }

    suspend fun send(message: Message.Client) {
        sendQueue.emit(message)
        // TODO: block and wait for a response?
    }
}