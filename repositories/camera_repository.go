package repositories

import (
	"go_restfull_api/models"
	"gorm.io/gorm"
)

/**
 * @author : Donald Trieu
 * @created : 8/25/21, Wednesday
**/
type ICameraRepository interface {
	FindAll(boxID int64) (*[]models.Camera, error)
	FindOneById(id int64, boxID int64) (*models.Camera, error)
	InsertCamera(createCamera models.Camera) (*models.Camera, error)
	UpdateCamera(id int64, updateCamera map[string]interface{}) (*models.Camera, error)
}

type CameraRepository struct {
	DB *gorm.DB
}

func (c *CameraRepository) FindAll(boxID int64) (*[]models.Camera, error) {
	var cameras *[]models.Camera
	result := c.DB.Model(&models.Camera{}).Where("box_id = ?", boxID).Find(&cameras)
	if result.Error != nil {
		return nil, result.Error
	}
	return cameras, nil
}

func (c *CameraRepository) FindOneById(id int64, boxID int64) (*models.Camera, error) {
	var camera *models.Camera
	result := c.DB.Model(&models.Camera{}).Where("box_id = ?", boxID).First(&camera, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return camera, nil
}

func (c *CameraRepository) InsertCamera(createCamera models.Camera) (*models.Camera, error) {
	result := c.DB.Model(&models.Camera{}).Create(&createCamera)
	if result.Error != nil {
		return nil, result.Error
	}
	return &createCamera, nil
}

func (c *CameraRepository) UpdateCamera(id int64, updateCamera map[string]interface{}) (*models.Camera, error) {
	result := c.DB.Model(&models.Camera{}).Where("id = ?", id).Updates(&updateCamera)
	if result.Error != nil {
		return nil, result.Error
	}
	cam, _ := c.FindOneById(id, *updateCamera["box_id"].(*int64))
	return cam, nil
}

func (*CameraRepository) BindKey() string {
	return "repo.Camera"
}
