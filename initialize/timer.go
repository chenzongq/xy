package initialize

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/xyxy/admin-service/config"
	"github.com/xyxy/admin-service/utils"

	"github.com/xyxy/admin-service/global"
)

func Timer() {
	if global.XY_CONFIG.Timer.Start {
		for i := range global.XY_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				var option []cron.Option
				if global.XY_CONFIG.Timer.WithSeconds {
					option = append(option, cron.WithSeconds())
				}
				_, err := global.XY_Timer.AddTaskByFunc("ClearDB", global.XY_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.XY_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				}, option...)
				if err != nil {
					fmt.Println("add timer error:", err)
				}
			}(global.XY_CONFIG.Timer.Detail[i])
		}
	}
}
