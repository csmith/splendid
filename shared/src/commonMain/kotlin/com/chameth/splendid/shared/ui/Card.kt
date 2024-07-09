package com.chameth.splendid.shared.ui

//import androidx.compose.desktop.ui.tooling.preview.Preview
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.BoxScope
import androidx.compose.foundation.layout.aspectRatio
import androidx.compose.foundation.layout.requiredHeight
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.unit.dp

@Composable
fun Card(
    modifier: Modifier = Modifier,
    composable: @Composable BoxScope.() -> Unit
) {
    Box(
        modifier = modifier
            .requiredHeight(LocalCardHeight.current)
            .clip(RoundedCornerShape(8.dp))
            .aspectRatio(0.66759f, matchHeightConstraintsFirst = true),
    ) {
        composable()
    }
}

// @Preview
@Composable
internal fun Card_Preview() = MaterialTheme {
    Card {
        Text("This is some text", modifier = Modifier.align(Alignment.Center))
    }
}