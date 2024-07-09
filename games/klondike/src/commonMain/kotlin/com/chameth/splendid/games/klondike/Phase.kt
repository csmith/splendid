package com.chameth.splendid.games.klondike

import kotlinx.serialization.Serializable

@Serializable
enum class Phase {
    Unstarted,
    Dealing,
    WaitingForMove,
    Finished
}