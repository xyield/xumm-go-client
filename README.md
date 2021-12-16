# xumm-go-client

![example workflow](https://github.com/xyield/xumm-go-client/actions/workflows/main.yml/badge.svg)

This is a client written in golang for using the XUMM API. The API docs can be found [here](https://xumm.readme.io/reference/about).
## Installation:
The latest version can be installed using `go get` :

```bash
go get github.com/xyield/xumm-go-client@master
```

## How to use this client

Register your app with the Xumm developer console
Set Xumm credentials as environment variables

```bash
export XUMM_API_KEY=<key_from_console>
export XUMM_API_SECRET=<secret_from_console>
```

```go
package main

func main() {
    cfg, err := xumm.NewConfig()
	if err != nil {
		log.Panicln(err)
	}

	client := client.New(cfg)

	pong, err := client.Meta.Ping()

	if err != nil {
		log.Println(err.Error())
	}

    // Do something
}
```


## Future work
Continue to develop all endpoints
</br>
Integration testing in pipeline to connect to Xumm API
</br>
Ability to use custom logging library
</br>
Ability to use a custom client library
</br>
