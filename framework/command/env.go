package command

import (
	"fmt"
	"github.com/theone-daxia/bdj/framework/cobra"
	"github.com/theone-daxia/bdj/framework/contract"
	"github.com/theone-daxia/bdj/framework/util"
)

// initEnvCommand 获取env相关的命令
func initEnvCommand() *cobra.Command {
	envCmd.AddCommand(envListCmd)
	return envCmd
}

// envCmd 获取当前的App环境
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "获取当前的App环境",
	Run: func(cmd *cobra.Command, args []string) {
		container := cmd.GetContainer()
		env := container.MustMake(contract.EnvKey).(contract.Env)
		fmt.Println("environment:", env.AppEnv())
	},
}

// envListCmd 获取所有的环境变量
var envListCmd = &cobra.Command{
	Use:   "list",
	Short: "获取所有的环境变量",
	Run: func(cmd *cobra.Command, args []string) {
		container := cmd.GetContainer()
		env := container.MustMake(contract.EnvKey).(contract.Env)
		envs := env.All()
		outs := [][]string{}
		for k, v := range envs {
			outs = append(outs, []string{k, v})
		}
		util.PrettyPrint(outs)
	},
}
