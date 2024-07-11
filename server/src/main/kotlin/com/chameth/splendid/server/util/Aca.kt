package com.chameth.splendid.server.util

private val adjectives by lazy { readResource("adjectives.txt") }
private val colours by lazy { readResource("colours.txt") }
private val animals by lazy { readResource("animals.txt") }

fun generateAca() = "${adjectives.random()}-${colours.random()}-${animals.random()}"

private fun readResource(name: String) =
    object {}.javaClass.getResourceAsStream(name)?.bufferedReader()?.readLines().orEmpty()
