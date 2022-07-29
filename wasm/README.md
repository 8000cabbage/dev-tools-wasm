#编译

```bash
 $Env:GOOS="js";$Env:GOARCH="wasm";go build -o ../static/go-tool.wasm;sleep(2);exit
```