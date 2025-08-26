package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sheip9/ninelink/internal/database"
	"net/http"
)

func GetRecord(c *gin.Context) {
	path := c.Param("path")
	accept := c.NegotiateFormat(gin.MIMEHTML, gin.MIMEJSON, gin.MIMEXML)

	var record database.Record

	recordErr := database.GetDB().First(&record, "path = ?", path).Error

	switch accept {
	case gin.MIMEJSON:
		if recordErr != nil {
			c.JSON(http.StatusNotFound, gin.H{})
		} else {
			c.JSON(http.StatusOK, record)
		}
	case gin.MIMEXML:
		if recordErr != nil {
			c.XML(http.StatusNotFound, gin.H{})
		} else {
			c.XML(http.StatusOK, record)
		}
	default:
		if recordErr != nil {
			c.HTML(http.StatusNotFound, "404.html", gin.H{
				"path": path,
			})
		} else {
			c.Redirect(http.StatusFound, record.Value)
		}
	}

}
