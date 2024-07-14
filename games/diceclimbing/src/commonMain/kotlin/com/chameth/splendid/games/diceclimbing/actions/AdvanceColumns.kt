package com.chameth.splendid.games.diceclimbing.actions

import com.chameth.splendid.games.diceclimbing.Phase
import com.chameth.splendid.games.diceclimbing.State
import com.chameth.splendid.games.diceclimbing.events.AdvanceBlackTokens
import com.chameth.splendid.games.diceclimbing.rules.canPlayInColumn
import com.chameth.splendid.games.diceclimbing.rules.canPlayInColumns
import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.shared.engine.Event
import kotlinx.serialization.Serializable

@Serializable
data class AdvanceColumns(override val actor: String, val columns: List<Int>) : Action<State> {

    override fun resolve(state: State): List<Event<State>> {
        return listOf(
            AdvanceBlackTokens(state.currentPlayerToken!!, columns),
        )
    }

    companion object {
        fun generate(state: State) = buildList {
            if (state.phase == Phase.WaitingForDiceSelection) {
                state.rollPermutations.forEach {
                    if (state.canPlayInColumns(it.first, it.second)) {
                        add(AdvanceColumns(state.currentPlayer!!, listOf(it.first, it.second)))
                    } else if (state.canPlayInColumn(it.first)) {
                        add(AdvanceColumns(state.currentPlayer!!, listOf(it.first)))
                    } else if (state.canPlayInColumn(it.second)) {
                        add(AdvanceColumns(state.currentPlayer!!, listOf(it.second)))
                    }
                }
            }
        }.distinct()
    }
}
