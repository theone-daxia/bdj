package orm

//import (
//	"github.com/theone-daxia/bdj/framework"
//	"github.com/theone-daxia/bdj/framework/contract"
//)
//
//type BdjGormProvider struct {
//}
//
//func (p *BdjGormProvider) Name() string {
//	return contract.ORMKey
//}
//
//func (p *BdjGormProvider) Register(c framework.Container) framework.NewInstance {
//	return NewBdjGormService
//}
//
//func (p *BdjGormProvider) Params(c framework.Container) []interface{} {
//	return []interface{}{c}
//}
//
//func (p *BdjGormProvider) IsDefer() bool {
//	return true
//}
//
//func (p *BdjGormProvider) Boot(c framework.Container) error {
//	return nil
//}
