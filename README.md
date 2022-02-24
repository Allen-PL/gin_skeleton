日志zerolog
- 采用链式调用
- 嵌套操作
- 全局log
    - log.Debug()
    - log.Info()
- 设定日志级别(7个级别)
    - panic
    - fatal
    - error
    - warn
    - info
    - debug
    - trace
  
配置文件管理神器Viper
- 设置默认值
- 从JSON、TOML、YAML、HCL...格式的配置文件读取配置信息
- 实时监控和重新读取配置文件
- 从环境变量中读取
- 从远程配置系统读取病监控配置变化
- 从命令行参数读取配置
- 从buffer读取配置
- 显式配置值

viper.SetDefault(key, value)设置默认值
viper.SetConfigFile(path)指定配置文件路径
viper.AddConfigPath(path)查找配置文件所在的路径
err := viper.ReadInConfig()查找并读取配置文件
viper.WriteConfig()将当前配置写入“viper.AddConfigPath()”和“viper.SetConfigName”设置的预定义路径
Get(key string) : interface{}
GetBool(key string) : bool
GetFloat64(key string) : float64
GetInt(key string) : int
GetIntSlice(key string) : []int
GetString(key string) : string
GetStringMap(key string) : map[string]interface{}
GetStringMapString(key string) : map[string]string
GetStringSlice(key string) : []string
GetTime(key string) : time.Time
GetDuration(key string) : time.Duration
IsSet(key string) : bool  判断读取的值是否存在
AllSettings() : map[string]interface{}










