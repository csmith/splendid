package com.chameth.splendid.games.klondike.events

import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.shared.engine.Event
import com.chameth.splendid.shared.util.replaceNth
import kotlinx.serialization.Serializable

@Serializable
data class BuildTableauFromFoundation(
    val tableau: Int,
    val foundation: Int
) : Event<State> {
    override fun resolve(state: State) = state.copy(
        tableau = state.tableau.replaceNth(tableau) { it + state.foundations[foundation].last() },
        foundations = state.foundations.replaceNth(foundation) { it.dropLast(1) }
    )
}