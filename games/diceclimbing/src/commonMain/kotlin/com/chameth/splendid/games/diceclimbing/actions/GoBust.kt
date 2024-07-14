package com.chameth.splendid.games.diceclimbing.actions

import com.chameth.splendid.games.diceclimbing.Phase
import com.chameth.splendid.games.diceclimbing.State
import com.chameth.splendid.games.diceclimbing.events.RemoveBlackTokens
import com.chameth.splendid.games.diceclimbing.events.SetTurn
import com.chameth.splendid.games.diceclimbing.rules.canPlayAny
import com.chameth.splendid.games.diceclimbing.rules.nextPlayer
import com.chameth.splendid.shared.engine.Action
import kotlinx.serialization.Serializable

@Serializable
data class GoBust(override val actor: String) : Action<State> {

    override fun resolve(state: State) = buildList {
        add(RemoveBlackTokens)
        add(SetTurn(state.nextPlayer()))
    }

    companion object {
        fun generate(state: State) = buildList {
            if (state.phase == Phase.WaitingForDiceSelection && !state.canPlayAny()) {
                add(GoBust(state.currentPlayer!!))
            }
        }
    }

}
