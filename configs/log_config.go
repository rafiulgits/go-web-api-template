package configs

type LogConfig struct {
	FilePath           string `json:"filePath"`
	MaxAgeInHour       int    `json:"maxAgeInHour"`
	RotationTimeInHour int    `json:"rotationTimeInHour"`
}
