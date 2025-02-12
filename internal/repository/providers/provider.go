package providers

import (
	"demoapi/internal/repository"
	"demoapi/internal/repository/xes_activity/live_room"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	repository.NewActivityDB,
	live_room.NewLiveRoomDao,
)
