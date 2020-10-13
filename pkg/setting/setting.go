package setting

import "github.com/spf13/viper"

// 配置文件读取器结构
type Setting struct {
	vp *viper.Viper
}

// 初始化文件读取器
func NewSetting() (*Setting,error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp},nil
}





