package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sheip9/ninelink/pkg/entity"
)

func GetRecord(c *gin.Context) {
	path := c.Param("path")
	var record *entity.Record = nil
	err := (*db).First(&record, "path = ?", path).Error
	if err != nil {
		c.HTML(404, "404.html", gin.H{
			"path": path,
		})
	}
	c.Redirect(302, record.Value)
}
