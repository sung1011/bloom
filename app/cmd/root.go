package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sung1011/bloom/fw"
)

// RunCommand  初始化根Command并运行
func RunCommand(pot *fw.TicklesPot) error {
	// 根Command
	var rootCmd = &cobra.Command{
		// 定义根命令的关键字
		Use: "bloom",
		// 简短介绍
		Short: "bloom 命令",
		// 根命令的详细介绍
		Long: "bloom 框架提供的命令行工具",
		// 根命令的执行函数
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		// 不需要出现cobra默认的completion子命令
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
	AddAppCommand(rootCmd)
	// 执行RootCommand
	return rootCmd.Execute()
}

// AddAppCommand 绑定业务的命令
func AddAppCommand(rootCmd *cobra.Command) {

	//rootCmd.AddCommand(foo.FooCommand)

	// 每秒调用一次Foo命令
	// rootCmd.AddCronCommand("* * * * * *", foo.FooCommand)

	// 启动一个分布式任务调度，调度的服务名称为init_func_for_test，每个节点每5s调用一次Foo命令，抢占到了调度任务的节点将抢占锁持续挂载2s才释放
	//rootCmd.AddDistributedCronCommand("foo_func_for_test", "*/5 * * * * *", demo.FooCommand, 2*time.Second)
}
