package com.chameth.splendid.games.diceclimbing.actions

import com.chameth.splendid.games.diceclimbing.Phase
import com.chameth.splendid.games.diceclimbing.State
import com.chameth.splendid.games.diceclimbing.Token
import com.chameth.splendid.games.diceclimbing.events.SitPlayer
import com.chameth.splendid.games.diceclimbing.rules.canSitMorePlayers
import com.chameth.splendid.shared.engine.Action
import kotlinx.serialization.Serializable

@Serializable
data class SitDown(override val actor: String) : Action<State> {
    override fun resolve(state: State) = listOf(
        SitPlayer(actor, state.randomUnusedToken)
    )

    private val State.unusedTokens: Set<Token>
        get() = setOf(Token.Red, Token.Blue, Token.Green, Token.Yellow) - players.values.toSet()

    private val State.randomUnusedToken: Token
        get() = unusedTokens.shuffled().first()

    companion object {
        fun generate(state: State) = buildList {
            if (state.phase == Phase.AssemblingPlayers && state.canSitMorePlayers()) {
                (state.lobbyMembers - state.players.keys.toSet()).forEach {
                    add(SitDown(it))
                }
            }
        }
    }
}
