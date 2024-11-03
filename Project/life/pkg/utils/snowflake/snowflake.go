package snowflake

//雪花算法生成ID
import (
	"errors"
	"time"

	"github.com/bwmarrin/snowflake"
)

var (
	node           *snowflake.Node
	errorSnowflake = errors.New("snowflake init failed")
)

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime) //第一个参数是时间格式占位符
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000 //将起始时间转换为毫秒级时间戳
	node, err = snowflake.NewNode(machineID)
	return
}
func GenID() int64 {
	return node.Generate().Int64()
}
