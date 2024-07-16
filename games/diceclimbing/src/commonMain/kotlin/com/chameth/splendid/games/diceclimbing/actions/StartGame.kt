package com.chameth.splendid.games.diceclimbing.actions

import com.chameth.splendid.games.diceclimbing.Phase
import com.chameth.splendid.games.diceclimbing.State
import com.chameth.splendid.games.diceclimbing.events.SetTurn
import com.chameth.splendid.games.diceclimbing.events.SetTurnOrder
import com.chameth.splendid.games.diceclimbing.rules.canStartGame
import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.shared.engine.Event
import kotlinx.serialization.Serializable

@Serializable
data class StartGame(override val actor: String) : Action<State> {

    override fun resolve(state: State): List<Event<State>> {
        val order = state.players.keys.shuffled().toList()

        return listOf(
            SetTurnOrder(order),
            SetTurn(order[0]),
        )
    }

    companion object {
        fun generate(state: State) = buildList {
            if (state.phase == Phase.AssemblingPlayers && state.canStartGame()) {
                state.players.keys.forEach {
                    add(StartGame(it))
                }
            }
        }
    }
}
