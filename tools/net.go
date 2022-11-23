package tools

// GetTcpConn 获取TCP连接数
func GetTcpConn() string {
	// 获取TCP连接数
	tcpConnResult, _ := ExecCmd("ss -s | grep -i \"TCP:\" | awk -F ':' '{print $2}'")
	// 去除前后的空格、换行符
	tcpConnResult = TrimSpace(tcpConnResult)
	// 返回结果
	return "TCP连接数：" + tcpConnResult + "\n"
}

// GetNetIn 网卡总接收字节数
func GetNetIn() string {
	// 获取网卡总接收字节数
	netInResult, _ := ExecCmd("cat /proc/net/dev | awk 'NR>2'| awk '{print NR\" \"$2}' | awk '{print $2}' |  awk '{sum+=$1} END {print sum}'")
	// 去除前后的空格、换行符
	netInResult = TrimSpace(netInResult)

	// bytes 转换为 MB
	netInResult = BytesToGB(netInResult)
	// 返回结果
	return "网卡总接收字节数：" + netInResult + " GB" + "    "
}

// GetNetOut 网卡总发送字节数
func GetNetOut() string {
	// 获取网卡总发送字节数
	netOutResult, _ := ExecCmd("cat /proc/net/dev | awk 'NR>2'| awk '{print NR\" \"$10}' | awk '{print $2}' |  awk '{sum+=$1} END {print sum}'")
	// 去除前后的空格、换行符
	netOutResult = TrimSpace(netOutResult)

	// bytes 转换为 MB
	netOutResult = BytesToGB(netOutResult)
	// 返回结果
	return "网卡总发送字节数：" + netOutResult + " GB" + "\n"
}
