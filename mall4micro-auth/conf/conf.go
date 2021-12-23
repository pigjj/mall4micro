package conf

var LocalSettings = new(YmlLocalConf)
var Settings = new(AuthConf)

func init() {
	loadLocalConf()
	err := Settings.LoadConf()
	if err != nil {
		panic(err)
	}
}
