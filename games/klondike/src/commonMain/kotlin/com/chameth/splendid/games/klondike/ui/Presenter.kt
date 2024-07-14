package com.chameth.splendid.games.klondike.ui

import androidx.compose.runtime.*
import com.chameth.splendid.games.klondike.Phase
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.rules.canAutoSolve
import com.chameth.splendid.games.klondike.ui.interactors.*
import com.chameth.splendid.games.klondike.ui.model.Selection
import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.shared.playingcards.Card
import com.chameth.splendid.shared.ui.LocalClientId
import kotlinx.coroutines.flow.MutableStateFlow

class Presenter(
    private val actionHandler: (Action<State>) -> Unit
) {

    private val gameState = MutableStateFlow(State())

    suspend fun updateState(state: State) = gameState.emit(state)

    @Composable
    fun present(): UiState {
        val state by gameState.collectAsState()

        var selected by remember { mutableStateOf<Selection?>(null) }

        LaunchedEffect(state) {
            selected = null
        }

        val clientId = LocalClientId.current

        return UiState(
            notStarted = state.phase == Phase.Unstarted,
            gameOver = state.phase == Phase.Finished,
            hasStock = state.stock.isNotEmpty(),
            canAutoSolve = state.canAutoSolve(),
            waste = state.waste.map { it.last().toSelectable(selected) },
            foundations = state.foundations.map { it.lastOrNull()?.toSelectable(selected) },
            tableau = state.tableau.map { it.toSelectable(selected) },
            eventSink = { event ->
                selected = when (event) {
                    UiEvent.StockClicked -> stockClicked(state, clientId, selected, actionHandler)
                    UiEvent.WasteClicked -> wasteClicked(state, clientId, selected, actionHandler)
                    UiEvent.RestartClicked -> restartClicked(state, clientId, selected, actionHandler)
                    UiEvent.AutoSolveClicked -> autoSolveClicked(state, clientId, selected, actionHandler)
                    is UiEvent.TableauClicked -> tableauClicked(state, clientId, selected, event.tableau, event.card, actionHandler)
                    is UiEvent.FoundationClicked -> foundationClicked(state, clientId, selected, event.foundation, actionHandler)
                    is UiEvent.StartGameClicked -> startClicked(state, clientId, selected, event.variant, actionHandler)
                }
            }
        )
    }

    private fun Card.toSelectable(selected: Selection?) = SelectableCard(
        card = this,
        selected = this == selected?.card
    )

    private fun List<Card>.toSelectable(selected: Selection?) = buildList {
        var foundSelected = false
        this@toSelectable.forEach { card ->
            if (card == selected?.card) {
                foundSelected = true
            }
            add(SelectableCard(card, foundSelected))
        }
    }

}