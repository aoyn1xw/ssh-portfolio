# ssh-portfolio

A terminal portfolio you can visit over SSH — built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Wish](https://github.com/charmbracelet/wish).

```
ssh <host> -p 2222
```

## Features

- About me page with bio
- Links page (GitHub, Instagram, Discord)
- 5 color themes — switch with `↑ ↓`
- Tab navigation with `← →`

## Run locally

```bash
go run main.go --local
```

## Deploy

```bash
# Generate a host key (only needed once)
mkdir -p .ssh
ssh-keygen -t ed25519 -f .ssh/id_ed25519 -N ""

# Build and run
go build -o ssh-portfolio
./ssh-portfolio
```

Server listens on port `2222` by default.

## Keybindings

| Key | Action |
|-----|--------|
| `← →` | Switch tabs |
| `↑ ↓` | Change color theme |
| `q` | Quit |

## License

[MIT](LICENSE)
