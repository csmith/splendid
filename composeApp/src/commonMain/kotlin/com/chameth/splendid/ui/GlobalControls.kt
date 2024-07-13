package com.chameth.splendid.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalClipboardManager
import androidx.compose.ui.text.AnnotatedString
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontFamily
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp

@Composable
fun GlobalControls(
    state: UiState.InGame,
    modifier: Modifier = Modifier,
) {
    Row(
        modifier = modifier
            .clip(RoundedCornerShape(topEnd = 8.dp))
            .background(Color.Gray.copy(alpha = 0.5f))
            .padding(vertical = 4.dp, horizontal = 8.dp),
        horizontalArrangement = Arrangement.spacedBy(16.dp)
    ) {
        Text(
            text = "This is Splendid!",
            style = MaterialTheme.typography.labelSmall
        )

        val clipboardManager = LocalClipboardManager.current

        Text(
            modifier = Modifier.clickable {
                clipboardManager.setText(AnnotatedString(state.gameId))
            },
            text = buildAnnotatedString {
                append("Game ID: ")
                withStyle(SpanStyle(fontFamily = FontFamily.Monospace)) {
                    append(state.gameId)
                }
            },
            style = MaterialTheme.typography.labelSmall,
        )
    }
}