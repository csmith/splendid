package com.chameth.splendid.games.klondike.actions

import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.games.klondike.Phase
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.events.BuildTableauFromFoundation
import com.chameth.splendid.games.klondike.rules.canAddCardToTableau

data class MoveFoundationToTableau(
    val foundation: Int,
    val tableau: Int
) : Action<State> {

    override fun resolve(state: State) = listOf(
        BuildTableauFromFoundation(
            tableau = tableau,
            foundation = foundation
        )
    )

    companion object {
        fun generate(state: State) = buildList<Action<State>> {
            if (state.phase == Phase.WaitingForMove) {
                state.foundations.forEachIndexed { foundation, fCards ->
                    if (fCards.isNotEmpty()) {
                        state.tableau.forEachIndexed { tableau, _ ->
                            if (state.canAddCardToTableau(fCards.last(), tableau)) {
                                add(MoveFoundationToTableau(foundation, tableau))
                            }
                        }
                    }
                }
            }
        }
    }

}
