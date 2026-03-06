package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ssh-portfolio/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/muesli/termenv"
)

const (
	host = "0.0.0.0"
	port = "2222"
)

func main() {
	// Local mode: go run main.go --local
	if len(os.Args) > 1 && os.Args[1] == "--local" {
		p := tea.NewProgram(tui.NewModel(), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			log.Fatalf("Error running local TUI: %v", err)
		}
		return
	}

	s, err := wish.NewServer(
		wish.WithAddress(host+":"+port),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			bubbletea.Middleware(teaHandler),
		),
	)
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("Starting SSH server on %s:%s", host, port)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-done

	log.Println("Stopping server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, active := s.Pty()
	if !active {
		m := tui.NewModel()
		return m, []tea.ProgramOption{tea.WithAltScreen()}
	}

	renderer := lipgloss.NewRenderer(s)
	renderer.SetColorProfile(termenv.TrueColor)
	m := tui.NewModelWithRenderer(renderer)

	return m, []tea.ProgramOption{
		tea.WithAltScreen(),
		tea.WithInput(s),
		tea.WithOutput(s),
		tea.WithEnvironment([]string{"TERM=" + pty.Term}),
	}
}
