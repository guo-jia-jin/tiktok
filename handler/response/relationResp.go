package response

import (
	"Tiktok/service/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RelationResponse struct {
	Response
	Data []server.RelationPageUser `json:"user_list"`
}

func OkWithDataforRelation(msg string, data *[]server.RelationPageUser, c *gin.Context) {
	c.JSON(http.StatusOK, RelationResponse{
		Response: Response{SUCCESS, msg},
		Data:     *data,
	})
}
