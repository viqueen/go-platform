## go-platform

- This is a template for a go project backend with a domain driven design
- It is a work in progress and will be updated as I learn more about go while unlearning Java

### development setup

- proto schema generate

```bash
make schema
```

- start backend harness

```bash
make harness
```

- start web client

```bash
make webapp
```

### domains

- init

```bash
./scripts/domain.sh init <name>
```

- tidy

```bash
./scripts/domain.sh tidy <name>
```

### platform

- init

```bash
./scripts/platform.sh init <name>
```

- tidy

```bash
./scripts/platform.sh tidy <name>
```
