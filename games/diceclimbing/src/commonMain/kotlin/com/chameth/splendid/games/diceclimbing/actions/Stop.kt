package com.chameth.splendid.games.diceclimbing.actions

import com.chameth.splendid.games.diceclimbing.Phase
import com.chameth.splendid.games.diceclimbing.State
import com.chameth.splendid.games.diceclimbing.events.GameOver
import com.chameth.splendid.games.diceclimbing.events.RemoveBlackTokens
import com.chameth.splendid.games.diceclimbing.events.SaveBlackTokens
import com.chameth.splendid.games.diceclimbing.events.SetTurn
import com.chameth.splendid.games.diceclimbing.rules.nextPlayer
import com.chameth.splendid.games.diceclimbing.rules.willWinIfStopping
import com.chameth.splendid.shared.engine.Action
import kotlinx.serialization.Serializable

@Serializable
data class Stop(override val actor: String) : Action<State> {

    override fun resolve(state: State) = buildList {
        add(SaveBlackTokens(state.currentPlayerToken!!))
        add(RemoveBlackTokens)

        if (state.willWinIfStopping()) {
            add(GameOver(state.currentPlayer!!))
        } else {
            add(SetTurn(state.nextPlayer()))
        }
    }

    companion object {
        fun generate(state: State) = buildList {
            if (state.phase == Phase.WaitingForDecision) {
                state.currentPlayer?.let {
                    add(Stop(it))
                }
            }
        }
    }

}
