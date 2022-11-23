package tools

import (
	"math"
	"strconv"
	"strings"
)

// GetMemUsage 获取内存使用率
func GetMemUsage() []float64 {
	// 获取内存使用率
	memUsageResult, _ := ExecCmd("free -m | grep Mem | awk '{print 100 - $4/$2*100}'")
	// 去除空格、换行符
	memUsageResult = TrimSpace(memUsageResult)

	// 转换为float64
	memUsage, _ := strconv.ParseFloat(memUsageResult, 64)

	// 提取前俩位小数
	memUsage = math.Floor(memUsage*100) / 100
	// 返回结果
	return []float64{memUsage}
}

// GetMemUsageString 内存使用率，返回值为string类型
func GetMemUsageString() string {
	// 获取内存使用率
	memUsageResult, _ := ExecCmd("free -m | grep Mem | awk '{print 100 - $4/$2*100}'")
	// 去除空格、换行符
	memUsageResult = TrimSpace(memUsageResult)

	// 提取前面的五位数字
	memUsageResult = memUsageResult[:5]

	return "内存使用率：" + memUsageResult + "%" + "    "
}

// GetMemNum 内存条数
func GetMemNum() string {
	// 获取内存条数
	memNumResult, _ := ExecCmd("dmidecode -t 17 | grep -c 'Size: '")
	// 去除空格、换行符
	memNumResult = TrimSpace(memNumResult)
	return "内存条数：" + memNumResult + "    "
}

// GetMemType 内存条类型
func GetMemType() string {
	// 获取内存条类型
	memTypeResult, _ := ExecCmd("dmidecode | grep -A16 \"Memory Device\" | grep 'Type:' |grep -v Unknown |uniq | awk -F ':' '{print $2}'")
	// 去除空格、换行符
	memTypeResult = TrimSpace(memTypeResult)
	return "内存条类型：" + memTypeResult + "\n"
}

// SwapUsage  交换分区使用率
func SwapUsage() string {
	// 获取交换分区使用率
	swapUsageResult, _ := ExecCmd("free -m | grep Swap | awk '{print $3/($2+1)*100}'")
	// 去除空格、换行符
	swapUsageResult = TrimSpace(swapUsageResult)
	return "交换分区使用率：" + swapUsageResult + "%" + "\n"
}

// GetTop8MemProcess 内存 Top8 进程
func GetTop8MemProcess() []string {
	top8MemProcessResult, _ := ExecCmd("ps aux|grep -v PID|sort -rn -k +4| awk '{print $1,$2,$3,$4,$11}'|head -8| column -t")
	// 去除最右边的换行符
	top8MemProcessResult = strings.TrimRight(top8MemProcessResult, "\n")

	// 结果按照 \n 分割
	top8MemProcessSlice := strings.Split(top8MemProcessResult, "\n")
	// 在最前面插入标题
	top8MemProcessSlice = append([]string{"用户    PID    CPU  内存  进程名"}, top8MemProcessSlice...)

	return top8MemProcessSlice

}
