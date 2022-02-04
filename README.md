# rss2

[![golangci-lint](https://github.com/91go/rss2/actions/workflows/golangci-lint.yml/badge.svg?branch=main)](https://github.com/91go/rss2/actions/workflows/golangci-lint.yml)
[![test and deploy](https://github.com/91go/rss2/actions/workflows/deploy.yml/badge.svg?branch=main)](https://github.com/91go/rss2/actions/workflows/deploy.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/91go/rss2)](https://goreportcard.com/report/github.com/91go/rss2)



## todo

1. ~~添加把webdav资源转rss的功能；~~
3. 只保留goframe，踢掉gin；
4. ~~把cicd从drone换成github-ci；~~
5. ~~优化docker构建；~~

开一个新branch, 把rss2按照rsshub的底层彻底改掉, 框架换成eagle; 可以参考rsshub和huginn

[花 2 小时撸一个 RSS 生成器](https://mp.weixin.qq.com/s/mRjoKgkq1PoqlVgOw8oRYw)



添加单元测试用例;




### 实现一个管理xxx的页面


1. 类似`rss.app`那样，在页面上增删改查某个信息源，把所有数据聚合到一个feed输出；
2. 删除是逻辑删除，可以revoke；



### 已取消

1. ~~用rss管理音乐文件；随便找个带歌单的音频播放器js，把本地音乐嵌进去，做成一个页面；~~ [maomao1996/Vue-mmPlayer: 🎵 基于 Vue 的在线音乐播放器（PC） Online music player](https://github.com/maomao1996/Vue-mmPlayer)



## 性能优化

1. ~~并发获取详情页数据~~
2. 全面实现rsshub的底层代码
3. 实现应用可观测
4. ~~实时监控网页变化，借鉴LogicJake/WebMonitor~~





