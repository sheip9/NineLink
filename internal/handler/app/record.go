package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sheip9/ninelink/internal/enum"
	"github.com/sheip9/ninelink/pkg/entity"
	"net/http"
)

func GetRecord(c *gin.Context) {
	path := c.Param("path")
	t, _ := c.Get(enum.RespType)

	var record *entity.Record = nil
	recordErr := (*db).First(&record, "path = ?", path).Error

	switch t {
	case enum.JSON:
		if recordErr != nil {
			c.JSON(http.StatusNotFound, gin.H{})
		}
		c.JSON(http.StatusOK, record)
	default:
		if recordErr != nil {
			c.HTML(http.StatusNotFound, "404.html", gin.H{
				"path": path,
			})
		}
		c.Redirect(http.StatusFound, record.Value)
	}
}
