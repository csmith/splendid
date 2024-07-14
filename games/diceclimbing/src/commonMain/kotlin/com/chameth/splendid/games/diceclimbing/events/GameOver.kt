package com.chameth.splendid.games.diceclimbing.events

import com.chameth.splendid.games.diceclimbing.Phase
import com.chameth.splendid.games.diceclimbing.State
import com.chameth.splendid.shared.engine.Event
import kotlinx.serialization.Serializable

@Serializable
data class GameOver(
    val winner: String
) : Event<State> {

    override fun resolve(state: State) = state.copy(
        phase = Phase.GameOver,
        currentPlayer = null,
        currentRoll = emptyList(),
        winner = winner
    )
    
}
