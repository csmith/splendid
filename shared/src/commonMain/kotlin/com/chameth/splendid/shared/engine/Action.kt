package com.chameth.splendid.shared.engine

interface Action<S : State> {

    fun resolve(state: S): List<Event<S>>

}