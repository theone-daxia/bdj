package command

import "github.com/theone-daxia/bdj/framework/cobra"

func AddKernelCommands(rootCmd *cobra.Command) {
	// 挂载 appCmd 命令
	rootCmd.AddCommand(initAppCommand())
}
