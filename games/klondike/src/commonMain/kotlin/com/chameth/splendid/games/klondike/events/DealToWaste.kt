package com.chameth.splendid.games.klondike.events

import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.shared.engine.Event
import com.chameth.splendid.shared.playingcards.Card
import kotlinx.serialization.Serializable

@Serializable
data class DealToWaste(
    val cards: List<Card>
) : Event<State> {

    override fun resolve(state: State) = state.copy(
        stock = state.stock.drop(cards.size),
        waste = listOf(state.waste.flatten() + cards[0])
                + cards.drop(1).map { listOf(it) }
    )

}