package com.chameth.splendid.games.klondike.rules

import com.chameth.splendid.shared.playingcards.Card
import com.chameth.splendid.shared.playingcards.Rank
import com.chameth.splendid.games.klondike.State

fun State.canAddCardToTableau(card: Card, tableauIndex: Int) =
    if (tableau[tableauIndex].isEmpty())
        card.rank == Rank.King
    else
        tableau[tableauIndex].last().let { t ->
            t.suit.black != card.suit.black && t.rank.ordinal == card.rank.ordinal + 1
        }
