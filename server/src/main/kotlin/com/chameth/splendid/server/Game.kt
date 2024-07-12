package com.chameth.splendid.server

import com.chameth.splendid.shared.SystemActor
import com.chameth.splendid.shared.engine.*
import com.chameth.splendid.shared.util.now
import kotlinx.coroutines.delay
import kotlinx.coroutines.flow.MutableSharedFlow
import kotlinx.coroutines.flow.SharedFlow
import kotlinx.datetime.LocalDateTime
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json

class Game<S : State>(
    val type: GameType<S>,
    val id: String
) {

    private val events = mutableListOf<Event<S>>()

    private val _eventFlow = MutableSharedFlow<EventWrapper<S>>(replay = Int.MAX_VALUE)

    private val actions = type.actionsGenerator(type.stateFactory()).toMutableList()

    private val json by lazy { Json { serializersModule = type.serializersModule } }

    val eventFlow: SharedFlow<EventWrapper<S>>
        get() = _eventFlow

    private var state: S = type.stateFactory()

    suspend fun invoke(action: Action<*>) {
        if (action !in actions) {
            println("Illegal action attempted: $action")
            return
        }

        @Suppress("UNCHECKED_CAST")
        (action as Action<S>).resolve(state).forEach { event ->
            emit(event)
            delay(50)
        }
    }

    fun save() = json.encodeToString(
        Save(
        version = type.version,
        game = type.name,
        gameId = id,
        events = events
    )
    )

    suspend fun load(data: String) {
        val save = json.decodeFromString<Save>(data)
        require(save.game == type.name) { "Save game is of type ${save.game} but this game is for ${type.name}." }
        require(save.version == type.version) { "Save version is different. Got: ${save.version}, expected: ${type.version}" }

        state = type.stateFactory()
        save.events.forEach {
            @Suppress("UNCHECKED_CAST")
            emit(it as Event<S>)
        }
    }

    @Suppress("UNCHECKED_CAST")
    suspend fun applyRemoteEvent(event: Event<*>) = emit(event as Event<S>)

    private suspend fun emit(event: Event<S>) {
        events += event
        state = event.resolve(state)
        _eventFlow.emit(EventWrapper(
            event = event,
            state = state,
            actions = type.actionsGenerator(state).also {
                actions.clear()
                actions.addAll(it)
            }
        ))
    }

    data class EventWrapper<S : State>(
        val event: Event<S>,
        val actor: String = SystemActor,
        val time: LocalDateTime = LocalDateTime.now(),
        val state: S,
        val actions: List<Action<S>>
    )

}