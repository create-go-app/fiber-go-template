<h1 align="center">Fiber backend template<br/>for <a href="https://github.com/create-go-app">Create Go App CLI</a></h1>

<img align="right" width="256px" src="https://golang.org/lib/godoc/images/go-logo-blue.svg" alt="Golang logo" />

Fiber is an `Express.js` styled HTTP framework implementation running on `FastHTTP`, the fastest HTTP engine for Go. The package make use of similar framework convention as they are in express.

📚 [Documentation](https://gofiber.github.io/fiber)

## Requirements

- Create Go App CLI `v0.5+` ([create-go-app/cli](https://github.com/create-go-app/cli))
- Go `v1.11+` with Go Modules ([golang/download](https://golang.org/dl/))

### Optional

- Docker `v19.x` ([docker/onboarding](https://hub.docker.com/?overlay=onboarding))

## Used packages

- [go-yaml/yaml](https://github.com/go-yaml/yaml) `v3.0.0`
- [zap](https://go.uber.org/zap) `v1.13.0`

## Template structure

```console
foo@bar:~/net_http-go-template$ tree .
.
├── .dockerignore
├── .editorconfig
├── .prettierignore
├── .gitignore
├── Dockerfile
├── Makefile
├── LICENSE
├── README.md
├── go.mod
├── go.sum
├── cmd
│   └── apiserver
│       └── main.go
├── configs
│   └── apiserver.yml
└── internal
    └── apiserver
        ├── apiserver.go
        ├── checker.go
        ├── config.go
        ├── logger.go
        └── routes.go

5 directories, 16 files
```

### Configs

All server/database/logging/static configs included in one YAML file `./configs/apiserver.yml`:

```yaml
# Server config
server:
    host: 127.0.0.1
    port: 8080
    timeout:
        read: 15
        write: 10
        idle: 5

# Database config
database:
    host: 127.0.0.1
    port: 5432
    username: postgres
    password: 1234

# Logging config
logging:
    level: debug

# Static files config
static:
    path: /static

```

### TODO (ASAP list)

- Add more API endpoints
- Add tests
- Add jmoiron/sqlx (Postgres)
- ...
- You're feel free to send ideas to [issue](https://github.com/create-go-app/net_http-go-template/issues/new/choose) 😉

## Developers

- Idea and active development by [Vic Shóstak](https://github.com/koddr) (aka Koddr).

## Project assistance

If you want to say «thank you» or/and support active development `Create Go App`:

1. Add a GitHub Star to project.
2. Twit about project [on your Twitter](https://twitter.com/intent/tweet?text=Set%20up%20a%20new%20Go%20%28Golang%29%20full%20stack%20app%20by%20running%20one%20CLI%20command%21%26url%3Dhttps%3A%2F%2Fgithub.com%2Fcreate-go-app).
3. Donate some money to project author via PayPal: [@paypal.me/koddr](https://paypal.me/koddr?locale.x=en_EN).
4. Join DigitalOcean at our [referral link](https://m.do.co/c/b41859fa9b6e) (your profit is **$100** and we get $25).
5. Become a sponsor.

Thanks for your support! 😘 Together, we make this project better every day.

### Sponsors

| Logo                                                                                                                                                              | Sponsor description                                                                                                                 | URL                              |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | -------------------------------- |
| <img align="center" width="100px" src="https://raw.githubusercontent.com/create-go-app/cli/master/images/sponsors/1wa.co_logo.png" alt="True web artisans logo"/> | **True web artisans** — IT specialists around the world, who are ready to share their experience to solve your business objectives. | [https://1wa.co](https://1wa.co) |
|                                                                                                                                                                   | <div align="center">💡 <a href="mailto:truewebartisans@gmail.com">Want to become a sponsor too?</a></div>                           |                                  |

## License

MIT
