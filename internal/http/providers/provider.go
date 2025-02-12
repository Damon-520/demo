package providers

import (
	"demoapi/internal/http/live_http"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	live_http.NewLiveRoomHttp,
)
