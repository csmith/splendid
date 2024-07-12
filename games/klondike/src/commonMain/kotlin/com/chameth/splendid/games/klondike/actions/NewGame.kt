package com.chameth.splendid.games.klondike.actions

import com.chameth.splendid.games.klondike.Phase
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.events.ResetState
import com.chameth.splendid.shared.engine.Action
import kotlinx.serialization.Serializable

@Serializable
data class NewGame(override val actor: String) : Action<State> {
    override fun resolve(state: State) = listOf(
        ResetState(emptyList())
    )

    companion object {
        fun generate(state: State) = buildList {
            if (state.phase != Phase.Unstarted) {
                state.players.forEach {
                    add(NewGame(it))
                }
            }
        }
    }
}