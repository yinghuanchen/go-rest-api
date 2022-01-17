package main

import (
	"fmt"
	"net/http"
	"go-rest-api/router"
	"go-rest-api/controller"
	"os"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
	postController controller.PostController = controller.NewPostController()
)

const (
	firebaseCred = "/Users/emily.chen/go/src/go-rest-api/golang-project-6de67-firebase-adminsdk-qg7fa-fd0ec7786a.json"
)

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", firebaseCred)
	const port string = ":8000"
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and Running...")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(port)
}