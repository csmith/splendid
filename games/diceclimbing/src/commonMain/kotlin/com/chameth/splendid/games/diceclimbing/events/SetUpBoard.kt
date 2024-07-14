package com.chameth.splendid.games.diceclimbing.events

import com.chameth.splendid.games.diceclimbing.Column
import com.chameth.splendid.games.diceclimbing.State
import com.chameth.splendid.shared.engine.Event
import kotlinx.serialization.Serializable

@Serializable
data object SetUpBoard : Event<State> {

    override fun resolve(state: State) = state.copy(
        board = listOf(
            Column(2, 3),
            Column(3, 5),
            Column(4, 7),
            Column(5, 9),
            Column(6, 11),
            Column(7, 13),
            Column(8, 11),
            Column(9, 9),
            Column(10, 7),
            Column(11, 5),
            Column(12, 3),
        )
    )

}
