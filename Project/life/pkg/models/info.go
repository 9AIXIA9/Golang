package models

type UserInfo struct {
	Username string `json:"username"`
	UserID   int64  `json:"user_id,string"`
	Token    string `json:"token"`
}

type CapInfo struct {
	Name  string `json:"capability_name" db:"capability_name"`   //能力名字
	Score int16  `json:"capability_score" db:"capability_score"` //能力数值
}

// UserCapability 用户能力值结构
type UserCapability struct {
	UserID       int64             `json:"user_id"`
	Capabilities map[string]string `json:"capabilities"` // 使用string类型以保持原始数据格式
}
