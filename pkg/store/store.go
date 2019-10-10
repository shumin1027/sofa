package store

import (
	"encoding/json"
	"go.uber.org/zap"
	"strings"
	"xtc/sofa/connect"
	"xtc/sofa/log"
	"xtc/sofa/model"
)

func Save(call *model.Call) {

	// json 编码
	j, err := json.Marshal(call)
	if err != nil {
		log.Logger.Error("marshal json error", zap.Error(err))
	}
	log.Logger.Info("received from client", zap.String("data", string(j)))

	// 存入redis给logstash继续处理
	redis := connect.RedisClient()
	redis.LPush(strings.ToLower(call.Platform+"-"+call.Command), j)
	log.Logger.Info("push to redis successed")

}
