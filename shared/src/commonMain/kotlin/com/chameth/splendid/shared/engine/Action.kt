package com.chameth.splendid.shared.engine

interface Action<S : State> {

    val actor: String

    fun resolve(state: S): List<Event<S>>

}