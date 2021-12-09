# xumm-go-client

![example workflow](https://github.com/xyield/xumm-go-client/actions/workflows/main.yml/badge.svg)

## Installation:
The latest version can be installed using `go get` :

```bash
go get github.com/xyield/xumm-go-client
```


## How to use xYield

Register your app with the Xumm developer console
Set Xumm credentials as environment variables

```bash
export XUMM_API_KEY=<key_from_console>
export XUMM_API_SECRET=<secret_from_console>
```

```go
func main() {
    c, err := xumm.NewClient()

    if err != nil {
        fmt.Println(err)
    }

    p, err := c.GetPing()

    if err != nil {
        // Handle error here
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
