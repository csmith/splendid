package com.chameth.splendid.games.klondike.events

import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.shared.engine.Event
import kotlinx.serialization.Serializable

@Serializable
data object MoveWasteToStock : Event<State> {

    override fun resolve(state: State) = state.copy(
        stock = state.waste.flatten().map { it.copy(visible = false) },
        waste = emptyList()
    )

}