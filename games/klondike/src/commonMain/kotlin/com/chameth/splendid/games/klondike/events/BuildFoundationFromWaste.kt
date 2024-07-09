package com.chameth.splendid.games.klondike.events

import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.shared.engine.Event
import com.chameth.splendid.shared.util.replaceNth
import kotlinx.serialization.Serializable

@Serializable
data object BuildFoundationFromWaste : Event<State> {

    override fun resolve(state: State): State {
        val card = state.waste.flatten().last()
        return state.copy(
            foundations = state.foundations.replaceNth(card.suit.ordinal) {
                it + card
            },
            waste = if (state.waste.size > 1)
                state.waste.take(state.waste.size - 1)
            else
                state.waste.replaceNth(0) { it.dropLast(1) }.filterNot { it.isEmpty() }
        )
    }

}