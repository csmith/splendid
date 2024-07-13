package com.chameth.splendid.games.klondike.events

import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.shared.engine.Event
import kotlinx.serialization.Serializable

@Serializable
data class AddPlayer(val playerId: String) : Event<State> {
    override fun resolve(state: State) = state.copy(players = state.players + playerId)
}