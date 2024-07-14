package com.chameth.splendid.games.all

import com.chameth.splendid.games.diceclimbing.DiceClimbing
import com.chameth.splendid.games.klondike.Klondike
import com.chameth.splendid.shared.engine.GameType
import kotlinx.serialization.json.Json
import kotlinx.serialization.modules.SerializersModule

object Games {
    val available = listOf<GameType<*>>(
        DiceClimbing,
        Klondike
    )

    val serializer by lazy {
        Json {
            serializersModule = SerializersModule {
                available.forEach {
                    include(it.serializersModule)
                }
            }
        }
    }
}
