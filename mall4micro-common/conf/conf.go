package conf

var LocalSettings = new(YmlLocalConf)
var Settings = new(MicroConf)

//
// ReloadConf
// @Description: 从consul加载对应模块的配置
// @param microServiceName
//
func ReloadConf(microServiceName string) {
	loadLocalConf(microServiceName)
	err := Settings.LoadConf()
	if err != nil {
		panic(err)
	}
}
