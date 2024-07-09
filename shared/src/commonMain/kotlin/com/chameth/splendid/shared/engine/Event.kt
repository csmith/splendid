package com.chameth.splendid.shared.engine

import kotlinx.serialization.Polymorphic

@Polymorphic
interface Event<S : State> {

    fun resolve(state: S): S

}
