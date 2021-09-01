package services

import (
	"github.com/mitchellh/mapstructure"
	"go_restfull_api/dto"
	"go_restfull_api/models"
	"go_restfull_api/repositories"
)

/**
 * @author : Donald Trieu
 * @created : 8/25/21, Wednesday
**/
type CameraService struct {
	CameraRepository repositories.ICameraRepository
}

func (c *CameraService) FindAll(boxID int64) (*[]models.Camera, error) {
	return c.CameraRepository.FindAll(boxID)
}

func (c *CameraService) FindOneById(id int64, boxID int64) (*models.Camera, error) {
	return c.CameraRepository.FindOneById(id, boxID)
}

func (c *CameraService) InsertCamera(createCamera dto.CreateCameraDto) (*models.Camera, error) {
	camera := models.Camera{}
	err := mapstructure.Decode(createCamera, &camera)
	if err != nil {
		return nil, err
	}
	return c.CameraRepository.InsertCamera(camera)
}

func (c *CameraService) UpdateCamera(id int64, cameraUpdateDto dto.UpdateCameraDto) (*models.Camera, error) {
	var camera map[string]interface{}
	err := mapstructure.Decode(cameraUpdateDto, &camera)
	if err != nil {
		return nil, err
	}
	return c.CameraRepository.UpdateCamera(id, camera)
}

func (*CameraService) BindKey() string {
	return "service.Camera"
}
