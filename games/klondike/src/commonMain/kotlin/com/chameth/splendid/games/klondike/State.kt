package com.chameth.splendid.games.klondike

import com.chameth.splendid.shared.SystemActor
import com.chameth.splendid.shared.engine.State
import com.chameth.splendid.shared.playingcards.Card
import com.chameth.splendid.shared.util.repeat
import kotlinx.serialization.Serializable

@Serializable
data class State(
    val players: Set<String> = setOf(SystemActor), // TODO: Use real players
    val variant: Variant = Variant.DrawOne,
    val phase: Phase = Phase.Unstarted,
    val tableau: List<List<Card>> = emptyList<Card>().repeat(7),
    val stock: List<Card> = emptyList(),
    val waste: List<List<Card>> = emptyList(),
    val foundations: List<List<Card>> = emptyList<Card>().repeat(4)
) : State {
    fun foundationFor(card: Card) = foundations[card.suit.ordinal]
}