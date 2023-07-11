package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	musicalbum "github.com/musicLibrary/controller/musicAlbum"
)

func GetAlbum(ctx *gin.Context){
	albumId,err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(500,map[string]string{
			"message":"something went wrong",
		})
		ctx.Abort()
		return
	}

	if albumId == 0  {
		ctx.JSON(400,map[string]string{
			"message":"bad request",
		})
		ctx.Abort()
		return
	}

	albumData,err := musicalbum.GetAlbum(uint(albumId))
	if err != nil {
		ctx.JSON(500,map[string]string{
			"message":err.Error(),
		})
		ctx.Abort()
		return
	}

	// if albumData.ID == 0 {
	// 	ctx.JSON(404,map[string]string{
	// 		"message":"album not found",
	// 	})
	// 	ctx.Abort()
	// 	return
	// }
	ctx.JSON(200,albumData)
}

func CreateAlbum(ctx *gin.Context){
	var albumData musicalbum.CreateMusicAlbum
	
	err := ctx.ShouldBindJSON(&albumData)
	if err != nil {
		ctx.JSON(500,map[string]string{
			"message":err.Error(),
		})
		ctx.Abort()
		return
	}

	id,err := musicalbum.CreateAlbum(albumData)
	if err != nil {
		ctx.JSON(500,map[string]string{
			"message":err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(200,map[string]uint{
		"album_id":id,
	})
}

func UpdateAlbum(ctx *gin.Context){

}

func DeleteAlbum(ctx *gin.Context){

}