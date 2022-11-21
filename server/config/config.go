package config

import "ve-blog-go/server/config/properties"

type Server struct {
	System properties.System `mapstructure:"system" json:"system" yaml:"system"`
	// gorm
	MysqlConfig properties.Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis       properties.Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	JWT         properties.JWT   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap         properties.Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
	Email       properties.Email `mapstructure:"email" json:"email" yaml:"email"`
	// oss
	//Local      properties.Local      `mapstructure:"local" json:"local" yaml:"local"`
	//TencentCOS properties.TencentCOS `mapstructure:"tencent-cos" json:"tencent-cos" yaml:"tencent-cos"`
	//AwsS3      properties.AwsS3      `mapstructure:"aws-s3" json:"aws-s3" yaml:"aws-s3"`
	//
	//Excel properties.Excel `mapstructure:"excel" json:"excel" yaml:"excel"`
	//Timer properties.Timer `mapstructure:"timer" json:"timer" yaml:"timer"`
}
