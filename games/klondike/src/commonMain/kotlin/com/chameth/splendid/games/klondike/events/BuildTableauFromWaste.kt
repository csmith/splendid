package com.chameth.splendid.games.klondike.events

import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.shared.engine.Event
import com.chameth.splendid.shared.util.replaceNth
import kotlinx.serialization.Serializable

@Serializable
data class BuildTableauFromWaste(val tableau: Int) : Event<State> {

    override fun resolve(state: State) = state.copy(
        tableau = state.tableau.replaceNth(tableau) { it + state.waste.flatten().last() },
        waste = if (state.waste.size > 1)
            state.waste.take(state.waste.size - 1)
        else
            state.waste.replaceNth(0) { it.dropLast(1) }.filterNot { it.isEmpty() }
    )

}