package com.chameth.splendid.games.diceclimbing.events

import com.chameth.splendid.games.diceclimbing.Phase
import com.chameth.splendid.games.diceclimbing.State
import com.chameth.splendid.shared.Die
import com.chameth.splendid.shared.engine.Event
import kotlinx.serialization.Serializable

@Serializable
data class SetDice(
    val rolls: List<Die>
) : Event<State> {

    override fun resolve(state: State) = state.copy(
        phase = Phase.WaitingForDiceSelection,
        currentRoll = rolls
    )
    
}
