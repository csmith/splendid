package com.chameth.splendid.games.klondike.events

import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.shared.engine.Event
import com.chameth.splendid.shared.util.replaceNth
import kotlinx.serialization.Serializable

@Serializable
data class BuildFoundationFromTableau(
    val tableau: Int,
) : Event<State> {

    override fun resolve(state: State): State {
        val card = state.tableau[tableau].last()
        return state.copy(
            foundations = state.foundations.replaceNth(card.suit.ordinal) {
                it + card
            },
            tableau = state.tableau.replaceNth(tableau) {
                it.dropLast(1)
            }
        )
    }

}