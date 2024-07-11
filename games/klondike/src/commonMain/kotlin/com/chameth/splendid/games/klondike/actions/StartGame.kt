package com.chameth.splendid.games.klondike.actions

import com.chameth.splendid.games.klondike.Phase
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.Variant
import com.chameth.splendid.games.klondike.events.DealToTableau
import com.chameth.splendid.games.klondike.events.ResetState
import com.chameth.splendid.games.klondike.events.SetPhase
import com.chameth.splendid.games.klondike.events.SetVariant
import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.shared.playingcards.Decks

data class StartGame(override val actor: String, val variant: Variant) : Action<State> {

    override fun resolve(state: State) = buildList {
        val stock = Decks.noJokers.shuffled().map { it.copy(visible = false) }

        add(ResetState(stock = stock))
        add(SetPhase(Phase.Dealing))

        var n = 0
        for (i in 0..6) {
            for (j in i..6) {
                add(DealToTableau(tableau = j, card = stock[n++].copy(visible = i == j)))
            }
        }

        add(SetVariant(variant = variant))
        add(SetPhase(phase = Phase.WaitingForMove))
    }

    companion object {
        fun generate(state: State) = buildList {
            if (state.phase == Phase.Unstarted) {
                Variant.entries.forEach { variant ->
                    state.players.forEach { actor ->
                        add(StartGame(actor, variant))
                    }
                }
            }
        }
    }
}