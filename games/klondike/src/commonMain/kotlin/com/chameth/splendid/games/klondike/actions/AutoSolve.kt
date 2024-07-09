package com.chameth.splendid.games.klondike.actions

import com.chameth.splendid.games.klondike.Phase
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.events.BuildFoundationFromTableau
import com.chameth.splendid.games.klondike.events.SetPhase
import com.chameth.splendid.games.klondike.rules.canAutoSolve
import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.shared.playingcards.Rank

data object AutoSolve : Action<State> {

    override fun resolve(state: State) = buildList {
        Rank.entries.forEach { rank ->
            state.tableau.forEachIndexed { index, cards ->
                if (cards.any { it.rank == rank }) {
                    add(BuildFoundationFromTableau(index))
                }
            }
        }

        add(SetPhase(phase = Phase.Finished))
    }

    fun generate(state: State) = buildList {
        if (state.canAutoSolve()) {
            add(AutoSolve)
        }
    }

}