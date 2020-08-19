package path

import (
	"net/http"
	"strings"

	"github.com/Slahser/coup-de-grace/pkg/ttserver/component"
	"github.com/gin-gonic/gin"
)

func TtWeb(context *gin.Context) {
	var apiHeader component.ApiHeader

	if err := context.BindHeader(&apiHeader); err != nil {
		context.JSON(200, err)
	}

	funcName := context.Param("funcName")

	fullPathStrings := []string{apiHeader.Org, apiHeader.Project, apiHeader.Env, funcName}

	context.String(http.StatusOK, "your code in "+strings.Join(fullPathStrings, "-")+" is blabala")
}
