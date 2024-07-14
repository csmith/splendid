package com.chameth.splendid.games.diceclimbing.events

import com.chameth.splendid.games.diceclimbing.State
import com.chameth.splendid.games.diceclimbing.Token
import com.chameth.splendid.shared.engine.Event
import kotlinx.serialization.Serializable

@Serializable
data class SitPlayer(
    val player: String,
    val token: Token
) : Event<State> {

    override fun resolve(state: State) = state.copy(
        players = state.players + (player to token)
    )

}
