package com.chameth.splendid.games.klondike.actions

import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.shared.engine.Event
import com.chameth.splendid.games.klondike.Phase
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.Variant
import com.chameth.splendid.games.klondike.events.DealToWaste
import kotlin.math.min

data object DrawFromStock : Action<State> {

    override fun resolve(state: State): List<Event<State>> {
        val cards = min(
            if (state.variant == Variant.DrawThree) 3 else 1,
            state.stock.size
        )

        return listOf(
            DealToWaste(
                cards = state.stock.take(cards).map{ it.copy(visible = true) }
            )
        )
    }

    fun generate(state: State) = buildList {
        if (state.phase == Phase.WaitingForMove && state.stock.isNotEmpty()) {
            add(DrawFromStock)
        }
    }

}