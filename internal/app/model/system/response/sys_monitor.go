package response

// Redis基础信息
type RedisInfor struct {
	Redis_version     string `json:"redis_version"`
	Redis_git_sha1    string `json:"redis_git_sha_1"`
	Redis_git_dirty   string `json:"redis_git_dirty"`
	Redis_build_id    string `json:"redis_build_id"`
	Redis_mode        string `json:"redis_mode"`
	Os                string `json:"os"`
	Arch_bits         string `json:"arch_bits"`
	Multiplexing_api  string `json:"multiplexing_api"`
	Process_id        string `json:"process_id"`
	Run_id            string `json:"run_id"`
	Tcp_port          string `json:"tcp_port"`
	Uptime_in_seconds string `json:"uptime_in_seconds"`
	Uptime_in_days    string `json:"uptime_in_days"`
	Hz                string `json:"hz"`
	Lru_clock         string `json:"lru_clock"`
	Executable        string `json:"executable"`
	Config_file       string `json:"config_file"`
}

// 实时状态
type RedisInfoCur struct {
	UsedMemory string `json:"used_memory"`
	DBSize     string `json:"db_size"`
}
