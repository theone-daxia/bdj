package app

import (
	"errors"
	"github.com/theone-daxia/bdj/framework"
	"github.com/theone-daxia/bdj/framework/util"
	"path/filepath"
)

type BdjApp struct {
	container  framework.Container // 服务容器
	baseFolder string              // 项目根目录
	configMap  map[string]string   // 配置加载
}

func NewBdjApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	c := params[0].(framework.Container)
	folder := params[1].(string)
	return &BdjApp{container: c, baseFolder: folder}, nil
}

func (b BdjApp) Version() string {
	return "0.0.1"
}

// BaseFolder 可代表开发场景目录，也可代表运行时候的目录
func (b BdjApp) BaseFolder() string {
	if b.baseFolder != "" {
		return b.baseFolder
	}

	// 没有设置，则使用命令行参数
	//var baseFolder string
	//flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数，默认为当前路径")
	//flag.Parse()
	//if baseFolder != "" {
	//	return baseFolder
	//}

	// 没有参数，则使用当前路径
	return util.GetExecDirectory()
}

func (b BdjApp) ConfigFolder() string {
	if folder, ok := b.configMap["config_folder"]; ok {
		return folder
	}
	return filepath.Join(b.BaseFolder(), "config")
}

func (b BdjApp) StorageFolder() string {
	if folder, ok := b.configMap["storage_folder"]; ok {
		return folder
	}
	return filepath.Join(b.BaseFolder(), "storage")
}

func (b BdjApp) LogFolder() string {
	if folder, ok := b.configMap["log_folder"]; ok {
		return folder
	}
	return filepath.Join(b.StorageFolder(), "log")
}

func (b BdjApp) AppFolder() string {
	return filepath.Join(b.BaseFolder(), "app")
}

func (b BdjApp) HttpFolder() string {
	if folder, ok := b.configMap["http_folder"]; ok {
		return folder
	}
	return filepath.Join(b.AppFolder(), "http")
}

func (b BdjApp) ProviderFolder() string {
	if folder, ok := b.configMap["provider_folder"]; ok {
		return folder
	}
	return filepath.Join(b.AppFolder(), "provider")
}

func (b BdjApp) MiddlewareFolder() string {
	if folder, ok := b.configMap["middleware_folder"]; ok {
		return folder
	}
	return filepath.Join(b.HttpFolder(), "middleware")
}

func (b BdjApp) ConsoleFolder() string {
	if folder, ok := b.configMap["console_folder"]; ok {
		return folder
	}
	return filepath.Join(b.AppFolder(), "console")
}

func (b BdjApp) CommandFolder() string {
	if folder, ok := b.configMap["command_folder"]; ok {
		return folder
	}
	return filepath.Join(b.ConsoleFolder(), "command")
}

func (b BdjApp) RuntimeFolder() string {
	if folder, ok := b.configMap["runtime_folder"]; ok {
		return folder
	}
	return filepath.Join(b.StorageFolder(), "runtime")
}

func (b BdjApp) TestFolder() string {
	if folder, ok := b.configMap["test_folder"]; ok {
		return folder
	}
	return filepath.Join(b.BaseFolder(), "test")
}

func (b BdjApp) LoadAppConfig(kv map[string]string) {
	for k, v := range kv {
		b.configMap[k] = v
	}
}
