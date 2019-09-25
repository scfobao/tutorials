package handler

import "testing"
import "context"
import us "github.com/micro-in-cn/tutorials/microservice-in-micro/part3/user-srv/proto/user"

func TestQueryUserByName(t *testing.T) {
	rsp, err := QueryUserByName(context.TODO(), &us.Request{
		UserName: "micro",
	})
	if 1 == 1 {
		t.Error("test")
	}

}
