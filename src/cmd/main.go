package main

import (
	"github.com/iarsham/shop-api/internal"
	"github.com/iarsham/shop-api/internal/common"
)

//	@title						Shop API Document
//	@contact.email				arshamdev2001@gmail.com
//	@BasePath					/api/v1/
//	@host						localhost:8000
//	@licence.Name				MIT
//	@licence.url				https://www.mit.edu/~amini/LICENSE.md
//	@schemes					http https
//	@securityDefinitions.apikey	Authorization
//	@in							header
//	@name						Authorization
//	@accept						json
func main() {
	logs := common.NewLogger()
	defer internal.CloseDB(logs)
	internal.InitialSrv(logs)
}
