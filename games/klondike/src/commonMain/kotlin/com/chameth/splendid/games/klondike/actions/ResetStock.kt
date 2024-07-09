package com.chameth.splendid.games.klondike.actions

import com.chameth.splendid.games.klondike.Phase
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.events.MoveWasteToStock
import com.chameth.splendid.shared.engine.Action

data object ResetStock : Action<State> {

    override fun resolve(state: State) = listOf(MoveWasteToStock)

    fun generate(state: State) = buildList {
        if (state.phase == Phase.WaitingForMove && state.stock.isEmpty() && state.waste.isNotEmpty()) {
            add(ResetStock)
        }
    }

}