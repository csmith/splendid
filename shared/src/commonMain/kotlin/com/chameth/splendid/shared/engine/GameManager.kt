package com.chameth.splendid.shared.engine

import com.chameth.splendid.shared.util.generateAca
import kotlinx.coroutines.FlowPreview
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.flow.debounce
import kotlinx.coroutines.launch
import kotlin.time.Duration.Companion.seconds

class GameManager(val types: List<GameType<*>>) {

    private val games = mutableMapOf<String, Game<*>>()

    @OptIn(FlowPreview::class)
    suspend fun createGame(type: String) = types
        .first { type.equals(it.name, ignoreCase = true) }
        .let { Game(it, generateAca()) }
        .also { games[it.id] = it }
        .also {
            // TODO: This should be scoped to the game and cancelled at some point
            GlobalScope.launch {
                it.eventFlow.debounce(1.seconds).collect { _ ->
                    save(it)
                }
            }
        }

    private fun save(game: Game<out State>) {
        println("Would save ${game.id}")
    }

}