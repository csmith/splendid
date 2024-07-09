package com.chameth.splendid.games.klondike.events

import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.shared.engine.Event
import com.chameth.splendid.shared.playingcards.Card
import com.chameth.splendid.shared.util.replaceNth
import kotlinx.serialization.Serializable

@Serializable
data class DealToTableau(
    val tableau: Int,
    val card: Card
) : Event<State> {

    override fun resolve(state: State) = state.copy(
        stock = state.stock.drop(1),
        tableau = state.tableau.replaceNth(tableau) { cards -> cards + card }
    )

}