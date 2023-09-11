//go:generate go run generate.go

package main

import (
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
)

// copied from src/utils.go
func OpenAndReadAllContent(path string) (string, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 借鉴文档中的用法, 将所有数据一次性读出
	var content_b []byte
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return "", err
		}
		//说明读取结束
		if n == 0 {
			break
		}
		//读取到最终的缓冲区中
		content_b = append(content_b, buf[:n]...)
	}
	return string(content_b), nil
}

// 生成checkOp函数, 用于检查op字符串是否合法, 并返回对应的opcode
func generateCheckOp(file *os.File) {
	// 从文件中读取#define
	content, err := OpenAndReadAllContent("ebpf_helper.h")
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile(`#define\s+BOESFS_(\w+)\s+(\d+)`)
	matches := re.FindAllStringSubmatch(content, -1)
	if matches == nil {
		log.Fatal("re.FindAllStringSubmatch(content, -1)")
	}

	// constMap := make(map[string]string) map会导致switch case中的case顺序不确定
	type caseStruct struct {
		Name  string // case的选项, 即操作类型字符串
		Value string // case内函数的返回值, 即对应的操作码
	}
	var caseList []caseStruct
	for _, match := range matches {
		name := match[1]
		value := match[2]
		// constMap["BOESFS_"+name] = value
		caseList = append(caseList, caseStruct{Name: strings.ToLower(name), Value: value})
	}

	const templateCode = `
// 反正我自己也不用宏定义, 所以这里直接将宏定义展开写死
// 要是改了 直接重新 go generate
func checkOp(s string) int {
	switch s {
{{- range .}}
	case "{{.Name}}":
		return {{.Value}}
{{- end}}
	default:
		return -1
	}
}
`
	// 解析模板
	tmpl := template.Must(template.New("opCode").Parse(templateCode))

	// 将数值插入模板中，并添加到文件
	err = tmpl.Execute(file, caseList)
	if err != nil {
		log.Fatal(err)
	}
	println("generate success")
}

// 生成enum
func generateEnums(file *os.File) {
	// 从文件中读取#define
	content, err := OpenAndReadAllContent("acl_public.h")
	if err != nil {
		log.Fatal(err)
	}

	// 定义模板
	const templateCode = `
// {{.Name}}
const (
	{{- .Value -}}
)
`
	tmpl := template.Must(template.New("opCode").Parse(templateCode))

	// 匹配所有的enum
	re := regexp.MustCompile(`typedef\s+enum\s+(\w+)\s*\{(\s*[^}]+)\}\s*(\w+);`)
	matches := re.FindAllStringSubmatch(content, -1)
	if matches == nil {
		log.Fatal("re.FindAllStringSubmatch(content, -1)")
	}

	type enumStruct struct {
		Name  string
		Value string
	}
	// 依次按照模板生成到文件中
	for _, match := range matches {
		enumName := match[1]
		content := match[2]
		content = strings.Replace(content, ",", " = iota", 1)
		content = strings.ReplaceAll(content, ",", "")
		err = tmpl.Execute(file, enumStruct{Name: "enum " + enumName, Value: content})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	// 打开目标文件
	file, err := os.Create("../src/c_headers.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 写入文件头
	const templateCode = `// Code generated by "go run generate.go"; DO NOT EDIT.

package gofs
`
	tmpl := template.Must(template.New("opCode").Parse(templateCode))
	tmpl.Execute(file, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 生成checkOp函数
	generateCheckOp(file)
	// 生成enum
	generateEnums(file)
}
