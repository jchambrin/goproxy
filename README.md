# goproxy

Reverse proxy with caching

## Quickstart

To start your reverse proxy, you only need to provide a YAML configuration file.
```yaml
name: reverse1
source: localhost:3000
destination:
  protocol: https
  host: github.com
  port: 443
cache:
  ttl: 60
  enable: true
  allowedMethods: 
    - GET
    - HEAD
```
Some examples can be found in the `data` folder

You can then, build and start the proxy
```bash
go build && ./goproxy -config ./data/github-config.yaml
```