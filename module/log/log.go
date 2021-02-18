package log

import (
	"jarvis/base/log"
	"jarvis/base/network"
)

type (
	logModule struct{}
)

const (
	ModuleName = "Log"
)

var (
	defaultLogModule = &logModule{}
)

func init() {}

func NewModule() network.Module {
	return defaultLogModule
}

func (lm *logModule) Name() string {
	return ModuleName
}

func (lm *logModule) Route() map[string][]network.RouteHandleFunc {
	return map[string][]network.RouteHandleFunc{
		"print": {lm.print},
	}
}

func (lm *logModule) print(ctx network.Context) {
	log.Info(string(ctx.Request().Data))
}
