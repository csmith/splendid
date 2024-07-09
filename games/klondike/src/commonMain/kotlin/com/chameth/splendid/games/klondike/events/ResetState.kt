package com.chameth.splendid.games.klondike.events

import com.chameth.splendid.games.klondike.Phase
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.shared.engine.Event
import com.chameth.splendid.shared.playingcards.Card
import com.chameth.splendid.shared.util.repeat
import kotlinx.serialization.Serializable

@Serializable
data class ResetState(
    val stock: List<Card>
) : Event<State> {

    override fun resolve(state: State) = state.copy(
        phase = Phase.Unstarted,
        stock = stock,
        waste = emptyList(),
        tableau = emptyList<Card>().repeat(7),
        foundations = emptyList<Card>().repeat(4)
    )

}
