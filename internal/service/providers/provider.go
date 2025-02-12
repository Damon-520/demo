package providers

import (
	"demoapi/internal/service/live_service"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	live_service.NewLiveRoomService,
)
