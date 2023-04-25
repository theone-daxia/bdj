package env

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/theone-daxia/bdj/framework/contract"
	"io"
	"os"
	"path"
	"strings"
)

type BdjEnvService struct {
	folder string            // 代表 .env 所在的目录
	maps   map[string]string // 保存所有的环境变量
}

func NewBdjEnvService(params ...interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("NewBdjEnvService param error")
	}

	folder := params[0].(string)

	// 实例化
	bdjEnv := &BdjEnvService{
		folder: folder,
		// APP_ENV 默认设置为开发环境
		maps: map[string]string{"APP_ENV": contract.EnvDevelopment},
	}

	file := path.Join(folder, ".env")
	f, err := os.Open(file)
	if err == nil {
		defer f.Close()

		br := bufio.NewReader(f)
		for {
			line, _, err := br.ReadLine()
			if err == io.EOF {
				break
			}
			s := bytes.SplitN(line, []byte{'='}, 2)
			if len(s) < 2 {
				continue
			}
			// 保存 map
			key := string(s[0])
			val := string(s[1])
			bdjEnv.maps[key] = val
		}
	}

	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) < 2 {
			continue
		}
		bdjEnv.maps[pair[0]] = pair[1]
	}

	return bdjEnv, nil
}

func (s *BdjEnvService) AppEnv() string {
	return s.Get("APP_ENV")
}

func (s *BdjEnvService) IsExist(key string) bool {
	_, ok := s.maps[key]
	return ok
}

func (s *BdjEnvService) Get(key string) string {
	if val, ok := s.maps[key]; ok {
		return val
	}
	return ""
}

func (s *BdjEnvService) All() map[string]string {
	return s.maps
}
