### JSON-to-GO：将JSON转换成Go结构体定义 【纯GOLANG实现】
##### 灵感来自[json-to-go](https://github.com/mholt/json-to-go)

##### 用法：
```golang
// 解析json元信息
meta := Parse(json)
// 根据元信息渲染输出结构体定义
// inline:true/false;结构体内联/分离输出
structDef := GenerateStruct(meta, inline)
```

### 安装：
```shell
go install github.com/AuroraV/json-to-go/cmd/g2j@v0.1.2

j2g -s '{"a":123, "b":{"c":1.23, "d":["9", "8"]}}' -f true --inline true 
```

### Features:
- [x] json转golang结构体定义输出
- [x] 输出模式可选:内联或分离
- [ ] CLI交互式生成结构体定义文件
- [ ] 内置简单web页，支持在线编辑
- [ ] 更多JSON生成代码的探索
- [ ] 支持对JSON流进行业务含义定义，类似字典赋能业务
- [ ] ...
> JSON结构是业内常用的传输语言； 若您有更多JSON可视化想法或对此项目有疑问，欢迎积极参与提ISSUE，您将获得积极的响应。

### 使用样例
[在线编辑](https://go.dev/play/p/C67JZortfSI)
```golang
// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"

	j2g "github.com/AuroraV/json-to-go"
)

func main() {
	meta := j2g.Parse(`[{
	"input_index": 100000001,
	"components": [{
		"primary_number": "1",
		"a": 0.83
	}, {
		"a": 1.01,
		"b": 1.28
	}]
}]`)
	structDef := j2g.GenerateStruct(meta, true)
	structDef = j2g.GenerateStruct(meta, false)
	fmt.Println(strings.Join(structDef, "\n----------------------------------------\n"))
	// output:
	/*
		1. output [inline = true] :
		type AutoGenerated []struct {
			InputIndex int64 `json:"input_index"`
			Components []struct {
				PrimaryNumber string `json:"primary_number,omitempty"`
				A float64 `json:"a"`
				B float64 `json:"b,omitempty"`
			} `json:"components"`
		}

		2. output [inline = false] :
		type AutoGenerated []struct {
				InputIndex int64 `json:"input_index"`
				Components []Components `json:"components"`
		}
		----------------------------------------
		type Components struct {
				PrimaryNumber string `json:"primary_number,omitempty"`
				A float64 `json:"a"`
				B float64 `json:"b,omitempty"`
		}
	*/
}
```

### License
This software is released under the MIT License, see LICENSE.
