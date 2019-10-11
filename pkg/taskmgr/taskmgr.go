package taskmgr

import (
	"encoding/json"
	"go.uber.org/zap"
	"time"
	"xtc/sofa/connect"
	"xtc/sofa/log"
	"xtc/sofa/model"
)

func Start() {
	loop()
}

// 负责从redis队列获取数据
func loop() {
	for {
		// timeout 0s

		msg := connect.RedisClient().BRPop(0*1000*1000*1000, "calls").Val()[1]

		var call model.Call
		err := json.Unmarshal([]byte(msg), &call)

		if err != nil {
			log.Logger.Error("unmarshal msg err", zap.Error(err))
			continue
		}

		call.SubmitTime = time.Now()

		go do(&call)

	}
}

func do(call *model.Call) {
	call.Exec()
	call.Save()
}
