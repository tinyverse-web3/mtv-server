package task

import (
	"context"
	"mtv/models"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/task"
	"github.com/gorilla/websocket"
)

func CheckRelayStatusTask() *task.Task {
	// 0 0 */1 * * *   0/30 * * * * *
	t := task.NewTask("tk1", "0 0,30 * * * *", func(ctx context.Context) error {
		logs.Info("check relay status task start---------------------------------------------")
		startTime := time.Now()
		// updateRelayStatus()
		var data []models.NostrRelay
		o := orm.NewOrm()

		var relay models.NostrRelay
		qt := o.QueryTable(relay)
		qt.All(&data)
		for _, value := range data {
			go updateRelayStatus(o, value)
		}
		endTime := time.Now()
		logs.Info("end time - start time = ", endTime.Sub(startTime).Seconds())
		logs.Info("check relay status task end---------------------------------------------")
		return nil
	})

	// check task
	// err := t.Run(context.Background())
	// if err != nil {
	// 	logs.Error(err)
	// }

	return t
}

func updateRelayStatus(o orm.Ormer, value models.NostrRelay) {
	// 状态置为检查中
	value.Status = 2
	_, err := o.Update(&value)
	if err != nil {
		logs.Error(err)
		return
	}
	// 检查连接状态
	conn, _, err := websocket.DefaultDialer.Dial(value.WsServer, nil)
	if err != nil {
		value.Status = 0 // 连接失败
		value.Remark = err.Error()
		logs.Error("Error connecting to Websocket Server:", err)
	} else {
		value.Status = 1 // 连接成功
		value.Remark = ""
		logs.Info("Connecting to Websocket Server Success")
		conn.Close()
	}
	//更新状态
	_, err = o.Update(&value)
	if err != nil {
		logs.Error(err)
	}
}
