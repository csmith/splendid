package com.chameth.splendid.shared.ui

// import androidx.compose.desktop.ui.tooling.preview.Preview
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.MaterialTheme
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.unit.dp

@Composable
fun CardPlaceholder(modifier: Modifier = Modifier) {
    Card(modifier = modifier) {
        Box(
            modifier = Modifier
                .fillMaxSize()
                .border(1.dp, Color.White, RoundedCornerShape(8.dp))
                .background(Color(53, 101, 77))
        ) {

        }
    }
}

// @Preview
@Composable
internal fun CardPlaceholder_Preview() {
    MaterialTheme {
        CardBack(modifier = Modifier.height(120.dp))
    }
}
