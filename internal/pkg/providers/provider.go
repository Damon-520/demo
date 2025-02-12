package pkg

import (
	"demoapi/internal/pkg/alipay_service"
	"demoapi/internal/pkg/middlewares/auth"
	"demoapi/internal/pkg/sidx"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	sidx.NewSid,
	auth.NewAdminAuth,
	alipay_service.NewAlipayService,
)
