package app

import (
	"NineLink/pkg/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRecord(c *gin.Context) {
	path := c.Param("path")
	var record *entity.Record = nil
	err := db.First(&record, "path = ?", path).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found"})
	}
	c.Redirect(302, record.Value)
}
