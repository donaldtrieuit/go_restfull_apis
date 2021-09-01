package dto

/**
 * @author : Donald Trieu
 * @created : 8/25/21, Wednesday
**/
type CreateCameraDto struct {
	Name string `json:"name" form:"name" binding:"required" mapstructure:"name,required"`

	StreamUri string `json:"stream_uri" form:"stream_uri" binding:"required,url" mapstructure:"stream_uri,required"`

	StreamingProtocol string `json:"streaming_protocol" form:"streaming_protocol" binding:"omitempty" mapstructure:"streaming_protocol,omitempty"`

	FFmpegServiceID string `json:"ffmpeg_service_id" form:"ffmpeg_service_id" binding:"omitempty" mapstructure:"ffmpeg_service_id,omitempty"`

	PlaybackServiceID string `json:"playback_service_id" form:"playback_service_id" binding:"omitempty" mapstructure:"playback_service_id,omitempty"`

	FFmpegGlobalServiceID string `json:"ffmpeg_global_service_id" form:"ffmpeg_global_service_id" binding:"omitempty" mapstructure:"ffmpeg_global_service_id,omitempty`

	BoxID int64 `json:"box_id" form:"box_id" binding:"required" mapstructure:"box_id,required"`
}

type UpdateCameraDto struct {
	ID *int64 `json:"id" form:"id" binding:"required" mapstructure:"id,required"`

	Name *string `json:"name" form:"name" binding:"omitempty" mapstructure:"name,omitempty"`

	StreamUri *string `json:"stream_uri" form:"stream_uri" binding:"omitempty,url" mapstructure:"stream_uri,omitempty"`

	StreamEndpoint *string `json:"stream_endpoint" form:"stream_endpoint" binding:"omitempty" mapstructure:"stream_endpoint,omitempty"`

	StreamName string `json:"stream_name" form:"stream_name" binding:"omitempty" mapstructure:"stream_name,omitempty"`

	StreamingProtocol *string `json:"streaming_protocol" form:"streaming_protocol" binding:"omitempty" mapstructure:"streaming_protocol,omitempty"`

	IsAddedStreamfile *int `json:"is_added_streamfile" form:"is_added_streamfile" binding:"omitempty" mapstructure:"is_added_streamfile,omitempty"`

	IsConnectedStreamfile *int `json:"is_connected_streamfile" form:"is_connected_streamfile" binding:"omitempty" mapstructure:"is_connected_streamfile,omitempty"`

	FFmpegServiceID string `json:"ffmpeg_service_id" form:"ffmpeg_service_id" binding:"omitempty" mapstructure:"ffmpeg_service_id,omitempty"`

	PlaybackServiceID *string `json:"playback_service_id" form:"playback_service_id" binding:"omitempty" mapstructure:"playback_service_id,omitempty"`

	FFmpegGlobalServiceID *string `json:"ffmpeg_global_service_id" form:"ffmpeg_global_service_id" binding:"omitempty" mapstructure:"ffmpeg_global_service_id,omitempty"`

	IsConnected *int `json:"is_connected" form:"is_connected" binding:"omitempty" mapstructure:"is_connected,omitempty"`

	Status *string `json:"status" form:"status" mapstructure:"status,omitempty"`

	BoxID *int64 `json:"box_id" form:"box_id" binding:"required" mapstructure:"box_id,required"`
}