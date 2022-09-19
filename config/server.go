/*
 * @Author: kingford
 * @Date: 2022-09-19 20:24:31
 * @LastEditTime: 2022-09-19 20:24:35
 */
package config

type Server struct {
	JWT JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`

	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
