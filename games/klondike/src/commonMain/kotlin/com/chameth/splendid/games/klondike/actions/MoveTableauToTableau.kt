package com.chameth.splendid.games.klondike.actions

import com.chameth.splendid.games.klondike.Phase
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.events.MoveCardsWithinTableau
import com.chameth.splendid.games.klondike.events.RevealCardInTableau
import com.chameth.splendid.games.klondike.rules.canAddCardToTableau
import com.chameth.splendid.shared.engine.Action
import kotlinx.serialization.Serializable

@Serializable
data class MoveTableauToTableau(
    override val actor: String,
    val from: Int,
    val to: Int,
    val count: Int
) : Action<State> {

    override fun resolve(state: State) = buildList {
        add(MoveCardsWithinTableau(from, to, count))

        val sourceTableau = state.tableau[from]
        if (sourceTableau.size > count && !sourceTableau[sourceTableau.size-count-1].visible) {
            add(
                RevealCardInTableau(
                    tableau = from,
                    card = sourceTableau[sourceTableau.size-count-1].copy(visible = true)
                )
            )
        }
    }

    companion object {
        fun generate(state: State) = buildList {
            if (state.phase == Phase.WaitingForMove) {
                addAll(
                    state.tableau
                        .flatMapIndexed { source, tableau ->
                            tableau.reversed().takeWhile { it.visible }.flatMapIndexed { index, card ->
                                state.tableau.indices
                                    .filter { j -> source != j && state.canAddCardToTableau(card, j) }
                                    .flatMap { dest ->
                                        state.players.map {
                                            MoveTableauToTableau(it, source, dest, index + 1)
                                        }
                                    }
                            }
                        }
                )
            }
        }
    }
}