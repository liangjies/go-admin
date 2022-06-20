package response

// Redis基础信息
type RedisInfor struct {
	Redis_version            string `json:"redis_version"`            // Redis版本
	Redis_mode               string `json:"redis_mode"`               // Redis模式
	Tcp_port                 string `json:"tcp_port"`                 // Redis TCP端口
	Uptime_in_days           string `json:"uptime_in_days"`           // Redis运行天数
	UsedMemory               string `json:"used_memory"`              // 内存信息
	DBSize                   string `json:"db_size"`                  // Key数量
	Connected_clients        string `json:"connected_clients"`        // 当前连接数
	Total_commands_processed string `json:"total_commands_processed"` // 已执行命令数
}

// 实时状态
type RedisInfoCur struct {
	UsedMemory string `json:"used_memory"`
	DBSize     string `json:"db_size"`
}
