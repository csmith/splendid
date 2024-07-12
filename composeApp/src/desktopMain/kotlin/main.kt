import androidx.compose.runtime.LaunchedEffect
import androidx.compose.ui.unit.dp
import androidx.compose.ui.window.Window
import androidx.compose.ui.window.application
import androidx.compose.ui.window.rememberWindowState
import com.chameth.splendid.ui.Root

fun main() = application {
    val state = rememberWindowState(width = 1024.dp, height = 800.dp)

    LaunchedEffect(state.size) {
        if (state.size.width < 900.dp) {
            state.size = state.size.copy(width = 900.dp)
        }

        if (state.size.height < 800.dp) {
            state.size = state.size.copy(height = 800.dp)
        }
    }

    Window(
        onCloseRequest = ::exitApplication,
        title = "Splendid!",
        state = state
    ) {
        Root()
    }
}