package com.chameth.splendid.shared.playingcards

object Decks {
    val noJokers = Suit.entries.flatMap { suit ->
        Rank.entries.map { rank ->
            Card(suit = suit, rank = rank)
        }
    }
}