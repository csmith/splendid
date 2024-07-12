package com.chameth.splendid.client

import com.chameth.splendid.shared.engine.GameType
import com.chameth.splendid.shared.engine.State

data class ClientState(
    val connected: Boolean,
    val gameType: GameType<*>?,
    val state: State?,
)
