# GO 分词服务
## 基于项目
[sego](https://github.com/huichen/sego/)
## 基本说明
[词典](https://github.com/huichen/sego/blob/master/dictionary.go)用双数组trie（Double-Array Trie）实现，
[分词器](https://github.com/huichen/sego/blob/master/segmenter.go)算法为基于词频的最短路径加动态规划。

支持普通和搜索引擎两种分词模式，支持用户词典、词性标注，可运行[JSON RPC](https://github.com/huichen/sego/blob/master/server/server.go)服务</a>。

分词速度[单线程](https://github.com/huichen/sego/blob/master/tools/benchmark.go)单线程9MB/s，[goroutines并发](https://github.com/huichen/sego/blob/master/tools/goroutines.go)42MB/s（8核Macbook Pro）。
