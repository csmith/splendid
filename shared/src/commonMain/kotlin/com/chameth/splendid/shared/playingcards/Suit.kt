package com.chameth.splendid.shared.playingcards

import kotlinx.serialization.Serializable

@Serializable
enum class Suit(val black: Boolean, private val short: String) {
    Spades(true, "♠"),
    Hearts(false, "♥"),
    Clubs(true, "♣"),
    Diamonds(false, "♦");

    override fun toString() = short
}