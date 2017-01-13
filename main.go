package main
import (
	"net/http"
	"strings"
	"gopkg.in/gin-gonic/gin"
	"io/ioutil"
	"encoding/json"
)

var router *gin.Engine

type Request struct {
	Site []string `json:"site"`
	SearchText string `json:"searchString"`
}

type Response struct {
	FoundAtSite string
}

func main() {

	router = gin.Default();

	router.POST("/",func(c *gin.Context){
		req := Request{}
		body, _ := ioutil.ReadAll(c.Request.Body)
		json.Unmarshal(body, &req)
		sites:=req.Site;
		searchText:=req.SearchText
		var match = false;
		var index = 0;
		for i:=range sites {
			page,err := http.Get(sites[i]);
			if err == nil{
				body,error := ioutil.ReadAll(page.Body);
				if error == nil{
					if strings.Contains(string(body),searchText){
						match = true;
						index = i;
					}
				}
			}
		}
		if match{
			c.JSON(http.StatusOK,&Response{
				FoundAtSite: sites[index],
			})
		}else {
			c.Writer.WriteHeader(http.StatusNoContent)
		}
	})

	router.Run()

}