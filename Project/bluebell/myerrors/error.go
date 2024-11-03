package myerrors

import "errors"

var (
	ErrorGetCurrentDir  = errors.New("获取当前工作目录失败")
	ErrorConfPathNil    = errors.New("配置文件路径为空")
	ErrorConfPathNone   = errors.New("未找到配置文件")
	ErrorViperRead      = errors.New("viper读取配置文件失败")
	ErrorViperUnmarshal = errors.New("viper反序列化配置文件失败")

	InitSettings  = errors.New("设置初始化失败")
	InitLogger    = errors.New("logger初始化失败")
	InitMysql     = errors.New("mysql初始化失败")
	InitRedis     = errors.New("redis初始化失败")
	InitSnowflake = errors.New("snowflake初始化失败")
	InitValidator = errors.New("validator初始化失败")
	GoRoutine     = errors.New("goroutine启动出错")

	DBConnect = errors.New("连接数据库失败")

	InvalidParam = errors.New("填写参数有误")
	InvalidToken = errors.New("错误token")

	SignUp  = errors.New("注册失败")
	Login   = errors.New("登录失败")
	NoRoute = errors.New("页面未定义")

	UserNotExist      = errors.New("用户不存在")
	UserExist         = errors.New("用户已存在")
	UserWrongPassword = errors.New("输入密码错误")
	UserNotLogin      = errors.New("用户未登录")

	CommunityNoData    = errors.New("不存在社区数据")
	CommunityInvalidID = errors.New("不存在该社区")

	PostNotExist      = errors.New("帖子不存在")
	PostCreate        = errors.New("帖子创建失败")
	PostIDGet         = errors.New("帖子id获取失败")
	PostListGet       = errors.New("帖子列表获取失败")
	PostPageGet       = errors.New("帖子分页获取失败")
	PostVote          = errors.New("帖子投票失败")
	PostVoteOverdue   = errors.New("帖子投票时间已过")
	PostTimeNotFound  = errors.New("帖子投票时间查找失败")
	PostVoteNotChange = errors.New("帖子未改变投票")

	ShutDownServer = errors.New("服务器关机")
)
