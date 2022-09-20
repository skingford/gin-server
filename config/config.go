/*
 * @Author: kingford
 * @Date: 2022-09-19 20:24:31
 * @LastEditTime: 2022-09-20 10:07:57
 */
package config

type Config struct {
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	// auto
	AutoCode AutoCode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
