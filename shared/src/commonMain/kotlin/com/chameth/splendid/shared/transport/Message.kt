package com.chameth.splendid.shared.transport

import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.shared.engine.Event
import kotlinx.serialization.Serializable

@Serializable
sealed interface Message {

    @Serializable
    sealed interface Client : Message {

        @Serializable
        data class SetId(val clientId: String) : Client

        @Serializable
        data class SetName(val name: String) : Client

        @Serializable
        data class JoinGame(val gameId: String) : Client

        @Serializable
        data class CreateGame(val type: String) : Client

        @Serializable
        data class PerformAction(val action: Action<*>) : Client

        @Serializable
        data object LeaveGame : Client

    }

    @Serializable
    sealed interface Server : Message {

        @Serializable
        data class MessageAcknowledged(val original: Client) : Server

        @Serializable
        data class MessageRejected(val original: Client, val error: String) : Server

        @Serializable
        data class GameJoined(val gameId: String, val type: String) : Server

        @Serializable
        data class EventOccurred(val event: Event<*>) : Server

    }
}