package com.chameth.splendid.games.klondike.ui

import androidx.compose.runtime.*
import com.chameth.splendid.games.klondike.Phase
import com.chameth.splendid.games.klondike.State
import com.chameth.splendid.games.klondike.rules.canAutoSolve
import com.chameth.splendid.games.klondike.ui.interactors.*
import com.chameth.splendid.games.klondike.ui.model.Selection
import com.chameth.splendid.shared.engine.Action
import com.chameth.splendid.shared.engine.Game
import com.chameth.splendid.shared.playingcards.Card
import kotlinx.coroutines.launch

class Presenter(private val game: Game<State>) {

    @Composable
    fun present(): UiState {
        val eventState by game.eventFlow.collectAsState(null)
        val state = eventState?.state ?: State()
        val coroutineScope = rememberCoroutineScope()

        var selected by remember { mutableStateOf<Selection?>(null) }

        LaunchedEffect(state) {
            selected = null
        }

        fun invoke(action: Action<State>) = coroutineScope.launch {
            game.invoke(action)
        }

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
                    UiEvent.StockClicked -> stockClicked(state, selected, ::invoke)
                    UiEvent.WasteClicked -> wasteClicked(state, selected, ::invoke)
                    UiEvent.RestartClicked -> restartClicked(state, selected, ::invoke)
                    UiEvent.AutoSolveClicked -> autoSolveClicked(state, selected, ::invoke)
                    is UiEvent.TableauClicked -> tableauClicked(state, selected, event.tableau, event.card, ::invoke)
                    is UiEvent.FoundationClicked -> foundationClicked(state, selected, event.foundation, ::invoke)
                    is UiEvent.StartGameClicked -> startClicked(state, selected, event.variant, ::invoke)
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