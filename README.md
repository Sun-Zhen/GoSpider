### TODO
* 修改glog记录的id为协程id

### 依赖包
* go get gopkg.in/yaml.v2 解析配置文件
* go get github.com/golang/glog 日志,去修改做定制实现
  1. -logtostderr=false
  2. -alsologtostderr=false
  3. -stderrthreshold=ERROR
  4. -log_dir=""
  5. -log_backtrace_at=""
  6. -v=0
  7. -vmodule=""
* github.com/PuerkitoBio/goquery 解析html
* "go.gopath": "/Users/alden/CustomeProjects/GoSpider",

### 注
#### URL的格式由三部分组成：
 1. 第一部分是协议(或称为服务方式)。
 2. 第二部分是存有该资源的主机IP地址(有时也包括端口号)。
 3. 第三部分是主机资源的具体地址，如目录和文件名等。