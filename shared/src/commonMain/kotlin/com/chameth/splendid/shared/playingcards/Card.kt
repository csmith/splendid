package com.chameth.splendid.shared.playingcards

import kotlinx.serialization.Serializable

@Serializable
data class Card(val suit: Suit, val rank: Rank, val visible: Boolean = true) {
    override fun toString() = if (visible) "${rank}${suit}" else "??"
}