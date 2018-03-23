package config

type SpiderConfig struct {
	BaseUrl        string // 域名信息
	StartPage      int    // 其实页码
	EndPage        int    // 结束页码
	OutPut         string // 输出路径
	AnalyzeGoSize  int    // 分析视频页面的Goroutine个数
	DownloadGoSize int    // 用于下载视频的Goroutine的个数
}
