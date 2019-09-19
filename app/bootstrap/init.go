/*
@Time : 2019/9/19 17:15
@Author : zxr
@File : init
@Software: GoLand
*/
package bootstrap

//系统配置加载
func InitBootstrap(confFile string) (err error) {
	//加载日志配置
	InitLogger()
	//加载 prometheus
	InitMetrics()
	//加载 配置文件
	if err = InitConfig(confFile); err != nil {
		return
	}
	//加载 DB ORM
	if err = InitDb(); err != nil {
		return
	}
	return nil
}
