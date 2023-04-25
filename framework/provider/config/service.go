package config

import (
	"bytes"
	"errors"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
	"github.com/theone-daxia/bdj/framework"
	"github.com/theone-daxia/bdj/framework/contract"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type BdjConfig struct {
	c         framework.Container    // 容器
	folder    string                 // 文件夹
	delimiter string                 // 路径的分隔符，默认为点
	lock      sync.RWMutex           // 配置文件读写锁
	envMaps   map[string]string      // 所有的环境变量
	confMaps  map[string]interface{} // 配置文件结构，key为文件名
	confRaws  map[string][]byte      // 配置文件的原始信息
}

func NewBdjConfig(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	envFolder := params[1].(string)
	envMaps := params[2].(map[string]string)

	if _, err := os.Stat(envFolder); os.IsNotExist(err) {
		return nil, errors.New("folder " + envFolder + " not exist: " + err.Error())
	}

	// 实例化
	bdjConf := &BdjConfig{
		c:         container,
		folder:    envFolder,
		delimiter: ".",
		lock:      sync.RWMutex{},
		envMaps:   envMaps,
		confMaps:  map[string]interface{}{},
		confRaws:  map[string][]byte{},
	}

	// 读取每个文件
	files, err := os.ReadDir(envFolder)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		fileName := file.Name()
		err := bdjConf.loadConfigFile(envFolder, fileName)
		if err != nil {
			log.Println(err)
			continue
		}
	}

	return bdjConf, nil
}

// replace 使用环境变量 maps 替换 content 中的 env(xxx) 的环境变量
func replace(content []byte, maps map[string]string) []byte {
	if maps == nil || len(maps) <= 0 {
		return content
	}

	for k, v := range maps {
		reKey := "env(" + k + ")"
		content = bytes.ReplaceAll(content, []byte(reKey), []byte(v))
	}
	return content
}

// searchMap 查找某个路径的配置项
func searchMap(source map[string]interface{}, path []string) interface{} {
	if len(path) == 0 {
		return source
	}

	// 判断是否有下个路径
	next, ok := source[path[0]]
	if !ok {
		return nil
	}

	if len(path) == 1 {
		return next
	}
	switch next.(type) {
	case map[interface{}]interface{}:
		return searchMap(cast.ToStringMap(next), path[1:])
	case map[string]interface{}:
		return searchMap(next.(map[string]interface{}), path[1:])
	default:
		return nil
	}
}

func (conf *BdjConfig) loadConfigFile(folder string, file string) error {
	conf.lock.Lock()
	defer conf.lock.Unlock()

	//  判断文件是否以 yaml 或者 yml 作为后缀
	s := strings.Split(file, ".")
	if len(s) == 2 && (s[1] == "yaml" || s[1] == "yml") {
		name := s[0]

		// 读取文件内容
		bf, err := os.ReadFile(filepath.Join(folder, file))
		if err != nil {
			return err
		}

		// 直接针对文本做环境变量的替换
		bf = replace(bf, conf.envMaps)
		// 解析对应的文件
		c := map[string]interface{}{}
		if err := yaml.Unmarshal(bf, &c); err != nil {
			return err
		}
		conf.confMaps[name] = c
		conf.confRaws[name] = bf

		// 读取app.path中的信息，更新app对应的folder
		if name == "app" && conf.c.IsBind(contract.AppKey) {
			if p, ok := c["path"]; ok {
				appService := conf.c.MustMake(contract.AppKey).(contract.App)
				appService.LoadAppConfig(cast.ToStringMapString(p))
			}
		}
	}

	return nil
}

// find 通过path来获取某个配置项
func (conf *BdjConfig) find(key string) interface{} {
	conf.lock.RLock()
	defer conf.lock.RUnlock()
	return searchMap(conf.confMaps, strings.Split(key, conf.delimiter))
}

func (conf *BdjConfig) IsExist(key string) bool {
	return conf.find(key) != nil
}

// Get 获取某个配置项
func (conf *BdjConfig) Get(key string) interface{} {
	return conf.find(key)
}

// GetBool 获取bool类型配置
func (conf *BdjConfig) GetBool(key string) bool {
	return cast.ToBool(conf.find(key))
}

// GetInt 获取int类型配置
func (conf *BdjConfig) GetInt(key string) int {
	return cast.ToInt(conf.find(key))
}

// GetFloat64 获取float64类型配置
func (conf *BdjConfig) GetFloat64(key string) float64 {
	return cast.ToFloat64(conf.find(key))
}

// GetTime 获取time类型配置
func (conf *BdjConfig) GetTime(key string) time.Time {
	return cast.ToTime(conf.find(key))
}

func (conf *BdjConfig) GetString(key string) string {
	return cast.ToString(conf.find(key))
}

func (conf *BdjConfig) GetIntSlice(key string) []int {
	return cast.ToIntSlice(conf.find(key))
}

func (conf *BdjConfig) GetStringSlice(key string) []string {
	return cast.ToStringSlice(conf.find(key))
}

func (conf *BdjConfig) GetStringMap(key string) map[string]interface{} {
	return cast.ToStringMap(conf.find(key))
}

func (conf *BdjConfig) GetStringMapString(key string) map[string]string {
	return cast.ToStringMapString(conf.find(key))
}

func (conf *BdjConfig) GetStringMapStringSlice(key string) map[string][]string {
	return cast.ToStringMapStringSlice(conf.find(key))
}

func (conf *BdjConfig) Load(key string, val interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "yaml",
		Result:  val,
	})
	if err != nil {
		return err
	}
	return decoder.Decode(conf.find(key))
}
