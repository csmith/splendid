package com.chameth.splendid.shared.engine

import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import kotlinx.serialization.modules.SerializersModule

interface GameType<S : State> {

    val stateFactory: () -> S

    val actionsGenerator: (S) -> List<Action<S>>

    val uiRoot: @Composable (Game<*>, Modifier) -> Unit

    val version: Int

    val serializersModule: SerializersModule

    val name: String

}