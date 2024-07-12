package com.chameth.splendid.shared.events

import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.shared.engine.Event
import com.chameth.splendid.shared.engine.State

data class PerformAction<T : State>(
    val action: Action<T>
) : Event<T> {

    override fun resolve(state: T) = state

}