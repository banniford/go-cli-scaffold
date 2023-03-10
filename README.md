# go-cli-scaffold
该项目参考 github-cli 使用 Cobra 框架创建一个cli开发脚手架

```txt
├── cmd/
│   ├── root.go
│   ├── command1.go
│   └── command2.go
├── internal/
│   ├── pkg1/
│   │   ├── file1.go
│   │   └── file2.go
│   ├── pkg2/
│   │   ├── file1.go
│   │   └── file2.go
│   └── config/
│       └── config.go
├── pkg/
│   ├── module1/
│   │   ├── file1.go
│   │   └── file2.go
│   ├── module2/
│   │   ├── file1.go
│   │   └── file2.go
│   └── common.go
├── vendor/
├── main.go
├── go.mod
└── go.sum
```

下面是每个目录的作用：

`cmd/` 目录：用于存放所有命令相关的代码，包括根命令和子命令。每个命令应该拥有一个单独的 Go 文件，该文件的名称应该与命令的名称相同。例如，如果有一个名为 `"command1"` 的命令，则该命令的代码应该存放在 `cmd/command1.go` 文件中。

`internal/` 目录：用于存放所有私有代码，这些代码只能被本项目中的其他代码引用。该目录下的子目录应该根据它们所提供的功能进行命名，并包含相关的代码文件。

`pkg/` 目录：用于存放所有公共的模块代码。该目录下的每个子目录应该包含一组相关的代码文件，并使用模块名称命名。例如，如果有一个名为 `"module1"` 的模块，则该模块的代码应该存放在 `pkg/module1/` 目录中。

`vendor/` 目录：用于存放项目所依赖的第三方库的代码。该目录可以由 Go 工具自动生成，并用于确保项目的依赖关系不会受到外部变化的影响。

`main.go` 文件：用于启动项目并注册所有命令。

`go.mod` 和 `go.sum` 文件：用于管理项目的依赖关系。`go.mod` 文件定义了项目的依赖项，`go.sum` 文件则记录了依赖项的确切版本信息，以确保所有团队成员在构建项目时使用的都是相同的依赖版本。

这是一个基本的 Cobra 项目结构，您可以根据自己的需求对其进行调整。