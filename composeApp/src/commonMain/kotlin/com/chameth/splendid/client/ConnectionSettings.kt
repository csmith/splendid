package com.chameth.splendid.client

data class ConnectionSettings(
    val defaultProtocol: String,
    val defaultHost: String,
    val defaultPort: Int,
    val defaultPath: String,
    val autoConnect: Boolean,
    val autoJoin: String
)
