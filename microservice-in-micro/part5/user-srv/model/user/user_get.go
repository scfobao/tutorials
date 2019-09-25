package user

import (
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part5/plugins/db"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part5/plugins/graylog"
	proto "github.com/micro-in-cn/tutorials/microservice-in-micro/part5/user-srv/proto/user"
	"github.com/micro/go-micro/util/log"
)

func (s *service) QueryUserByName(userName string) (ret *proto.User, err error) {
	queryString := `SELECT user_id, user_name, pwd FROM user WHERE user_name = ?`

	// 获取数据库
	o := db.GetDB()

	ret = &proto.User{}
	graylog.GetLog(nil).Write(map[string]interface{}{"data": ret, "tag": "QueryUserByName"})
	// 查询
	err = o.QueryRow(queryString, userName).Scan(&ret.Id, &ret.Name, &ret.Pwd)
	if err != nil {
		log.Logf("[QueryUserByName] 查询数据失败，err：%s", err)
		return
	}
	return
}
