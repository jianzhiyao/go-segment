# GO 中文分词服务(HTTP)
## 基于Sego 
### 项目地址
[github.com/huichen/sego](https://github.com/huichen/sego/)
### 基本说明
[词典](https://github.com/huichen/sego/blob/master/dictionary.go)用双数组trie（Double-Array Trie）实现，
[分词器](https://github.com/huichen/sego/blob/master/segmenter.go)算法为基于词频的最短路径加动态规划。

支持普通和搜索引擎两种分词模式，支持用户词典、词性标注，可运行[JSON RPC](https://github.com/huichen/sego/blob/master/server/server.go)服务</a>。

分词速度[单线程](https://github.com/huichen/sego/blob/master/tools/benchmark.go)单线程9MB/s，[goroutines并发](https://github.com/huichen/sego/blob/master/tools/goroutines.go)42MB/s（8核Macbook Pro）。


## 运行项目
### GO环境
具体百度

### 安装 Sego 命令
    ```shell
    go get -u github.com/huichen/sego
    ```    

### 编译项目
编译项目可以编译出当前系统的可执行文件，这里以 Linux 系统为例
```shell
go build
```    

### 运行项目
```shell
/go-segment-path/go-segment {port}
```

样例
```shell
/go-segment-path/go-segment 8787
```

### 守护进程（deamon）
可以用 supervisor 等进程管理工具挂起 


## 提供接口(HTTP)
### 接口地址
-X POST http://127.0.0.1:{port}/segment
### 请求参数
| 参数 | 可选| 类型 | 说明 |
|:--- |:---|:--- |:--- |
| content |必选| string| 需要分词的字符串，最大允许 10000 长度的字符串 |
| deep_search | 可选 | int| 是否进行深度分词，注意：有此参数就会进行深度分词，无论什么值|


### 响应结果
| 参数 |  类型 | 说明 |
|:--- |:--- |:--- |
| Status | int| 分词结果，1：成功，0：失败 |
| Msg |  string| 返回的相关信息|
| Response |  string| 分词结果|

#### 分词结果格式说明
1.每组分词以空格隔开
2.每组分词以：词 + 斜杠 + 词性组成


#### 成功实例
```json
{
	"Status": 1,
	"Msg": "",
	"Response": "测试/vn  /x 九/m 华山/ns 九华山/ns 龙泉/nz 寺庙/n "
}
```

#### 失败实例
```json
{
	"Status": 0,
	"Msg": "超过最大处理内容长度",
	"Response": ""
}
```

#### 请求实例 CURL 命令
```shell
curl -X POST "http://127.0.0.1:8787/segment" --data "deep_search=w&content=华 山龙泉寺庙测试"
```

