## go程序版本控制

## 下载

    go get github.com/blurty/go-version

### 使用

    package main

    import (
        "flag"

        "github.com/blurty/go-version"
    )

    func main() {
        showVersion := flag.Bool("v", false, "version info")
        flag.Parse()

        if *showVersion == true {
            version.PrintVersion()
        }
        // your code here
    }

### 编译

1. 将example目录下的build.sh拷贝到你所在的项目顶级目录下
2. 修改脚本中的版本号
3. 运行编译脚本`./build.sh`

### 例子

    ➜  example git:(master) ✗ ls
    build.sh test.go
    ➜  example git:(master) ✗ ./build.sh
    fatal: Needed a single revision
    build finish !!
    Version: 0.0.1
    Git commit: unsupported
    Go version: go version go1.8.3 darwin/amd64
    ➜  example git:(master) ✗ ls
    build.sh example  test.go
    ➜  example git:(master) ✗ ./example -v
    Version:  0.0.1
    Git commit: unsupported
    Go version: go version go1.8.3 darwin/amd64
    Build time: 2017-11-27 23:28:16