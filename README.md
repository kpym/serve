# serve

Dead simple localhost server.

Start serving the current folder at the first available port after `8080` and open the browser to the corresponding adress.

```bash
> serve
2021/12/22 15:24:57 serve [dev]: start serving the current folder to localhost:8080.
```

Or if there is piped data it will be served.

```bash
> echo test | serve
2021/12/22 15:24:20 serve [dev]: start serving the piped data to localhost:8080.
2021/12/22 15:24:21 serve from stdin
```

## To install

### Using go

```bash
> go install github.com/kpym/serve@latest
```

### Downloading precompiled binaries

You can download the precompiled single executable from [the realease page](https://github.com/kpym/serve/releases).

