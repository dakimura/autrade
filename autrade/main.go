package autrade

import (
	"awesomeProject/conf"
	"awesomeProject/logger"
	"fmt"
	"go.uber.org/zap"
)

func Run(config conf.MainConfig) error {
	err := logger.Init()
	if err != nil {
		return err
	}
	zap.L().Debug("initialized", zap.String("config", fmt.Sprintf("%v", config)))

	result := analyzer.Run(strategy)

	notifier.notifyOutput(result)

	return nil
}
