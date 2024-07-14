package com.chameth.splendid.games.diceclimbing.events

import com.chameth.splendid.games.diceclimbing.State
import com.chameth.splendid.shared.engine.Event
import kotlinx.serialization.Serializable

@Serializable
data class SetTurnOrder(
    val turnOrder: List<String>
) : Event<State> {

    override fun resolve(state: State) = state.copy(
        turnOrder = turnOrder
    )

}
