package settings

// AppConfig 都必须用mapstructure
type AppConfig struct {
	App        *AppInfo `mapstructure:"app"`
	*Auth      `mapstructure:"auth"`
	*LogConf   `mapstructure:"log"`
	*MysqlConf `mapstructure:"mysql"`
	*RedisConf `mapstructure:"redis"`
}

type AppInfo struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
	Language  string `mapstructure:"language"`
	Port      string `mapstructure:"port"`
}

type Auth struct {
	TokenDuration int    `mapstructure:"token_duration"`
	TokenLocation string `mapstructure:"token_location"`
	TokenHeader   string `mapstructure:"token_header"`
	TokenSecret   string `mapstructure:"token_secret"`
}

type LogConf struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConf struct {
	Host               string `mapstructure:"host"`
	User               string `mapstructure:"user"`
	Password           string `mapstructure:"password"`
	DatabaseName       string `mapstructure:"database_name"`
	Port               string `mapstructure:"port"`
	MaxOpenConnections int    `mapstructure:"max_open_connections"`
	MaxIdleConnections int    `mapstructure:"max_idle_connections"`
}

type RedisConf struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Database int    `mapstructure:"database"`
	PoolSize int    `mapstructure:"pool_size"`
}
