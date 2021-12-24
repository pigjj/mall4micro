package conf

var LocalSettings = new(YmlLocalConf)
var Settings = new(MicroConf)

func ReloadConf(microServiceName string) {
	loadLocalConf(microServiceName)
	err := Settings.LoadConf()
	if err != nil {
		panic(err)
	}
}
