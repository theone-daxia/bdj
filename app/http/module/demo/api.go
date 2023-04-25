package demo

import (
	demoService "github.com/theone-daxia/bdj/app/provider/demo"
	"github.com/theone-daxia/bdj/framework/contract"
	"github.com/theone-daxia/bdj/framework/gin"
)

type DemoApi struct {
	service *Service
}

func Register(r *gin.Engine) error {
	api := NewDemoApi()
	//r.Bind(&demoService.DemoServiceProvider{})

	r.GET("demo/demo", api.Demo)
	r.GET("demo/demo2", api.Demo2)
	r.POST("demo/demo_post", api.DemoPost)
	r.GET("demo/env", api.Env)

	return nil
}

func NewDemoApi() *DemoApi {
	service := NewService()
	return &DemoApi{service: service}
}

func (api *DemoApi) Demo(c *gin.Context) {
	users := api.service.GetUsers()
	usersDTO := UserModelsToUserDTOs(users)
	c.JSON(200, usersDTO)
}

func (api *DemoApi) Demo2(c *gin.Context) {
	demoProvider := c.MustMake(demoService.Key).(demoService.IService)
	students := demoProvider.GetAllStudent()
	usersDTOs := StudentsToUserDTOs(students)
	c.JSON(200, usersDTOs)
}

func (api *DemoApi) DemoPost(c *gin.Context) {
	type Foo struct {
		Name string
	}
	foo := &Foo{}
	err := c.BindJSON(&foo)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, nil)
}

func (api *DemoApi) Env(c *gin.Context) {
	configService := c.MustMake(contract.ConfigKey).(contract.Config)
	password := configService.GetString("database.mysql.password")
	c.JSON(200, password)
}
