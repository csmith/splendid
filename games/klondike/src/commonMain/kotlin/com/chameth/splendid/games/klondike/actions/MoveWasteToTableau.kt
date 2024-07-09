package com.chameth.splendid.games.klondike.actions

import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.games.klondike.Phase
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.events.BuildTableauFromWaste
import com.chameth.splendid.games.klondike.rules.canAddCardToTableau

data class MoveWasteToTableau(val tableau: Int) : Action<State> {

    override fun resolve(state: State) = listOf(
        BuildTableauFromWaste(tableau = tableau)
    )

    companion object {
        fun generate(state: State) = buildList {
            if (state.phase == Phase.WaitingForMove && state.waste.isNotEmpty()) {
                state.tableau.forEachIndexed { index, cards ->
                    if (state.canAddCardToTableau(state.waste.flatten().last(), index)) {
                        add(MoveWasteToTableau(index))
                    }
                }
            }
        }
    }

}