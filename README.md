<h3 align="center">Project Title</h3>


<div align="center">


[![Status](https://img.shields.io/badge/status-active-success.svg)](https://github.com/eryajf/care-screenshot)
[![GitHub Issues](https://img.shields.io/github/issues/kylelobo/The-Documentation-Compendium.svg)](https://github.com/eryajf/care-screenshot/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/kylelobo/The-Documentation-Compendium.svg)](https://github.com/eryajf/care-screenshot/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)


</div>


---

<p align="center"> 给我一个URL，我能截图发给企业微信机器人🤖
    <br> 
</p>



## 使用


查看命令帮助:


```bash
$ ./care-screenshot exec -h
命令行工具，可使用此工具订阅一些你关心的网页服务状态，然后添加到定时任务中。

Usage:
  care-screenshot exec [flags]

Flags:
  -b, --bot string       机器人地址 (default "d63e3f22-3a88-43fb-a2ad-ad78ba5b43b5")
  -e, --element string   给我你关心的页面元素 (default "#s_lg_img")
  -g, --gao string       页面高度 (default "800")
  -h, --help             help for exec
  -k, --kuan string      页面宽度 (default "1200")
  -u, --url string       给我一个你想要截图的URL (default "https://baidu.com")
```

## 上手

如果我们想要拿到百度首页的logo，那么访问百度之后，点击检查，找到想要的元素：

![](http://t.eryajf.net/imgs/2021/09/4a491889fa6629a1.jpg)

> 1. 进入页面，点击检查。
> 2. 使用选择工具找到自己想要截图的区域。
> 3. 右键点击对应元素的选项。
> 4. 找到Copy。
> 5. 选择 `Copy selector`。

此时在剪切板中会复制到对应页面的元素标签。

然后就可以运行如下命令拿到我们想要的图片并自动发送到微信群：

```bash
$ ./care-screenshot exec  -u 'https://www.baidu.com' -e '#s_lg_img' -k 800 -g 800
```

然后就能在群里看到程序截的图了。

![](http://t.eryajf.net/imgs/2021/09/a02d2f8664a60a73.png)

## 实践

再举个栗子，比如我们经常做一些变更，希望通过日志能观测到变化，但是总去`kibana`看不太方便，就可以订阅一下关心的图，比如我想要这个`dashboard`中状态码`5xx`的框框，注意：尽量使用资源的短链接。

```bash
$ ./care-screenshot exec  -u 'http://10.6.6.5:5601/goto/d700abb7461c1e5b91cf5c6579a15b69' -e '#dashboardViewport > div > div > div:nth-child(2) > div' -k 2000 -g 800
```

接着就能在群里看到对应区块的截图了：

![](http://t.eryajf.net/imgs/2021/09/86560a9c04357548.jpg)

## 容器

程序在主机上运行的时候，如果检查到运行环境不满足需求，会自动下载相关依赖，如果你想来去无痕地运行程序，可以使用如下程序运行：

```bash
docker run -it --rm registry.cn-hangzhou.aliyuncs.com/ali_eryajf/chrome-go-rod:v0.0.1 care-screenshot exec -u 'https://www.baidu.com' -e '#s_lg_img' -k 800 -g 800
```

通过这种方式，也方便我们将程序放到系统定时任务中定期拿到我们关心的数据了。

## 感谢

感谢如下优秀的项目：

- [rod](https://github.com/go-rod/rod)

## 另外

- 如果觉得项目不错，麻烦动动小手点个⭐️start⭐️!
- 如果你还有其他项目或者需求，欢迎在issue中交流！
- 程序还有很多bug，部分页面可能截图失败，可重试一波，如果还失败，那就与我一起来维护这个项目吧！