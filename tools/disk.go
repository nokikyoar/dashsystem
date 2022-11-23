package tools

import "strconv"

// GetDiskUsage 获取磁盘使用率
func GetDiskUsage() []float64 {
	// 获取磁盘使用率
	diskUsageResult, _ := ExecCmd("df -TH | grep \"/$\" | awk '{print $6}'")
	// 去除空格、换行符、%
	diskUsageResult = TrimSpace(diskUsageResult)
	// 转换为float64
	diskUsage, _ := strconv.ParseFloat(diskUsageResult, 64)
	// 返回结果
	return []float64{diskUsage}
}
