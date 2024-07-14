package com.chameth.splendid.games.diceclimbing.rules

import com.chameth.splendid.games.diceclimbing.Phase
import com.chameth.splendid.games.diceclimbing.State

fun State.canSitMorePlayers() = phase == Phase.AssemblingPlayers && players.size < 4