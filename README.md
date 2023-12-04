# ko2dockerfile

Easily generate `Dockerfile` from `ko` "commands" and configuration

## Idea

- Extract images from yamls (or do like `ko build` or both)
- Read =KO_CONFIG_DIR=
- Add possibility to add extras (through configuration)
  - annotations (or labels)
  - envs
  - runs (=RUN=)

Keep it that simple. Make downstream /generate/ that configuration.
Almost /drop-in/ replacement to =ko= command.

```bash
ko apply -f -R config # …
lord generate -f -R config # … or apply
lord generate --output dockerfiles -f -R config # … or apply
# …
ko build ./cmd/foo
lord generate ./cmd/foo # output a dockerfile
lord build ./cmd/foo # build using dockerfile
```
