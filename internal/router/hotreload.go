package router

import (
	"net/http"
	"sync"
	"time"

	datastar "github.com/starfederation/datastar/sdk/go"
)

var hotReloadOnlyOnce sync.Once

func hotReloadHandler(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)
	hotReloadOnlyOnce.Do(func() {
		// Refresh the client page as soon as connection
		// is established. This will occur only once
		// after the server starts.
		sse.ExecuteScript(
			"window.location.reload()",
			datastar.WithExecuteScriptRetryDuration(time.Second),
		)
	})

	// Freeze the event stream until the connection
	// is lost for any reason. This will force the client
	// to attempt to reconnect after the server reboots.
	<-r.Context().Done()
}
