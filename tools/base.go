package tools

// GetNowTime 获取服务器当前时间
func GetNowTime() string {
	dateResult, _ := ExecCmd("date")
	return "服务器当前时间：" + dateResult
}

// GetUptime 获取服务器运行时间
func GetUptime() string {
	uptimeResult, _ := ExecCmd("uptime -p")
	return "服务器运行时间：" + uptimeResult
}

// GetOsVersion 操作系统版本
func GetOsVersion() string {
	osVersionResult, _ := ExecCmd("cat  /etc/redhat-release")
	return "操作系统版本：" + osVersionResult
}

// GetServerModel 服务器型号
func GetServerModel() string {
	serverModelResult, _ := ExecCmd("dmidecode -s system-product-name")
	return "服务器型号：" + serverModelResult
}

// GetServerCpu 服务器CPU信息
func GetServerCpu() string {
	serverCpuResult, _ := ExecCmd("cat /proc/cpuinfo | grep 'model name' | uniq | awk -F: '{print $2}'")
	return "服务器CPU型号：" + serverCpuResult
}

// GetServerNicCount 服务器网卡数量
func GetServerNicCount() string {
	serverNicCountResult, _ := ExecCmd("lspci | grep -i 'ethernet controller' | wc -l")
	return "服务器网卡数量：" + serverNicCountResult
}

// GetHostName 主机名
func GetHostName() string {
	hostnameResult, _ := ExecCmd("hostname -s")
	return "主机名：" + hostnameResult
}

// GetCurrentUser 获取当前用户
func GetCurrentUser() string {
	currentUserResult, _ := ExecCmd("whoami")
	return "当前用户：" + currentUserResult
}
