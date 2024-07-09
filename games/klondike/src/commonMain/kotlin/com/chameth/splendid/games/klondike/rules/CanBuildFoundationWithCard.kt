package com.chameth.splendid.games.klondike.rules

import com.chameth.splendid.shared.playingcards.Card
import com.chameth.splendid.shared.playingcards.Rank
import com.chameth.splendid.games.klondike.State

fun State.canBuildFoundationWithCard(card: Card): Boolean {
    val foundation = foundationFor(card)

    return if (foundation.isEmpty())
        card.rank == Rank.Ace
    else
        foundation.last().rank.ordinal == card.rank.ordinal - 1
}
