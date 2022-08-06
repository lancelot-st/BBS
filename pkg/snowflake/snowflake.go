package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

var Node *snowflake.Node

//先进性初始化节点
func Init(machineID int64) error {
	var st time.Time
	st, err := time.Parse("2006-01-02", "2022-06-01")
	if err != nil {
		return err
	}
	//设置时间
	snowflake.Epoch = st.UnixNano() / 1000000
	Node, err = snowflake.NewNode(machineID)
	if err != nil {
		return err
	}
	return nil
}

func GenID() int64 {
	return Node.Generate().Int64()
}
