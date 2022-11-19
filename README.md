# udpgo

#### 介绍
统一开发平台（Go语言版）
实现对公共能力的统一封装
主要功能
1、集成日志功能（参考lang包）
2、集成kafka功能（参考kafka包）
3、集成influxdb line protocol功能(参考influxdb)
4、常用的字符串等处理功能

#### 软件架构



#### 安装教程

go get github.com/chaolihf/udpgo

#### 使用说明

参考测试用例
1、格式化字符串处理FormatStringer，可以实现对字符串中命名的变量进行替换等操作，命名的变量格式为{变量名称}，如字符串中需要输出{、}、\字符，则需要进行转义，如\{,\},\\。
示例：  `你好{祖国}` ，代表变量为祖国，可以使用replace方法对特定位置的变量进行替换

#### 参与贡献


#### 特性

#### 变更
v0.0.5 增加字符串功能: LeftString , 支持utf-8字符串截取，并且不会当索引超过长度不会报异常

v0.0.7 基于gojson(https://github.com/ChengjinWu/gojson.git),增加修改和序列化功能

v0.0.9 增加格式化字符串替换

v0.0.10 修复json数组，字符串等获取错误，增加keys方法