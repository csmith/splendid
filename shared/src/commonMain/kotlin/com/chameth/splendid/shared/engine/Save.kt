package com.chameth.splendid.shared.engine

import com.chameth.splendid.shared.util.now
import kotlinx.datetime.LocalDateTime
import kotlinx.serialization.Serializable

@Serializable
data class Save(
    val version: Int,
    val game: String,
    val gameId: String,
    val events: List<Event<*>>,
    val saveTime: LocalDateTime = LocalDateTime.now()
)