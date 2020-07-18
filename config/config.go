package config

// DataDirPath 数据存放目录 这个要做成可配置的
var DataDirPath string = "./data"

// DataVersionFile 数据版本号
var DataVersionFile string = "/version"

// SecDescFile 知识区块的描述文件
var SecDescFile string = "/desc.json"

// CardsDirPath 卡片子目录
var CardsDirPath string = "/cards"

// RandomCardsNumByOneQuery 一次随机获取卡片返回的数量
var RandomCardsNumByOneQuery int = 1
