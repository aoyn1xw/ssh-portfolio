# ssh-portfolio

A terminal-based portfolio accessible via SSH, built with Go and Charm.sh.

## Connect
```bash
ssh hi.ayon1xw.me -p 2222
```

## Built With

- [Wish](https://github.com/charmbracelet/wish) — SSH server framework
- [Bubbletea](https://github.com/charmbracelet/bubbletea) — Terminal UI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) — Terminal styling

## Features

- About me page
- Links / contact page
- Theme switcher (press `t`)
- Keyboard navigation (`← →`)

## Running Locally
```bash
git clone https://github.com/aoyn1xw/ssh-portfolio.git
cd ssh-portfolio
go run main.go
```

Then in another terminal:
```bash
ssh localhost -p 2222
```
