package apps

// 业务加载区

import (
	// api impl
	_ "go18/book/v4/apps/book/api"

	// service impl
	_ "go18/book/v4/apps/book/impl"
	_ "go18/book/v4/apps/comment/impl"
)
