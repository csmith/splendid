package com.chameth.splendid.games.klondike.events

import com.chameth.splendid.games.klondike.Phase
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.shared.engine.Event
import kotlinx.serialization.Serializable

@Serializable
data class SetPhase(
    val phase: Phase
) : Event<State> {

    override fun resolve(state: State) = state.copy(phase = phase)

}
