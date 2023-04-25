# serve

Dead simple localhost server.

Start serving the current folder at the first available port after `8080` and open the browser to the corresponding adress.

```bash
> serve
2022/07/07 14:31:44 [serve] Start serving the current folder at http://localhost:8080.
2022/07/07 14:31:44 [serve] Request: /
```

Or if there is piped data it will be served.

```bash
> echo test | serve
2022/07/07 14:33:46 [serve] Piped data is present.
2022/07/07 14:33:46 [serve] Start serving the current folder at http://localhost:8080.
2022/07/07 14:33:55 [serve] Request: /
2022/07/07 14:33:55 [serve] Serve piped data at root.
```
Only two flags are available:

- `-p` : choose the port to use. If not set the first available after `8080` will be used.
  ```bash
  > serve -p 8888
  2022/07/07 14:34:56 [serve] Start serving the current folder at http://localhost:8888.
  2022/07/07 14:34:56 [serve] Request: /
  ```
- `-t` : serve at the given path
  ```bash
  > serve -t 'foo'
  2022/07/07 14:36:54 [serve] Start serving at http://localhost:8080/foo/.
  2022/07/07 14:36:54 [serve] Request: /foo/
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
goreleaser --snapshot --skip-publish --clean
```

You will find the resulting binaries in the `dist/` sub-folder.

## License

[MIT](LICENSE) for this code _(but all used libraries may have different licences)_.
