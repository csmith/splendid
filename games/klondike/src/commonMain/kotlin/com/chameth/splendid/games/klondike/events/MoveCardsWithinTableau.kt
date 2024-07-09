package com.chameth.splendid.games.klondike.events

import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.shared.engine.Event
import com.chameth.splendid.shared.util.replaceNth
import kotlinx.serialization.Serializable

@Serializable
data class MoveCardsWithinTableau(
    val from: Int,
    val to: Int,
    val count: Int
) : Event<State> {

    override fun resolve(state: State) = state.copy(
        tableau = state.tableau.replaceNth(to) {
            it + state.tableau[from].takeLast(count)
        }.replaceNth(from) {
            it.dropLast(count)
        }
    )

}