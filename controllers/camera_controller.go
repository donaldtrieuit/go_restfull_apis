package controllers

/**
 * @author : Donald Trieu
 * @created : 8/25/21, Wednesday
**/
import (
	"github.com/gin-gonic/gin"
	"go_restfull_api/dto"
	"go_restfull_api/helpers"
	"go_restfull_api/services"
	"net/http"
	"strconv"
	"strings"
)

type CameraController struct {
	Service *services.CameraService
}

// Get all camera
// @Summary Get all camera
// @Description Get all camera
// @Tags cameras
// @Accept  json
// @Produce json
// @Param box_id query string true "Box ID"
// @Success 200 {object} helpers.Response{data=[]models.Camera}
// @Failure 400,404 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Failure default {object} helpers.Response
// @Router /api/v1/cameras [get]
func (c *CameraController) GetAllCamera(ctx *gin.Context) {
	boxID := ctx.Query("box_id")
	if len(strings.TrimSpace(boxID)) == 0 {
		res := helpers.BuildErrorResponse("Failed to process request", "Param query box_id not found", helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		reqBoxID, _ := strconv.ParseInt(boxID, 10, 64)
		result, err := c.Service.FindAll(reqBoxID)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			response := helpers.BuildResponse(true, "OK", result)
			ctx.JSON(http.StatusOK, response)
		}
	}
}

// Get single camera
// @Summary Get single camera
// @Description Get single camera
// @Tags cameras
// @Accept  json
// @Produce json
// @Param id path int true "Camera ID"
// @Param box_id query string true "Box ID"
// @Success 200 {object} helpers.Response{data=models.Camera}
// @Failure 400,404 {object} helpers.Response{}
// @Failure 500 {object} helpers.Response{}
// @Failure default {object} helpers.Response{}
// @Router /api/v1/cameras/{id} [get]
func (c *CameraController) GetCameraSingle(ctx *gin.Context) {
	id := ctx.Param("id")
	boxID := ctx.Query("box_id")
	if len(strings.TrimSpace(boxID)) == 0 || len(strings.TrimSpace(id)) == 0 {
		res := helpers.BuildErrorResponse("Failed to process request", "Param query box_id or id not found", helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		reqBoxID, _ := strconv.ParseInt(boxID, 10, 64)
		camID, _ := strconv.ParseInt(id, 10, 64)
		result, err := c.Service.FindOneById(camID, reqBoxID)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			response := helpers.BuildResponse(true, "OK", result)
			ctx.JSON(http.StatusOK, response)
		}
	}
}

// Get all camera
// @Summary Get all camera
// @Description Get all camera
// @Tags cameras
// @Accept  json
// @Produce json
// @Param request body dto.CreateCameraDto true "CreateCameraDto"
// @Success 200 {object} helpers.Response{data=models.Camera}
// @Failure 400,404 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Failure default {object} helpers.Response
// @Router /api/v1/cameras [post]
func (c *CameraController) InsertCamera(ctx *gin.Context) {
	var cameraCreateDto dto.CreateCameraDto
	errDTO := ctx.ShouldBind(&cameraCreateDto)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		result, err := c.Service.InsertCamera(cameraCreateDto)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			response := helpers.BuildResponse(true, "OK", result)
			ctx.JSON(http.StatusCreated, response)
		}
	}
}
// Get single camera
// @Summary Get single camera
// @Description Get single camera
// @Tags cameras
// @Accept  json
// @Produce json
// @Param id path int true "Camera ID"
// @Param request body dto.UpdateCameraDto true "UpdateCameraDto"
// @Success 200 {object} helpers.Response{data=models.Camera}
// @Failure 400,404 {object} helpers.Response{}
// @Failure 500 {object} helpers.Response{}
// @Failure default {object} helpers.Response{}
// @Router /api/v1/cameras/{id} [put]
func (c *CameraController) UpdateCamera(ctx *gin.Context) {
	var cameraUpdateDto dto.UpdateCameraDto
	errDTO := ctx.ShouldBind(&cameraUpdateDto)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObject{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		result, err := c.Service.UpdateCamera(*cameraUpdateDto.ID, cameraUpdateDto)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObject{})
			ctx.JSON(http.StatusBadRequest, res)
		} else {
			response := helpers.BuildResponse(true, "OK", result)
			ctx.JSON(http.StatusOK, response)
		}
	}
}
