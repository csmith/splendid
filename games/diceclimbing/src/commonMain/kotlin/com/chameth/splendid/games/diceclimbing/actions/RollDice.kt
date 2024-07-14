package com.chameth.splendid.games.diceclimbing.actions

import com.chameth.splendid.games.diceclimbing.Phase
import com.chameth.splendid.games.diceclimbing.State
import com.chameth.splendid.games.diceclimbing.events.SetDice
import com.chameth.splendid.shared.Dice
import com.chameth.splendid.shared.engine.Action
import kotlinx.serialization.Serializable

@Serializable
data class RollDice(override val actor: String) : Action<State> {

    override fun resolve(state: State) = listOf(
        SetDice(buildList {
            repeat(4) {
                add(Dice.D6.roll())
            }
        })
    )

    companion object {
        fun generate(state: State) = buildList {
            if (state.phase == Phase.WaitingForRoll || state.phase == Phase.WaitingForDecision) {
                state.currentPlayer?.let {
                    add(RollDice(it))
                }
            }
        }
    }
}
