package com.chameth.splendid.shared.playingcards

import kotlinx.serialization.Serializable

@Serializable
enum class Rank(private val short: String) {
    Ace("A"),
    Two("2"),
    Three("3"),
    Four("4"),
    Five("5"),
    Six("6"),
    Seven("7"),
    Eight("8"),
    Nine("9"),
    Ten("T"),
    Jack("J"),
    Queen("Q"),
    King("K");

    override fun toString() = short
}