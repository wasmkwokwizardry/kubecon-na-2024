# Snippets

## Extender

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

## Framework

```sh
go get github.com/wasmkwokwizardry/kubecon-na-2024/scheduling-framework-regex-plugin
```

```
github.com/wasmkwokwizardry/kubecon-na-2024/scheduling-framework-regex-plugin v0.0.0-20241114191727-7386f4e5bea3
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
name: elphaba
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