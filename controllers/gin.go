package controllers

import (
	"errors"
	"net/http"
	"project/database"

	"github.com/gin-gonic/gin"
)

func ReadPreman(c *gin.Context) {
	var preman database.Preman
	id := c.Param("id")
	res := database.DB.Find(&preman, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "preman not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"preman": preman,
	})
	return
}

func ReadPremans(c *gin.Context) {
	var premans []database.Preman
	res := database.DB.Find(&premans)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("premans not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": premans,
	})
	return
}

func CreatePreman(c *gin.Context) {
	var preman *database.Preman
	err := c.ShouldBind(&preman)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := database.DB.Create(preman)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a book",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"preman": preman,
	})
	return
 }

 func UpdatePreman(c *gin.Context) {
	var preman database.Preman
	id := c.Param("id")
	err := c.ShouldBind(&preman)
   
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
  
	var updatePreman database.Preman
	res := database.DB.Model(&updatePreman).Where("id = ?", id).Updates(preman)
  
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "preman not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"preman": preman,
	})
	return
 }

 func DeletePreman(c *gin.Context) {
	var preman database.Preman
	id := c.Param("id")
	res := database.DB.Find(&preman, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "preman not found",
		})
		return
	}
	database.DB.Delete(&preman)
	c.JSON(http.StatusOK, gin.H{
		"message": "preman deleted successfully",
	})
	return
 }
 
 