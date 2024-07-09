package com.chameth.splendid.games.klondike.actions

import com.chameth.splendid.games.klondike.Phase
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.events.ResetState
import com.chameth.splendid.shared.engine.Action

data object NewGame : Action<State> {
    override fun resolve(state: State) = listOf(
        ResetState(emptyList())
    )

    fun generate(state: State) = buildList {
        if (state.phase != Phase.Unstarted) {
            add(NewGame)
        }
    }
}