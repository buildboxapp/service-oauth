package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/foolin/gin-template"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"html/template"
	"os"
)

const (
	authroute         = "/auth/login/auth"
	sendresulttofront = true
)

var (
	vendorslist = []string{
		"google",
		"yandex",
		"vk",
		"mailru",
	}
	authcallback string
	conf         Config
)

func main() {
	// парсим параметры запуска апки
	appOpts = parseAppOpts()
	if appOpts.Test {
		spew.Dump("App opts:", appOpts)
	}

	authcallback = appOpts.Host + authroute

	conf = readConfig()
	setup()

	// инициализируем параметры вендоров аутентификации
	authconfs = vendorsConfInit()

	if !appOpts.Test {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	// подключаем шаблонизатор гина
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "views",
		Extension: ".tpl",
		Funcs: template.FuncMap{
			"tohtml": func(i interface{}) template.HTML {
				return template.HTML(i.(string))
			},
		},
	})

	router.GET("/ping", proxyping)
	router.GET("/help", info)
	router.GET("/", loginform)
	router.GET("/auth", auth)
	router.POST("/auth", auth)
	//router.GET("/getauthresult", reqauthresult)

	// статика в гине
	router.Use(static.Serve("/static", static.LocalFile("static", false)))

	// стартуем апку
	err := router.Run(fmt.Sprintf(":%d", appOpts.Port))
	if err != nil {
		fmt.Println("Error start router:", err)
		os.Exit(1)
	}
}
