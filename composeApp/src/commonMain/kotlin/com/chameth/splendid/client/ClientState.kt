package com.chameth.splendid.client

import com.chameth.splendid.shared.engine.GameType
import com.chameth.splendid.shared.engine.State

data class ClientState(
    val connected: Boolean = false,
    val gameType: GameType<*>? = null,
    val state: State? = null,
    val gameId: String? = null,
    val clientId: String? = null
)
