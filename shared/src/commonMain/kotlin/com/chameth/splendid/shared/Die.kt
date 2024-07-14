package com.chameth.splendid.shared

import kotlinx.serialization.Serializable
import kotlin.random.Random

@Serializable
data class Die(val sides: Int, val value: Int) {
    fun roll() = copy(value = 1 + Random.nextInt(sides))
}

object Dice {
    val D6 = Die(6, 0)
}