<h1 align="center">
    <img align="center" height="96px" src=".github/images/fiber_logo.svg" alt="Fiber logo" /><br/>
    Fiber backend template<br/>
    for <a href="https://github.com/create-go-app">Create Go App</a>
</h1>

<p align="center"><img src="https://img.shields.io/badge/Go-1.11+-00ADD8?style=for-the-badge&logo=go" alt="go version" />&nbsp;<a href="https://gocover.io/github.com/create-go-app/fiber-go-template/pkg/apiserver" target="_blank"><img src="https://img.shields.io/badge/Go_Cover-88%25-success?style=for-the-badge&logo=none" alt="go cover" /></a>&nbsp;<a href="https://goreportcard.com/report/github.com/create-go-app/fiber-go-template" target="_blank"><img src="https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none" alt="go report" /></a>&nbsp;<img src="https://img.shields.io/badge/license-mit-red?style=for-the-badge&logo=none" alt="lisense" /></p>

## ⚡️ Quick start guide

1. Create a new app with this template by [Create Go App CLI](https://github.com/create-go-app/cli):

```console
cgapp -p ./my-app -b fiber
```

2. Go to the `./my-app` folder
3. Run app by command:

```console
task -s
```

> ☝️ Please note: we're using [`Taskfile`](https://taskfile.dev) as task manager by default.

## ☕️ Description

Fiber is an `Express` inspired web framework build on top of `Fasthttp`, the fastest HTTP engine for Go. Designed to ease things up for fast development with zero memory allocation and performance in mind.

People switching from Node.js to Go often end up in a bad learning curve to start building their webapps, this project is meant to ease things up for fast development, but with zero memory allocation and performance in mind.

📚 [Documentation](https://docs.gofiber.io/)

## ✅ Used packages

- [gofiber/fiber](https://github.com/gofiber/fiber) `v1.12.1`
- [go-yaml/yaml](https://github.com/go-yaml/yaml) `v2.3.0`
- [stretchr/testify](https://github.com/stretchr/testify) `v1.6.1`

## 🗄 Template structure

```console
.
├── .dockerignore
├── .editorconfig
├── .gitignore
├── Dockerfile
├── LICENSE
├── README.md
├── Taskfile.yml
├── go.mod
├── go.sum
├── cmd
│   └── apiserver
│       └── main.go
├── configs
│   └── apiserver.yml
├── static
│   └── index.html
└── pkg
    └── apiserver
        ├── config.go
        ├── config_test.go
        ├── error_checker.go
        ├── error_checker_test.go
        ├── new_server.go
        ├── new_server_test.go
        └── routes.go

6 directories, 17 files
```

## ⚙️ Configuration

```yaml
# ./configs/apiserver.yml

# Server config
server:
  host: 127.0.0.1
  port: 8080

# Database config
database:
  host: 127.0.0.1
  port: 5432
  username: postgres
  password: 1234

# Static files config
static:
  prefix: /public
  path: ./static
```

## ⚠️ License

MIT &copy; [Vic Shóstak](https://github.com/koddr) & [True web artisans](https://1wa.co/).
