# Txray
Txray是一款 xray 终端版客户端，使用go编写。

项目地址：https://github.com/hsernos/Txray

Project X core： https://github.com/XTLS/Xray-core

## 注意
此文档只针对于最新的commit, 可能不适用于已发布的最新版本.

<!-- toc -->

## 目录

- [特色](#特色)
- [编译/交叉编译 说明](#编译交叉编译-说明)
- [下载/运行 说明](#下载运行-说明)
- [命令列表及说明](#命令列表及说明)
  * [命令总览](#命令总览)
  * [查看基本设置帮助文档](#查看基本设置帮助文档)
    + [查看基本设置](#查看基本设置)
    + [修改基本设置](#修改基本设置)
  * [查看测试设置帮助文档](#查看测试设置帮助文档)
      + [查看测试设置](#查看测试设置)
      + [修改测试设置](#修改测试设置)    
  * [查看订阅帮助文档](#查看订阅帮助文档)
    + [添加订阅](#添加订阅)
    + [查看订阅](#查看订阅)
    + [修改订阅](#修改订阅)
    + [删除订阅](#删除订阅)
    + [从订阅更新节点](#从订阅更新节点)
  * [查看节点帮助文档](#查看节点帮助文档)
    + [添加节点](#添加节点)
    + [查看节点](#查看节点)
    + [删除节点](#删除节点)
    + [tcping测试](#tcping测试)
    + [节点查找](#节点查找)
    + [导出节点](#导出节点)
  * [查看DNS帮助文档](#查看DNS帮助文档)
    + [查看DNS设置](#查看DNS设置)
    + [修改DNS设置](#修改DNS设置)
  * [查看路由帮助文档](#查看路由帮助文档)
    + [添加路由](#添加路由)
    + [domain路由规则](#domain路由规则)
    + [ip路由规则](#ip路由规则)
  * [启动或重启xray-core服务](#启动或重启xray-core服务)
  * [停止xray-core服务](#停止xray-core服务)
- [已知问题](#已知问题)
- [交流反馈](#交流反馈)

<!-- tocstop -->

# 特色

1. 多平台支持, 支持 Windows, macOS, linux.
2. Tab键命令补齐
3. 支持VMess、Shadowsocks、Trojan协议

#  编译/交叉编译 说明

1. 在终端下进入项目目录

2. 设置`GOPROXY`,提高编译所需依赖的下载速度
   Linux/Mac下，运行 `GOPROXY=https://goproxy.cn,direct`
   Windows下,运行 `set GOPROXY=https://goproxy.cn,direct`

3. 编译常用平台
   运行 `go build Txray.go`, 可编译当前平台的版本
   运行 `python3 build.py`, 可编译常用平台的版本

4. 编译其他平台
   运行 `go tool dist list` 查看所有支持的 GOOS/GOARCH

   Linux/Darwin 例子: 编译 Windows 下的 64 位程序

   `GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build Txray.go`
   
   Windows 例子: 编译 Linux 下的 32 位程序
   
   `set GOOS=linux`
   `set GOARCH=386`
   `set CGO_ENABLED=0`
   `go build Txray.go`


# 下载/运行 说明

需要下载[xray-core](https://github.com/XTLS/Xray-core/releases)

1. Txray检测xray-core所在的优先级: 环境变量 `CORE_HOME` > Txray所在目录
2. 配置文件目录优先级:  环境变量 `TXRAY_HOME` > Txray所在目录


# 命令列表及说明

> 在终端中运行Txray进入shell交互

## 命令总览
```
Commands:
    base                     基础设置             使用 'base help' 查看详细用法
    dns                      DNS 设置             使用 'dns help' 查看详细用法
    test                     测试设置             使用 'test help' 查看详细用法
    node                     节点管理             使用 'node help' 查看详细用法
    sub                      订阅管理             使用 'sub help' 查看详细用法
    routing                  路由管理             使用 'routing help' 查看详细用法
    help, -h                 查看帮助信息
    version, -v              查看版本
    clear                    清屏
    exit                     退出程序
    stop                     停止服务
    run                      启动或重启服务

Usage: run [索引式 | -t [索引式]]

    run [索引式]      默认为上一次运行节点，如果选中多个节点，则选择访问YouTube延迟最小的
    run -t [索引式]   按tcp延迟选择节点，默认'1'，比如 'run -t 1-10' 为选择tcp延迟最小的10个节点


说明：
一、索引式：更简单的批量选择
1.选择前6个：'1,2,3,4,5,6' 或 '1-3,4-6' 或 '1-6' 或 '-6'
2.选择第6个及后面的所有：'6-'
3.选择第6个：'6'
4.选择所有：'all' 或 '-'
注意：超出部分会被忽略，'all' 只能单独使用

二、[] 和 {}：帮助说明中的中括号和大括号
1. []: 表示该选项可忽略
2. {}: 表示该选项为必须，不可忽略
```

## 查看基本设置帮助文档

```
>>> base help
base {commands}

Commands:
                                  查看基本设置
    help                          查看帮助
    port {port}                   设置socks端口
    http {port}                   设置http端口, 0为关闭http监听
    udp {y|n}                     是否启用udp
    sniffing {y|n}                是否启用流量监听
    lanconn {y|n}                 是否启用局域网连接
    mux {y|n}                     是否启用多路复用
    bypass {y|n}                  是否绕过局域网及大陆
    routing {1|2|3}               设置路由策略为{AsIs|IPIfNonMatch|IPOnDemand}
```

### 查看基本设置

```
# 
>>> base
+------------+----------+---------+--------------+----------+----------------+------------------+--------------+
|  SOCKS端口  | HTTP端口  | UDP转发 |   启用流量监听  |  多路复用  |  允许局域网连接  | 绕过局域网和大陆    |   路由策略     |
+------------+----------+---------+--------------+----------+----------------+------------------+--------------+
|    2333    |   2334   |  true   |    false     |  false   |     false      |       true       | IPIfNonMatch |
+------------+----------+---------+--------------+----------+----------------+------------------+--------------+
```

### 修改基本设置

```
# 修改socks监听端口为3333
>>> base socks 3333

# 修改http监听端口为3334
>>> base http 3334

# 修改不绕过局域网和大陆
>>> base bypass n

# 修改路由策略为IPIfNonMatch, {1|2|3}=>{AsIs|IPIfNonMatch|IPOnDemand}
>>> base routing 2
```


## 查看测试设置帮助文档

```
>>> test help
test [commands]

Commands:
                                  查看测试设置
    help                          查看帮助
    url {url}                     设置测试网站
    timeout {time}                设置测试超时时间 (秒)
```

### 查看测试设置

```
>>> test
+-------------------------+----------------+
|         测试URL         | 超时时间（秒）  |
+-------------------------+----------------+
| https://www.youtube.com |       5        |
+-------------------------+----------------+
```

### 修改测试设置

```
# 修改测试URL为google
>>> test url https://google.com

# 修改超时时间为10秒
>>> test timeout 10
```

## 查看订阅帮助文档

```
>>> sub help
sub {commands} [flags] ...

Commands:
    [索引式]                      查看订阅信息
    help                          查看帮助
    rm {索引式}                   删除订阅
    add {订阅url} [flags]         添加订阅
    mv {索引式} {flags}           修改订阅
    update-node [索引式] [flags]  从订阅更新节点, 索引式会忽略是否启用

add Flags
    -r, --remarks {别名}          定义别名

rm Flags
    -u, --url {订阅url}           修改订阅链接
    -r, --remarks {别名}          定义别名
    --using {y|n}                 是否启用此订阅

update-node Flags
    -s, --socks5 [port]           通过本地的socks5代理更新, 默认为设置中的socks5端口
    -h, --http [port]             通过本地的http代理更新, 默认为设置中的http端口
    -a, --addr {address}          对上面两个参数的补充, 修改代理地址
```

### 添加订阅

```
# 添加订阅链接为https://sublink.com
>>> sub add https://sublink.com

# 添加订阅链接为https://sublink.com，并命名为test
>>> sub add https://sublink.com -r test
```

### 查看订阅

```
# 查看全部订阅
>>> sub
+------+-------+---------------------+----------+
| 索引  | 别名   |       URL          |  是否启用  |
+------+-------+---------------------+----------+
|  1   | test1 | https://sublink.com |   true   |
|  2   | test2 | https://sublink.com |   true   |
|  3   | test3 | https://sublink.com |   true   |
|  4   | test4 | https://sublink.com |   true   |
|  5   | test5 | https://sublink.com |   true   |
|  6   | test6 | https://sublink.com |   true   |
+------+-------+---------------------+----------+

# 查看索引为2,3,4的订阅
>>> sub 2-4
+------+-------+---------------------+----------+
| 索引  | 别名  |         URL         |  是否启用  |
+------+-------+---------------------+----------+
|  2   | test2 | https://sublink.com |   true   |
|  3   | test3 | https://sublink.com |   true   |
|  4   | test4 | https://sublink.com |   true   |
+------+-------+---------------------+----------+
```

### 修改订阅

```
# 修改索引为1的订阅链接为https://test.com，别名为test8
>>> sub mv 1 -u https://test.com -r test8
>>> sub 1
+------+-------+------------------+----------+
| 索引  | 别名  |        URL        |  是否启用 |
+------+-------+------------------+----------+
|  1   | test8 | https://test.com |   true   |
+------+-------+------------------+----------+

# 禁用索引为3和5的订阅链接
>>> sub mv 3,5 --using n
>>> sub 
+------+-------+---------------------+----------+
| 索引 |  别名  |         URL         |  是否启用  |
+------+-------+---------------------+----------+
|  1   | test8 | https://sublink.com |   true   |
|  2   | test2 | https://sublink.com |   true   |
|  3   | test3 | https://sublink.com |  false   |
|  4   | test4 | https://sublink.com |   true   |
|  5   | test5 | https://sublink.com |  false   |
|  6   | test6 | https://sublink.com |   true   |
+------+-------+---------------------+----------+
```

### 删除订阅

```
# 删除索引为3和5的订阅
>>> sub rm 3,5

# 删除所有订阅
>>> sub rm all
```

### 从订阅更新节点

```
# 从启用的订阅且不使用代理更新节点
>>> sub update-node

# 从索引范围更新节点，无论是否启用
>>> sub update-node 1,3,6

# 使用端口为2333的本地socks5代理更新节点
>>> sub update-node -s 2333

# 使用设置中的socks端口通过本地socks5代理更新节点
>>> sub update-node -s

# 使用端口为2334的本地http代理更新节点
>>> sub update-node -h 2334

# 使用端口为2333，地址为1.2.3.4的socks代理更新节点
>>> sub update-node -s 2333 -a 1.2.3.4
```



## 查看节点帮助文档

```
>>> node help
node {commands} [flags] ...

Commands:
    [索引式|t]                    查看节点信息, 默认'all', 't'表示按延迟降序查看
    help                          查看帮助
    info {索引}                   查看单个节点详细信息
    rm {索引式}                   删除节点
    tcping {索引式}               测试节点tcp延迟
    find {关键词}                 查找节点（按别名）
    add [flags]                   添加节点
    export [索引式] [flags]       导出节点链接, 默认'all'

add Flags
    -l, --link {link}             从链接导入一条节点
    -f, --file {path}             从节点链接文件或订阅文件导入节点
    -c, --clipboard               从剪贴板读取的节点链接或订阅文本导入节点

export Flags
    -c, --clipboard               导出节点链接到剪贴板
```

### 添加节点

```
# 添加一个vmess节点
>>> node add  -l vmess://xxxxxxXXXXxxxxxXX

# 添加一个trojan节点
>>> node add  -l trojan://xxxxxxXXXXxxxxxXX

# 由链接文件批量添加节点
>>> node add -f /home/links.txt

# 解析订阅文件添加节点，可以将订阅文件下载下来然后从本地导入
>>> node add -f /home/subtext.txt

# 从剪贴板读取的节点链接或订阅文本导入节点, 功效和上面从文件导入一样
>>> node add -c

# 手动添加一个节点
>>> node add
```

### 查看节点

```
# 查看前20个节点
>>> node 1-20

# 查看某个节点的全部信息
>>> node info 1

# 查看按tcp延迟排序的节点
>>> node t

```

### 删除节点

```
# 删除前20个节点
>>> node rm 1-20
```

### tcping测试

```
# tcping前20个节点 （'-20' 等价于 '1-20'）
>>> node tcping -20
```

### 节点查找

```
# 查找关键词为'vip'的节点
>>> node find vip

# 查找关键词为'香港'的节点
>>> node find "香港"
```

### 导出节点

```
# 导出前20个节点到终端
>>> node export -20

# 导出前20个节点到剪贴板
>>> node export -20 -c
```

## 查看DNS帮助文档

```
>>> dns help
dns {commands}

Commands:
                                  查看DNS设置
    help                          查看帮助
    port {port}                   设置dns端口, 0为关闭
    inland {dns}                  设置一条境内DNS
    outland {dns}                 设置一条境外DNS
    backup {dns}                  设置备用DNS，多条以 ',' 分隔
```

### 查看DNS设置

```
# 查看DNS设置
>>> dns
+---------+---------+-----------+---------+
| DNS端口  | 境外DNS  |  境内DNS   | 备用DNS |
+---------+---------+-----------+---------+
|  23333  | 1.1.1.1 | 223.6.6.6 |         |
+---------+---------+-----------+---------+
```

### 修改DNS设置

```
# 修改dns监听端口为23334
>>> dns port 23334

# 关闭dns监听端口
>>> dns port 0

# 修改境外DNS为8.8.8.8
>>> dns outland 8.8.8.8

# 修改境内DNS为180.76.76.76
>>> dns inland 180.76.76.76

# 修改备用DNS为180.76.76.76
>>> dns backup 180.76.76.76

# 修改备用DNS为180.76.76.76和localhost
>>> dns backup 180.76.76.76,localhost
```


## 查看路由帮助文档

```
>>> route help
routing {commands} [flags] ...

Commands:
    block [索引式] | [flags]      查看或管理禁止路由规则
    direct [索引式] | [flags]     查看或管理直连路由规则
    proxy [索引式] | [flags]      查看或管理代理路由规则

block, direct, proxy Flags
    -a, --add {规则}              添加路由规则
    -r, --rm {索引式}             删除路由规则
    -f, --file {path}             从文件导入规则
    -c, --clipboard               从剪贴板导入规则
```

### 添加路由

```
# 添加www.baidu.com到黑名单
>>> routing block -a www.baidu.com

# 添加www.google.com到代理名单
>>> routing proxy -a www.google.com

# 从文件批量导入到黑名单
>>> routing block -f /home/xxx/block.txt

# 从剪贴板导入到黑名单
>>> routing block -c

```

### domain路由规则

- 纯字符串: 当此字符串匹配目标域名中任意部分，该规则生效。比如`sina.com`可以匹配`sina.com`、sina.com.cn和www.sina.com，但不匹配`sina.cn`。
- 正则表达式: 由`regexp:`开始，余下部分是一个正则表达式。当此正则表达式匹配目标域名时，该规则生效。例如`regexp:\\.goo.*\\.com$`匹配`www.google.com`、`fonts.googleapis.com`，但不匹配`google.com`。
- 子域名 (推荐): 由`domain:`开始，余下部分是一个域名。当此域名是目标域名或其子域名时，该规则生效。例如`domain:xray.com`匹配`www.xray.com`、`xray.com`，但不匹配`xxray.com`。
- 完整匹配: 由`full:`开始，余下部分是一个域名。当此域名完整匹配目标域名时，该规则生效。例如`full:xray.com`匹配`xray.com`但不匹配`www.xray.com`。
- 预定义域名列表：由`"geosite:"`开头，余下部分是一个名称，如`geosite:google`或者`geosite:cn`。名称及域名列表参考[预定义域名列表](https://www.v2ray.com/chapter_02/03_routing.html#dlc)。
- 从文件中加载域名: 形如`ext:file:tag`，必须以`ext:`（小写）开头，后面跟文件名和标签，文件存放在[资源目录](https://www.v2ray.com/chapter_02/env.html#asset-location)中，文件格式与`geosite.dat`相同，标签必须在文件中存在。

### ip路由规则

- IP: 形如`127.0.0.1`。
- [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing): 形如`10.0.0.0/8`.
- GeoIP: 形如`geoip:cn`，必须以`geoip:`（小写）开头，后面跟双字符国家代码，支持几乎所有可以上网的国家。
- 特殊值：`geoip:private` (xray 3.5+)，包含所有私有地址，如`127.0.0.1`。
- 从文件中加载 IP: 形如`ext:file:tag`，必须以`ext:`（小写）开头，后面跟文件名和标签，文件存放在[资源目录](https://www.v2ray.com/chapter_02/env.html#asset-location)中，文件格式与`geoip.dat`相同标签必须在文件中存在。

### 

## 启动或重启xray-core服务

```
# 启动或重启索引为3的节点
>>> run 3

# 自动选择所有节点中访问YouTube延迟最小的那个节点
>>> run all

# 自动选择1-10中访问YouTube延迟最小的那个节点
>>> run 1-10

# 自动选择tcp延迟最小的10个中访问YouTube延迟最小的那个节点
>>> run -t -10
```



## 停止xray-core服务

```
>>>stop
```



# 已知问题

- 有时直接从订阅更新节点失败，可以用浏览器下载订阅文本然后使用 'node add -f {绝对路径}' 导入，或者使用代理导入（sub update-node -s [端口]）
- ss://链接只支持形如 ss://base64编码#别名，trojan://链接只支持形如 trojan://密码@地址:端口#别名

# 交流反馈

提交Issue: [Issues](https://github.com/hsernos/Txray/issues)
