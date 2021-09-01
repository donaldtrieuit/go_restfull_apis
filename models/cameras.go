package models

import (
	"gorm.io/gorm"
	"time"
)

/**
 * @author : Donald Trieu
 * @created : 8/25/21, Wednesday
**/

type Camera struct {
	ID int64 `gorm:"primary_key:auto_increment" json:"id"`

	TenantID int64 `gorm:"not null;default:1" json:"tenant_id"`

	FFmpegServiceID string `gorm:"type:varchar(255);column:ffmpeg_service_id" json:"ffmpeg_service_id"`

	PlaybackServiceID string `gorm:"type:varchar(255);column:playback_service_id" json:"playback_service_id"`

	FFmpegGlobalServiceID string `gorm:"type:varchar(255);column:ffmpeg_global_service_id" json:"ffmpeg_global_service_id"`

	BoxID int64 `gorm:"not null" json:"box_id"`

	LocationID int64 `gorm:"not null" json:"location_id"`

	Name string `gorm:"type:varchar(255);not null" json:"name"`

	Driver string `gorm:"type:varchar(255);default:'f-edge'" json:"driver"`

	VideoCodingStandard string `gorm:"type:char(10);default:'h.264'" json:"video_coding_standard"`

	StreamingProtocol string `gorm:"type:enum('rtsp', 'rtmp');default:'rtsp'" json:"streaming_protocol"`

	StreamUri string `gorm:"type:varchar(255);not null" json:"stream_uri"`

	StreamEndpoint string `gorm:"type:varchar(255);not null" json:"stream_endpoint"`

	StreamName string `gorm:"type:varchar(255);not null" json:"stream_name"`

	IsRecordingSet int `gorm:"default:0" json:"is_recording_set"`

	IsPTZEnabled int `gorm:"default:0" json:"is_PTZ_enabled"`

	IsConnected int `gorm:"default:0" json:"is_connected"`

	Status string `gorm:"default:'disconnected'" json:"status"`

	DeletedAt *gorm.DeletedAt `gorm:"default:null" json:"deleted_at"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Camera) TableName() string {
	return "devices"
}