# genrun

Generate some files then run a command.

Supports following backends for generating files:

- Lua 5.1 (powered by [gopher-lua](https://github.com/yuin/gopher-lua)
- Jsonnet (powered by [go-jsonnet](https://github.com/google/go-jsonnet)
- Go templates with various datasources (powered by [gomplate](https://github.com/hairyhenderson/gomplate))
- YAML/JSON/TOML

## Usage

```console
$ genrun Genrunfile -- helmfile sync
```

Given the following files, `gnerun` generates the `helmfile.yaml` and `.envrc` according to `Genrunfile`, and then runs `helmfile sync` according to the command-line args after `--`.

`Genrun.yaml`:

```yaml
files:
- source: .envrc.gotmpl # generates .envrc by rendering go text/template
- source: helmfile.yaml.lua # generates helmfile.yaml by evaluating the lua script

## Advanced configuration

files:
- source: .genrun/.envrc.gotmpl
  target: .envrc
  datasources:
  - name: cities
    url: env:///CITIES?type=application/yaml
  - name: weather
    url: https://wttr.in/?0
    header: "weather=User-Agent: curl"
- source: .genrun/helmfile.yaml.lua
  target: helmfile.yaml
```
