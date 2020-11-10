package config

const (
	ListenPort     = 50000 // Check ip列表的时候的监听端口
	Threshold      = 100   // 当包的大小大于阈值的时候，计数接受
	WaitTime       = 10    // 全部发包完毕后，等待其余数据包的时间
	Blacklists     = "/data/blacklists/blacklists"
	MaxAtktime     = 300 // 最大攻击时间
	AttackInterval = 0
	ShodanApi      = ""
	ShodanPage     = 10 // 默认搜索页数，10页=1000个
	ZoomeyeApi     = ""
	ZoomeyePage    = 50    // 默认搜索页数，50页=1000个
	HttpPort       = 65466 // Http默认监听端口
)
