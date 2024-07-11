package com.chameth.splendid.games.klondike.actions

import com.chameth.splendid.games.klondike.Phase
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.events.BuildTableauFromWaste
import com.chameth.splendid.games.klondike.rules.canAddCardToTableau
import com.chameth.splendid.shared.engine.Action

data class MoveWasteToTableau(override val actor: String, val tableau: Int) : Action<State> {

    override fun resolve(state: State) = listOf(
        BuildTableauFromWaste(tableau = tableau)
    )

    companion object {
        fun generate(state: State) = buildList {
            if (state.phase == Phase.WaitingForMove && state.waste.isNotEmpty()) {
                state.tableau.forEachIndexed { index, _ ->
                    if (state.canAddCardToTableau(state.waste.flatten().last(), index)) {
                        state.players.forEach {
                            add(MoveWasteToTableau(it, index))
                        }
                    }
                }
            }
        }
    }

}