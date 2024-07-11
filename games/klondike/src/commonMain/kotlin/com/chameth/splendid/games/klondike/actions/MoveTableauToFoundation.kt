package com.chameth.splendid.games.klondike.actions

import com.chameth.splendid.games.klondike.Phase
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.events.BuildFoundationFromTableau
import com.chameth.splendid.games.klondike.events.RevealCardInTableau
import com.chameth.splendid.games.klondike.events.SetPhase
import com.chameth.splendid.games.klondike.rules.canBuildFoundationWithCard
import com.chameth.splendid.games.klondike.rules.willWin
import com.chameth.splendid.shared.engine.Action

data class MoveTableauToFoundation(
    override val actor: String,
    val tableau: Int,
) : Action<State> {

    override fun resolve(state: State) = buildList {
        add(BuildFoundationFromTableau(tableau = tableau))

        if (state.willWin()) {
            add(SetPhase(phase = Phase.Finished))
            return@buildList
        }

        if (state.tableau[tableau].size > 1) {
            add(
                RevealCardInTableau(
                    tableau = tableau,
                    card = state.tableau[tableau][state.tableau[tableau].size - 2].copy(visible = true)
                )
            )
        }
    }

    companion object {
        fun generate(state: State) = buildList {
            if (state.phase == Phase.WaitingForMove) {
                addAll(
                    state.tableau
                        .mapIndexed { i, tableau -> i to tableau.lastOrNull() }
                        .mapNotNull { (i, topCard) -> topCard?.let { i to it } }
                        .filter { (_, topCard) -> state.canBuildFoundationWithCard(topCard) }
                        .flatMap { (i, _) ->
                            state.players.map {
                                MoveTableauToFoundation(actor = it, tableau = i)
                            }
                        }
                )
            }
        }
    }
}