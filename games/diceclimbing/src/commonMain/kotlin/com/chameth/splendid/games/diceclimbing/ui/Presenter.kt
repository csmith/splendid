package com.chameth.splendid.games.diceclimbing.ui

import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import com.chameth.splendid.games.diceclimbing.Phase
import com.chameth.splendid.games.diceclimbing.State
import com.chameth.splendid.games.diceclimbing.actions.*
import com.chameth.splendid.games.diceclimbing.rules.*
import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.shared.ui.LocalClientId
import kotlinx.coroutines.flow.MutableStateFlow

class Presenter(private val actionHandler: (Action<State>) -> Unit) {

    private val gameState = MutableStateFlow(State())

    suspend fun updateGameState(state: State) = gameState.emit(state)

    @Composable
    fun present(): UiState {
        val state by gameState.collectAsState()
        val clientId = LocalClientId.current

        return UiState(
            canJoin = state.canSitMorePlayers() && clientId !in state.players,
            canStart = state.canStartGame() && clientId in state.players,
            players = state.players.map { (id, token) ->
                token to if (id == clientId) "You" else id.substring(0..7)
            },
            board = state.board,
            dice = state.currentRoll,
            canRoll = (state.phase == Phase.WaitingForRoll || state.phase == Phase.WaitingForDecision) && state.currentPlayer == clientId,
            canStop = state.phase == Phase.WaitingForDecision && state.currentPlayer == clientId,
            goBust = state.phase == Phase.WaitingForDiceSelection && state.currentPlayer == clientId && !state.canPlayAny(),
            options = if (state.phase == Phase.WaitingForDiceSelection && state.currentPlayer == clientId)
                state.options()
            else
                emptyList(),
            winner = state.winner,
            eventSink = {
                when (it) {
                    UiEvent.JoinGameClicked -> actionHandler(SitDown(clientId))
                    UiEvent.StartGameClicked -> actionHandler(StartGame(clientId))
                    is UiEvent.AdvanceColumnsClicked -> actionHandler(AdvanceColumns(clientId, it.columns))
                    UiEvent.GoBustClicked -> actionHandler(GoBust(clientId))
                    UiEvent.RollClicked -> actionHandler(RollDice(clientId))
                    UiEvent.StopClicked -> actionHandler(Stop(clientId))
                }
            }
        )
    }

    private fun State.options() = buildList<List<Int>> {
        rollPermutations.forEach { (first, second) ->
            if (canPlayInColumns(first, second)) {
                add(listOf(first, second))
            } else if (canPlayInColumn(first)) {
                add(listOf(first))
            } else if (canPlayInColumn(second)) {
                add(listOf(second))
            }
        }
    }

}