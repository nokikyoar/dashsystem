package tools

import (
	"strconv"
	"strings"
)

// GetCpuUsage 获取CPU使用率,返回一个数组
func GetCpuUsage() []float64 {
	cpuUsageResult, _ := ExecCmd("top -bn 1 | grep 'Cpu(s)' | awk '{print $2 + $4}'")
	// 去除空格、换行符
	cpuUsageResult = TrimSpace(cpuUsageResult)
	// 转换为float64
	cpuUsage, _ := strconv.ParseFloat(cpuUsageResult, 64)
	// 返回结果
	return []float64{cpuUsage}
}

// GetLoad 系统负载
func GetLoad() string {
	loadResult, _ := ExecCmd("uptime | awk -F 'average:' '{print $2}'")
	return "系统负载：" + loadResult

}

// GetUserCpuUsage 用户态CPU使用率
func GetUserCpuUsage() string {
	userCpuUsageResult, _ := ExecCmd("top -b -n1 | grep \"Cpu(s)\" | awk '{print $2 + $6}'")
	// 去除空格、换行符
	userCpuUsageResult = TrimSpace(userCpuUsageResult)
	return "用户态CPU使用率：" + userCpuUsageResult + "%" + "    "
}

// GetSysCpuUsage 系统态CPU使用率
func GetSysCpuUsage() string {
	sysCpuUsageResult, _ := ExecCmd("top -b -n1 | grep \"Cpu(s)\" | awk '{print $4 + $12 + $14+ $16}'")
	// 去除空格、换行符
	sysCpuUsageResult = TrimSpace(sysCpuUsageResult)
	return "系统态CPU使用率：" + sysCpuUsageResult + "%" + "\n"
}

// GetCpuWait CPU 等待率
func GetCpuWait() string {
	cpuWaitResult, _ := ExecCmd("top -b -n1 | grep \"Cpu(s)\" | awk '{print $10}'")
	// 去除空格、换行符
	cpuWaitResult = TrimSpace(cpuWaitResult)
	return "CPU等待率：" + cpuWaitResult + "%" + "    "
}

// GetCpuCore CPU 核数
func GetCpuCore() string {
	cpuCoreResult, _ := ExecCmd("cat /proc/cpuinfo | grep \"cpu cores\" | uniq | awk -F ':' '{print $2}'")
	// 去除空格、换行符
	cpuCoreResult = TrimSpace(cpuCoreResult)
	return "CPU核数：" + cpuCoreResult + "    "
}

// GetCpuPhysical 物理CPU个数
func GetCpuPhysical() string {
	cpuPhysicalResult, _ := ExecCmd("cat /proc/cpuinfo | grep \"physical id\" | sort | uniq | wc -l")
	// 去除空格、换行符
	cpuPhysicalResult = TrimSpace(cpuPhysicalResult)
	return "物理CPU个数：" + cpuPhysicalResult + "    " + "\n"
}

// GetWaitIOProcess 等待 IO进程数
func GetWaitIOProcess() string {
	waitIOProcessResult, _ := ExecCmd("ps -eo state,pid,cmd | grep \"^D\" | wc -l")
	// 去除空格、换行符
	waitIOProcessResult = TrimSpace(waitIOProcessResult)
	return "等待IO进程数：" + waitIOProcessResult + "    "
}

// GetZombieProcess 僵尸进程数
func GetZombieProcess() string {
	zombieProcessResult, _ := ExecCmd("ps -eo state,pid,cmd | grep \"^Z\" | wc -l")
	// 去除空格、换行符
	zombieProcessResult = TrimSpace(zombieProcessResult)
	return "僵尸进程数：" + zombieProcessResult + "\n"
}

// GetTop8CpuProcess top8 CPU 占用进程
func GetTop8CpuProcess() []string {
	top8CpuProcessResult, _ := ExecCmd("ps aux|grep -v PID|sort -rn -k +3| awk '{print $1,$2,$3,$4,$11}'|head -n 8 | column -t")
	// 去除最右边的换行符
	top8CpuProcessResult = strings.TrimRight(top8CpuProcessResult, "\n")

	// 结果按照 \n 分割
	top8CpuProcessSlice := strings.Split(top8CpuProcessResult, "\n")
	// 在最前面插入标题
	top8CpuProcessSlice = append([]string{"用户  PID    CPU  内存  进程名"}, top8CpuProcessSlice...)

	return top8CpuProcessSlice

}
