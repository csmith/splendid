package com.chameth.splendid.shared.engine

import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import kotlinx.serialization.modules.SerializersModule

interface GameType<S : State> {

    val stateFactory: () -> S

    val actionsGenerator: (S) -> List<Action<S>>

    val uiRoot: @Composable (S, (Action<S>) -> Unit, Modifier) -> Unit

    val version: Int

    val serializersModule: SerializersModule

    val name: String

    fun mask(state: S, event: Event<*>, actor: String): Event<*>

    fun newAddPlayerEvent(playerId: String): Event<*>

    @Composable
    fun root(
        state: State,
        actionSink: (Action<State>) -> Unit,
        modifier: Modifier
    ) {
        @Suppress("UNCHECKED_CAST")
        uiRoot(
            state as S,
            { actionSink(it as Action<State>) },
            modifier
        )
    }

}