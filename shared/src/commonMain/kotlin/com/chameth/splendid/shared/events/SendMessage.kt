package com.chameth.splendid.shared.events

import com.chameth.splendid.shared.engine.Event
import com.chameth.splendid.shared.engine.State

data class SendMessage<T : State>(
    val actor: String,
    val message: String
) : Event<T> {

    override fun resolve(state: T) = state

}