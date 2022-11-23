/*
Copyright © 2022 kikyoar

*/
package cmd

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/spf13/cobra"
	"log"
	"time"
	"yuhao.com/dashSystem/tools"
)

// ntopCmd represents the ntop command
var ntopCmd = &cobra.Command{
	Use:   "ntop",
	Short: "获取系统信息图表",
	Long:  `获取系统信息图表.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := ui.Init(); err != nil {
			log.Fatalf("failed to initialize termui: %v", err)
		}
		defer ui.Close()
		/*
			基础信息
		*/
		// 文本框
		base := widgets.NewParagraph()
		base.Title = "基础信息"
		base.TitleStyle.Fg = ui.ColorRed
		base.SetRect(0, 0, 80, 10)
		base.Text = tools.GetHostName() + tools.GetCurrentUser() + tools.GetNowTime() +
			tools.GetUptime() + tools.GetOsVersion() + tools.GetServerModel() + tools.GetServerCpu() +
			tools.GetServerNicCount()

		/*
			重点指标信息
		*/
		// 文本框
		metrics := widgets.NewParagraph()
		metrics.Title = "重点指标信息"
		metrics.TitleStyle.Fg = ui.ColorRed
		metrics.SetRect(144, 0, 64, 10)
		metrics.Text = tools.GetLoad() + tools.GetUserCpuUsage() + tools.GetSysCpuUsage() + tools.GetCpuWait() +
			tools.GetCpuCore() + tools.GetCpuPhysical() +
			tools.GetMemUsageString() + tools.GetMemNum() + tools.GetMemType() +
			tools.GetWaitIOProcess() + tools.GetZombieProcess() + tools.GetTcpConn() +
			tools.SwapUsage() + tools.GetNetIn() + tools.GetNetOut()

		/*
		 CPU、内存、硬盘信息检查
		*/
		// CPU
		cpu := widgets.NewPlot()
		cpu.Title = "CPU 使用率"
		cpu.TitleStyle.Fg = ui.ColorRed
		cpu.SetRect(0, 10, 48, 20)
		cpu.Marker = widgets.MarkerDot
		cpu.Data = [][]float64{tools.GetCpuUsage()}
		cpu.AxesColor = ui.ColorWhite
		cpu.LineColors[0] = ui.ColorYellow
		cpu.DrawDirection = widgets.DrawLeft

		// 硬盘
		disk := widgets.NewPlot()
		disk.Title = "硬盘根目录使用率"
		disk.TitleStyle.Fg = ui.ColorRed
		disk.SetRect(48, 10, 96, 20)
		disk.Marker = widgets.MarkerDot
		disk.Data = make([][]float64, 1)
		disk.Data[0] = tools.GetDiskUsage()
		disk.AxesColor = ui.ColorWhite
		disk.DotMarkerRune = '*'
		disk.LineColors[0] = ui.ColorYellow
		disk.PlotType = widgets.ScatterPlot
		disk.Border = true
		disk.BorderStyle.Fg = ui.ColorBlue

		// 内存
		mem := widgets.NewPlot()
		mem.Title = "内存使用率"
		mem.TitleStyle.Fg = ui.ColorRed
		mem.SetRect(96, 10, 144, 20)
		mem.Marker = widgets.MarkerDot
		mem.Data = make([][]float64, 1)
		mem.Data[0] = tools.GetMemUsage()
		mem.AxesColor = ui.ColorWhite
		mem.LineColors[0] = ui.ColorYellow
		mem.PlotType = widgets.ScatterPlot

		// 根据余数显示不同颜色
		updateParagraph := func(count int) {
			base.Text = tools.GetHostName() + tools.GetCurrentUser() + tools.GetNowTime() +
				tools.GetUptime() + tools.GetOsVersion() + tools.GetServerModel() + tools.GetServerCpu() +
				tools.GetServerNicCount()

			metrics.Text = tools.GetLoad() + tools.GetUserCpuUsage() + tools.GetSysCpuUsage() + tools.GetCpuWait() +
				tools.GetCpuCore() + tools.GetCpuPhysical() +
				tools.GetMemUsageString() + tools.GetMemNum() + tools.GetMemType() +
				tools.GetWaitIOProcess() + tools.GetZombieProcess() + tools.GetTcpConn() +
				tools.SwapUsage() + tools.GetNetIn() + tools.GetNetOut()

			if count%2 == 0 {
				base.TextStyle.Fg = ui.ColorWhite
				metrics.TextStyle.Fg = ui.ColorWhite
			} else {
				base.TextStyle.Fg = ui.ColorYellow
				metrics.TextStyle.Fg = ui.ColorGreen
			}
		}
		/*
		 CPU TOP5 进程信息
		*/
		cpuProcess := widgets.NewList()
		cpuProcess.Title = "CPU TOP8 进程"
		cpuProcess.TitleStyle.Fg = ui.ColorRed
		//process.Rows = []string{"USER      PID    %CPU  %MEM  COMMAND",
		//	"root  6569   4.0  2.0  /bin/prometheus",
		//	"root  19013  3.9  0.4  /usr/bin/kubelet",
		//	"root  11478  3.4  7.9  /opt/jdk1.8.0_211/bin/java",
		//	"hive  15858  2.9  3.4  /opt/jdk1.8.0_211/bin/java",
		//	"root  2979   1.9  0.5  /usr/bin/python",}
		cpuProcess.Rows = tools.GetTop8CpuProcess()
		cpuProcess.TextStyle = ui.NewStyle(ui.ColorYellow)
		cpuProcess.WrapText = false
		cpuProcess.Border = true
		cpuProcess.BorderStyle.Fg = ui.ColorGreen
		cpuProcess.SetRect(0, 20, 72, 31)

		/*
		 内存 TOP5 进程信息
		*/
		memProcess := widgets.NewList()
		memProcess.Title = "内存 TOP8 进程"
		memProcess.TitleStyle.Fg = ui.ColorRed
		memProcess.Rows = tools.GetTop8MemProcess()
		memProcess.TextStyle = ui.NewStyle(ui.ColorYellow)
		memProcess.WrapText = true
		memProcess.SetRect(72, 20, 144, 31)
		memProcess.Border = true
		memProcess.BorderStyle.Fg = ui.ColorYellow

		/*
			日志检查
		*/
		logs := widgets.NewParagraph()
		logs.Title = "日志检查"
		logs.TitleStyle.Fg = ui.ColorRed
		logs.Text = tools.GetLastTenLog()
		logs.TextStyle = ui.NewStyle(ui.ColorYellow)
		logs.WrapText = true
		logs.SetRect(0, 31, 144, 50)
		logs.Border = true
		logs.BorderStyle.Fg = ui.ColorBlue

		// 刷新指标结果
		updateResult := func(count int) {
			// cpu
			cpu.Data[0] = append(cpu.Data[0], tools.GetCpuUsage()[0])
			// 硬盘
			disk.Data[0] = append(disk.Data[0], tools.GetDiskUsage()[0])
			// 内存
			mem.Data[0] = append(mem.Data[0], tools.GetMemUsage()[0])
			// 系统文本指标
			metrics.Text = tools.GetLoad() + tools.GetUserCpuUsage() + tools.GetSysCpuUsage() + tools.GetCpuWait() +
				tools.GetCpuCore() + tools.GetCpuPhysical() +
				tools.GetMemUsageString() + tools.GetMemNum() + tools.GetMemType() +
				tools.GetWaitIOProcess() + tools.GetZombieProcess() + tools.GetTcpConn() +
				tools.SwapUsage() + tools.GetNetIn() + tools.GetNetOut()
			// CPU TOP5 进程
			cpuProcess.Rows = tools.GetTop8CpuProcess()
			// 内存 TOP5 进程
			memProcess.Rows = tools.GetTop8MemProcess()
			// 日志检查
			logs.Text = tools.GetLastTenLog()

			if count > 40 {
				// 输出前count-1个数据
				cpu.Data[0] = cpu.Data[0][1:]
				disk.Data[0] = disk.Data[0][1:]
				mem.Data[0] = mem.Data[0][1:]
			}

		}

		draw := func(count int) {

			ui.Render(base, cpu, disk, mem, metrics, cpuProcess, memProcess, logs)
		}

		tickerCount := 1
		draw(tickerCount)
		tickerCount++
		uiEvents := ui.PollEvents()
		ticker := time.NewTicker(time.Second * 3).C
		for {
			select {
			case e := <-uiEvents:
				switch e.ID {
				case "q", "<C-c>":
					return
				}
			case <-ticker:
				// 更新基础信息
				updateParagraph(tickerCount)
				// 更新CPU、内存、进程信息
				updateResult(tickerCount)

				draw(tickerCount)

				tickerCount++
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(ntopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ntopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ntopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
