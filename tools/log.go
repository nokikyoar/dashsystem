package tools

// GetLastTenLog 系统日志最后10行检查
func GetLastTenLog() string {
	cmd, err := ExecCmd("dmesg | tail -n 10")
	if err != nil {
		return "获取系统日志失败"
	}
	return cmd
}
