package orm

//import (
//	"github.com/go-sql-driver/mysql"
//	"github.com/theone-daxia/bdj/framework"
//	"github.com/theone-daxia/bdj/framework/contract"
//	"sync"
//)
//
//type BdjGormService struct {
//	container framework.Container
//	dbs       map[string]*gorm.DB
//	lock      *sync.RWMutex
//}
//
//func NewBdjGormService(params ...interface{}) (interface{}, error) {
//	container := params[0].(framework.Container)
//	dbs := make(map[string]*gorm.DB)
//	lock := &sync.RWMutex{}
//	return &BdjGormService{
//		container: container,
//		dbs:       dbs,
//		lock:      lock,
//	}, nil
//}
//
//func (s *BdjGormService) GetDB(opts ...contract.DBOption) (*gorm.DB, error) {
//	logger := s.container.MustMake(contract.LogKey).(contract.Log)
//
//	// 读取默认配置
//	config := GetBaseConfig(s.container)
//
//	logService := s.container.MustMake(contract.LogKey).(contract.Log)
//
//	// 设置Logger
//	ormLogger := NewOrmLogger(logService)
//	config.Config = &gorm.Config{
//		Logger: ormLogger,
//	}
//
//	for _, opt := range opts {
//		if err := opt(s.container, config); err != nil {
//			return nil, err
//		}
//	}
//
//	// 如果最终的 config 没有设置 dsn，就生成 dsn
//	if config.Dsn == "" {
//		dsn, err := config.FormatDsn()
//		if err != nil {
//			return nil, err
//		}
//		config.Dsn = dsn
//	}
//
//	// 判断是否已经实例化了gorm.DB
//	s.lock.RLock()
//	if db, ok := s.dbs[config.Dsn]; ok {
//		s.lock.RUnlock()
//		return db, nil
//	}
//	s.lock.RUnlock()
//
//	// 没有实例化 gorm.DB，那么就要进行实例化操作
//	s.lock.Lock()
//	defer s.lock.Unlock()
//
//	// 实例化gorm.DB
//	var db *gorm.DB
//	var err error
//	switch config.Driver {
//	case "mysql":
//		db, err = gorm.Open(mysql.Open(config.Dsn), config)
//	case "postgres":
//		db, err = gorm.Open(postgres.Open(config.Dsn), config)
//	case "sqlite":
//		db, err = gorm.Open(sqlite.Open(config.Dsn), config)
//	case "sqlserver":
//		db, err = gorm.Open(sqlserver.Open(config.Dsn), config)
//	case "clickhouse":
//		db, err = gorm.Open(clickhouse.Open(config.Dsn), config)
//	}
//}
