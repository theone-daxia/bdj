package console

import (
	"github.com/theone-daxia/bdj/framework"
	"github.com/theone-daxia/bdj/framework/cobra"
	"github.com/theone-daxia/bdj/framework/command"
)

// RunCommand 初始化根 Command 并运行
func RunCommand(container framework.Container) error {
	var rootCmd = &cobra.Command{
		// 根命令的关键字
		Use:   "bdj",
		Short: "bdj 命令",
		Long:  "bdj 框架提供的命令行工具，使用此工具可方便的执行框架自带命令，也能方便的编写业务命令",
		// 根命令的执行函数
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		// 不需要出现 cobra 默认的 completion 子命令
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	// 为根命令设置服务容器
	rootCmd.SetContainer(container)
	// 挂载框架命令
	command.AddKernelCommands(rootCmd)
	// 挂载业务命令
	AddAppCommands(rootCmd)
	// 执行根命令
	return rootCmd.Execute()
}

func AddAppCommands(rootCmd *cobra.Command) {

}
