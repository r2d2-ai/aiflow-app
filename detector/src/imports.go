package main

import (
	_ "github.com/r2d2-ai/aiflow/action/flow"
	_ "github.com/r2d2-ai/aiflow/action/flow/activity/subflow"
	_ "github.com/r2d2-ai/aiflow/activity/common/actreturn"
	_ "github.com/r2d2-ai/aiflow/activity/common/log"
	_ "github.com/r2d2-ai/aiflow/activity/vision/object_detection"
	_ "github.com/r2d2-ai/aiflow/activity/vision/transform"
	_ "github.com/r2d2-ai/aiflow/function/utils"
	_ "github.com/r2d2-ai/aiflow/trigger/net/ipcam"
)
