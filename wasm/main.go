package main

import (
	"errors"
	"fmt"
	"go-tool/wasm/component"
	"syscall/js"
)

var (
	input     = js.Global().Get("userInput")
	output    = js.Global().Get("userOutput")
	genButton = js.Global().Get("genButton")
	error_    = js.Global().Get("error")
	method    = js.Global().Get("method")
)

func main() {
	fmt.Println("golang wasm running")

	js.Global().Set("ClearContent", js.FuncOf(ClearContent))
	Handler()

	<-make(chan int, 0)
}

func Handler() {
	//	绑定点击事件
	genButton.Call("addEventListener", "click", js.FuncOf(makeData))
}

func ClearContent(this js.Value, args []js.Value) any {
	output.Set("innerHTML", "")
	input.Set("innerHTML", "")
	outputError("")
	return nil
}

func outputError(content string) {
	error_.Set("innerHTML", content)
}

func makeData(this js.Value, args []js.Value) any {
	var res string
	var err error

	defer func() {
		output.Set("innerHTML", "")
		input.Set("innerHTML", "")
		outputError("")
		if err != nil {
			outputError(err.Error())
		} else if res != "" {
			output.Set("innerHTML", res)
		}

	}()

	//输入内容
	inputContent := input.Get("value").String()
	if inputContent == "" {
		err = errors.New("请输入内容")
		return nil
	}

	//方法
	methodContent := method.Call("getAttribute", "value").String()
	if methodContent == "" {
		return nil
	}

	switch methodContent {
	case "json2go":
		res = component.Json2Go(inputContent)
		break
	case "sql2gorm":
		var sList []string
		sList, err = component.Sql2Gorm(inputContent)
		if err != nil {
			return nil
		}
		for _, v := range sList {
			res += v
		}

	default:
		fmt.Println("未找到类型")
	}

	return nil
}
