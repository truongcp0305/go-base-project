package init

import (
	"go-project/adapters/connection"
	"go-project/infrastructure/router"
	"go-project/registry"
	"os"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	userCol := connection.Conn()
	es := connection.EsConn()
	c := registry.New(userCol, es).NewAppController()
	g := gin.Default()
	router.NewRouter(g, c)
	g.Run(os.Getenv("SEVER_HOST"))
}
