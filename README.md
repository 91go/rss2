# rss2

[![golangci-lint](https://github.com/91go/rss2/actions/workflows/golangci-lint.yml/badge.svg?branch=main)](https://github.com/91go/rss2/actions/workflows/golangci-lint.yml)
[![test and deploy](https://github.com/91go/rss2/actions/workflows/deploy.yml/badge.svg?branch=main)](https://github.com/91go/rss2/actions/workflows/deploy.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/91go/rss2)](https://goreportcard.com/report/github.com/91go/rss2)



## todo

1. 用rss管理音乐文件；随便找个带歌单的音频播放器js，把本地音乐嵌进去，做成一个页面；
    1. [maomao1996/Vue-mmPlayer: 🎵 基于 Vue 的在线音乐播放器（PC） Online music player](https://github.com/maomao1996/Vue-mmPlayer)
2. ~~添加把webdav资源转rss的功能；~~
3. 只保留goframe，踢掉gin；
4. ~~把cicd从drone换成github-ci；~~
5. 优化docker构建；


## 性能优化

1. ~~并发获取详情页数据~~
2. 全面实现rsshub的底层代码
3. 实现应用可观测
4. ~~实时监控网页变化，借鉴LogicJake/WebMonitor~~



## 实现一个管理xxx的页面


1. 类似`rss.app`那样，在页面上增删改查某个信息源，把所有数据聚合到一个feed输出；
2. 删除是逻辑删除，可以revoke；





```markdown
有哪些删掉的视频博主？


1. 辩论(新国辩之类的辩论视频)
2. 历史&键政(古代史、近代史、各国家历史、时政评论)
    1. 南海望龙
    2. 范勇鹏
    3. 万归藏-西瓜视频
    4. 慕有枝613
    5. 草说木言✅
    6. 安森垚✅
    7. IC实验室
    8. 小王Albert
    9. 经济研究室-祈祷
    10. 一条闲木鱼
    11. 马督工
3. 投资&财经(比如基金投资视频以及财经新闻等泛财经视频)
    1. 基金研究员阿鸡fundchick
4. 科普
5. 评测
    1. 穿搭
        1. ITAKE-STUDIO
        2. 吕政懋MaoMao
        3. 登山者black(户外)✅
    2. PC
        1. 搞机所
        2. Eixa工作室(专业itx测评)✅
        3. 翼王
        4. 微机分WekiHome
        5. 有梦想的阿肯老师
        6. FUN科技
        7. 潮玩客
        8. 科技美学(基本上全是广告)
    3. 数码&家电
        1. 先看评测&选品君✅
        2. 钟文泽(只有开箱，没有评测，节目效果好但是不推荐)
    4. 综合
        1. 老爸评测
    5. 鞋子
        1. 鞋吧Sneakersbar
        2. 球鞋师说
        3. 亚平宁的蓝色(专业跑鞋测评)✅
    6. 装修
        1. Mr迷瞪✅
    7. 其他
        1. 内幕纠察局
        2. 韭菜实验室
        3. 凰家实验室
        4. 滤镜粉碎机
    8. 汽车
        1. 懂车帝原创
6. 代码相关
    1. FunInCode
    2. 跟李沐学AI
    3. `MikeTang84`Rust唠嗑室

```


