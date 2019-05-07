# genrun

Generate some files and envvars then run a command.

## Usage

```console
$ genrun Genrunfile -- helmfile sync
```

Given the following files, `gnerun` generates the `helmfile.yaml` and `.envrc` according to `Genrunfile`, and then runs `helmfile sync` according to the command-line args after `--`.

`Genrun.yaml`:

```yaml
files:
- .envrc.gotmpl # generates .envrc by rendering go text/template
- helmfile.yaml.lua # generates helmfile.yaml by evaluating the lua script

## Alternative syntax

files:
- source: .genrun/.envrc.gotmpl
  target: .envrc
- source: .genrun/helmfile.yaml.lua
  target: helmfile.yaml
```
