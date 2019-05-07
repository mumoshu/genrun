# genrun

Free developers from config-file-format lock-ins.

---

**gen**run generates some files then **run** a command.

The following fiile generators are supported:

- Lua 5.1 (powered by [gopher-lua](https://github.com/yuin/gopher-lua)
- Jsonnet (powered by [go-jsonnet](https://github.com/google/go-jsonnet)
- Go templates with various datasources (powered by [gomplate](https://github.com/hairyhenderson/gomplate))
- YAML/JSON/TOML

## Usage

```console
$ genrun Genrunfile -- helmfile sync
```

Given the following files, `genrun` generates the `helmfile.yaml` and `.envrc` evaluating Lua and Gomplate code according to `Genrunfile`, and then runs `helmfile sync` according to the command-line args after `--`.

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

## Shebang support

`genrun` supports the usage from shebangs. That is, add `genrun` on top of your genrun config file:

`bin/helmfile`:

```
#!/usr/bin/env genrun

files:
- source: .envrc.gotmpl
- source: helmfile.yaml.lua
```

Now make it an executable so that it behaves as a native commnad:

```
$ bin/helmfile sync

# `genrun` generates files from the definitions in `bin/helmfile` and then runs `helmfile` according to the basename of $0(=helmfile)
```

In case the command being wrapped by `genrun` must be customized, set `command` in your config file:

```
#!/usr/bin/env genrun

command:
- envexec
- helmfile

files:
- source: .envrc.gotmpl
- source: helmfile.yaml.lua
```
