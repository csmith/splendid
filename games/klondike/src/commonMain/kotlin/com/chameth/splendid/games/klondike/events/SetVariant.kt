package com.chameth.splendid.games.klondike.events

import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.Variant
import com.chameth.splendid.shared.engine.Event
import kotlinx.serialization.Serializable

@Serializable
data class SetVariant(
    val variant: Variant
) : Event<State> {

    override fun resolve(state: State) = state.copy(variant = variant)

}
