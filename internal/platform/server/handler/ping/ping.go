package ping

import (
	"net/http"

	"github.com/arkiant/hexagonal-golang-api/internal/sending/ping"
	"github.com/arkiant/hexagonal-golang-api/internal/platform/server"
	"github.com/arkiant/hexagonal-golang-api/kit/cqrs/query"
	"github.com/gin-gonic/gin"
)

func PingHandler(queryBus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := queryBus.Dispatch(ctx, ping.NewPingQuery())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, server.Response{Data: response})
	}
}