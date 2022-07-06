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
Only two flags are available: 

- `--port` : choose the port to use. If not set the first available after `8080` will be used.
  ```bash
  > serve --port 8888
  2021/12/22 15:24:57 serve [dev]: start serving the current folder to localhost:8888.
  ```
- `-prefix` : serve to the specified prefix folder of localhost. 
  ```bash
  > serve --prefix 'foo'
  2021/12/22 15:24:57 serve [dev]: start serving the current folder to localhost:8080/foo/.
  ```

## To install

### Using go

```bash
> go install github.com/kpym/serve@latest
```

### Downloading precompiled binaries

You can download the precompiled single executable from [the realease page](https://github.com/kpym/serve/releases).

### Using goreleaser

After cloning this repo you can compile the sources with [goreleaser](https://github.com/goreleaser/goreleaser/) for all available platforms:

```
git clone https://github.com/kpym/lol.git .
goreleaser --snapshot --skip-publish --rm-dist
```

You will find the resulting binaries in the `dist/` sub-folder.

## License

[MIT](LICENSE) for this code _(but all used libraries may have different licences)_.
