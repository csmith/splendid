package com.chameth.splendid.games.klondike.events

import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.shared.engine.Event
import com.chameth.splendid.shared.playingcards.Card
import com.chameth.splendid.shared.util.replaceNth
import kotlinx.serialization.Serializable

@Serializable
data class RevealCardInTableau(
    val tableau: Int,
    val card: Card
) : Event<State> {

    override fun resolve(state: State) = state.copy(
        tableau = state.tableau.replaceNth(tableau) { cards ->
            cards.dropLast(1) + card
        }
    )

}