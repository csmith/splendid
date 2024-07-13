import androidx.compose.ui.ExperimentalComposeUiApi
import androidx.compose.ui.window.ComposeViewport
import com.chameth.splendid.client.ConnectionSettings
import com.chameth.splendid.ui.Root
import kotlinx.browser.document

@OptIn(ExperimentalComposeUiApi::class)
fun main() {
    ComposeViewport(document.body!!) {
        val isSecure = document.location?.protocol == "https:"

        Root(
            ConnectionSettings(
                defaultProtocol = if (isSecure) "wss" else "ws",
                defaultHost = document.location?.host.orEmpty(),
                defaultPort = document.location?.port?.toIntOrNull() ?: if (isSecure) 443 else 80,
                defaultPath = "/client",
                autoConnect = true
            )
        )
    }
}