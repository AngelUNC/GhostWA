package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var debugMode = false

func Start() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt())
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		if cmd == "" {
			continue
		}

		handle(cmd)
	}
}

func prompt() string {
	if debugMode {
		return "dev> "
	}
	return "> "
}

func handle(cmd string) {
	switch {

	case cmd == "help":
		fmt.Println(`
🧭 Comandos disponibles:
 help              → Muestra esta ayuda
 clear             → Limpia la pantalla
 close             → Cierra el monitor
 developer-mode    → Activa o desactiva debug mode
`)
	case cmd == "clear":
		fmt.Print("\033[2J\033[H")

	case cmd == "close":
		fmt.Println("👋 Cerrando GhostWA…")
		os.Exit(0)

	case cmd == "developer-mode":
		debugMode = !debugMode
		fmt.Printf("🛠️ Modo desarrollador: %v\n", debugMode)

	default:
		fmt.Printf("❓ Comando desconocido: %s (usa 'help')\n", cmd)
	}
}
