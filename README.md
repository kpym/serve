# serve

Dead simple localhost server.

Start serving the current folder at the first available port after `8080` and open the browser to the corresponding adress.

```bash
> serve
2021/05/05 11:18:58 Start serving the current folder to localhost:8080.
```

## The go code

1. Find the first available port:
    ```go
      var hostport string
      for i := 8080; i < 8180; i++ {
        hostport = fmt.Sprintf("localhost:%d", i)
        if ln, err := net.Listen("tcp", hostport); err == nil {
          ln.Close()
          break
        }
      }
    ```
1. Open the browser at `url :="http://" + hostport`:
    - on Windows
      ```go
        exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
      ```
    - on Linux
      ```go
        exec.Command("xdg-open", url).Start()
      ```
    - on Mac
      ```go
        exec.Command("open", url).Start()
      ```
1. Start serving the local folder:
    ```go
      http.ListenAndServe(hostport, http.FileServer(http.Dir(".")))
    ```

