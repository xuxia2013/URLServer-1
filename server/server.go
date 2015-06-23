package server

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"URLServer/logs"
	"path/filepath"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"URLServer/models"
	"URLServer/utils"
)
type Config struct {
	DB         map[string][]string `yaml:"db,omitempty"`
}
var Settings Config

func loadConfig() {
	filename, _ := filepath.Abs("./config.yaml")

	yamlFile, err := ioutil.ReadFile(filename)
	utils.Check(err)

	err = yaml.Unmarshal(yamlFile, &Settings)
	utils.Check(err)
}
func Start(){
	logger.Init()
	loadConfig()
	logger.Log.Info("create gin")
	router:=gin.Default()
	models.Eventip=Settings.DB["eventip"][0]
	logger.Log.Info("create postgresql")
	Pq.Create()
	v1:=router.Group("/v1")
	{
	   v1.POST("/webtag",setwebtag)
    	   v1.DELETE("/webtag/:id",delwebtag)
		
	   v1.GET("/webtag",getwebtag)
	   v1.GET("/webtag/:id",getwebtagid)
	}
	logger.Log.Info("linsten 10013")
	http.ListenAndServe(":10013",router)
	defer Pq.Close()
}
