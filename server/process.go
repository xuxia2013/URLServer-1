package server

import(
	"github.com/gin-gonic/gin"
	"URLServer/models"
	"fmt"
)
var (
	Pq Postgresql 
	Url=models.URLLIST{}
	UrlList=models.URLLISTS{}
	EventInfo =models.EVENTINFO{}
)
func setwebtag (c *gin.Context){
	Url.Init()
	c.Bind(&Url)
	if Url.Check(){
		UrlList.Insert(Url)
		EventInfo.InitAndPost("urlweb",Url.ID)
	//	EventInfo.InitAndPost("urlweb",Url.URL)
		fmt.Println(UrlList)	
		c.JSON(200,Url.ID)
	}else{
		c.JSON(404,gin.H{"status":"false"})
	}
	
}
func delwebtag (c *gin.Context){
	delid:=c.Params.ByName("id")
	if  delid==""{
	  c.JSON(404,gin.H{"status":"nil"})
	}else if delid=="all"{
		UrlList.Clean()
		c.JSON(200,gin.H{"status":"true"})
		fmt.Println(UrlList)
	}else{
		UrlList.Delete(delid)
		c.JSON(200,gin.H{"status":"ok"})
		fmt.Println(UrlList)
	}
}
func getwebtag (c *gin.Context){
	Url=UrlList.Get()
	if Url.Check(){
		c.JSON(200,Url)	
		fmt.Println(UrlList)
	}else{
		c.JSON(404,gin.H{"status":"false"})
	  
	}
}
func getwebtagid (c *gin.Context){
	getid:=c.Params.ByName("id")
	if getid==""{
		c.JSON(404,gin.H{"status":"nil"})
	}else {
		Url=UrlList.GetId(getid)
		Url.Check()
		c.JSON(200,Url)
		fmt.Println(Url)

	}
}
