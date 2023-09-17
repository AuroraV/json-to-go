### JSON-to-GO：将JSON转换成Go结构体定义 【GOLANG实现】
##### 灵感和输出结果来自[json-to-go](https://github.com/mholt/json-to-go）

##### 用法：
```golang
// 解析json元信息
meta := Parse(json)
// 根据元信息渲染输出结构体定义
structDef := GenerateStruct(meta, inline)
```

### 安装：
```shell
go get github.com/AuroraV/json-to-go
```

### 使用样例
[在线编辑](https://go.dev/play/p/-MKUWeDBml7)
```golang
```

### License
This software is released under the MIT License, see LICENSE.
