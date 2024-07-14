package com.chameth.splendid.games.diceclimbing.rules

import com.chameth.splendid.games.diceclimbing.Phase
import com.chameth.splendid.games.diceclimbing.State

fun State.canStartGame() = phase == Phase.AssemblingPlayers && players.size in 2..4