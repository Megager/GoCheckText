package main
import (
	"net/http"
	"strings"
	"gopkg.in/gin-gonic/gin"
	"io/ioutil"
)

var router *gin.Engine

type Request struct {
	Site []string
	SearchText string
}

type Response struct {
	FoundAtSite string
}

func main() {

	router = gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)

	})

	router.POST("/",func(c *gin.Context){
		searchText := c.PostForm("searchText");
		sites := strings.Split(c.PostForm("sites"),";");
		/*json := &Request{
			Site:   sites,
			SearchText: searchText};*/
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
			c.AbortWithStatus(204);
		}
	})

	router.Run()

}