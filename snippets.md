# Snippets

## Extender

### v1

```yaml
extenders:
- urlPrefix: http://extender:8000/
  filterVerb: filter
  weight: 10
```

```yaml
name: glinda
  annotations:
    scheduler.wasmkwokwizardry.io/regex: 'oz-.*'
```

### v2

```yaml
nodeCacheCapable: true
```

```yaml
name: elphaba
  annotations:
    scheduler.wasmkwokwizardry.io/regex: 'oz-.*'
```

## Framework

### v1

```sh
go get github.com/wasmkwokwizardry/kubecon-na-2024/scheduling-framework-regex-plugin
```

```go
regex "github.com/wasmkwokwizardry/kubecon-na-2024/scheduling-framework-regex-plugin/v1/plugin"
```

```go
regex.Name: regex.New,
```

```yaml
  - name: RegexScheduling
```

```yaml
name: albus
  annotations:
    scheduler.wasmkwokwizardry.io/regex: 'phoenix-.*'
```

### v2

```go
regex "github.com/wasmkwokwizardry/kubecon-na-2024/scheduling-framework-regex-plugin/v2/plugin"
```

```yaml
  - name: RegexScheduling
```

```yaml
name: sirius
  annotations:
    scheduler.wasmkwokwizardry.io/regex: 'phoenix-.*'
```

### WASM

### v1

```yaml
  - name: WasmRegexScheduling
```
```yaml
  - name: WasmRegexScheduling
        args:
          guestURL: http://static-webserver/regex_v1.wasm
```

```yaml
name: morrible
  annotations:
    scheduler.wasmkwokwizardry.io/regex: 'oz-.*'
```

### v2

```yaml
              guestURL: http://static-webserver/regex_v2.wasm
```

```yaml
name: minerva
  annotations:
    scheduler.wasmkwokwizardry.io/regex: 'phoenix-.*'
```