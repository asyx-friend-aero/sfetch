package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"
)

// ANSI-коди для кольорів
const (
	ColorReset  = "\033[0m"
	ColorBold   = "\033[1m"
	ColorBlue   = "\033[36m" // Arch Blue
	ColorPurple = "\033[35m" // Gentoo Purple
	ColorGreen  = "\033[32m" // Void Green
	ColorRed    = "\033[31m" // Debian Red
)

// Отримуємо назву ОС із /etc/os-release
func getOSName() (string, string) {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return "Arch Linux", "arch"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	osID := "arch"
	prettyName := "Linux"

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "ID=") {
			osID = strings.Trim(strings.TrimPrefix(line, "ID="), `"`)
		}
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			prettyName = strings.Trim(strings.TrimPrefix(line, "PRETTY_NAME="), `"`)
		}
	}

	return prettyName, osID
}

func main() {
	u, err := user.Current()
	username := "user"
	if err == nil {
		username = u.Username
	}

	hostname := "localhost"
	if h, err := os.Hostname(); err == nil {
		hostname = h
	}

	prettyOS, osID := getOSName()

	var logo []string
	var primaryColor string

	// Вибір логотипу та кольору залежно від дистрибутива
	switch osID {
	case "gentoo":
		primaryColor = ColorPurple
		logo = []string{
			ColorPurple + "      -/oyddmdhs+:." + ColorReset,
			ColorPurple + "      -odNMMMMMMMMNNmhy+-" + ColorReset,
			ColorPurple + "    -yNMMMMMMMMMMMNNNmmdhy+-" + ColorReset,
			ColorPurple + "  `omMMMMMMMMMMMMNmdmmmmddhhy/" + ColorReset,
			ColorPurple + "  omMMMMMMMMMMMNhhyyyohmdddhhhdo`" + ColorReset,
			ColorPurple + " .ydMMMMMMMMMMdhs++so/smdddhhhhdm+" + ColorReset,
			ColorPurple + "  oyhdmNMMMMMMMNdyooydmddddhhhhyhNd." + ColorReset,
			ColorPurple + "   :oyhhdNNMMMMMMMNNNmmdddhhhhhyymMh" + ColorReset,
			ColorPurple + "     .:+sydNMMMMMNNNmmmdddhhhhhhmMmy" + ColorReset,
			ColorPurple + "        /mMMMMMMNNNmmmdddhhhhhmMNhs:" + ColorReset,
			ColorPurple + "     `oNMMMMMMMNNNmmmddddhhdmMNhs+" + ColorReset,
			ColorPurple + "   `sNMMMMMMMMNNNmmmdddmNMmhs/." + ColorReset,
			ColorPurple + "  /NMMMMMMMMNNNNmmmdddmNMNdso:`" + ColorReset,
			ColorPurple + " +MMMMMMMNNNNNmmmmdmNMNdso/-" + ColorReset,
			ColorPurple + " yMMNNNNNNNmmmmmNNMmhs+/-`" + ColorReset,
			ColorPurple + " /hMMNNNNNNNNMNdhs++/-`" + ColorReset,
			ColorPurple + " `/ohdmmddhys+++/:.`" + ColorReset,
			ColorPurple + "   `-//////:--." + ColorReset,
		}
	case "void":
		primaryColor = ColorGreen
		logo = []string{
			ColorGreen + "    ____   " + ColorReset,
			ColorGreen + "  _--------_ " + ColorReset,
			ColorGreen + " / \\______/ \\ " + ColorReset,
			ColorGreen + "|  /      \\  |" + ColorReset,
			ColorGreen + "| |        | |" + ColorReset,
			ColorGreen + "|  \\______/  |" + ColorReset,
			ColorGreen + " \\__________/ " + ColorReset,
		}
	case "debian":
		primaryColor = ColorRed
		logo = []string{
			ColorRed + "  _____     " + ColorReset,
			ColorRed + " /  __ \\    " + ColorReset,
			ColorRed + "|  /    |   " + ColorReset,
			ColorRed + "|  \\___-    " + ColorReset,
			ColorRed + " -_         " + ColorReset,
			ColorRed + "   --_      " + ColorReset,
		}
	default: // Arch Linux за замовчуванням
		primaryColor = ColorBlue
		logo = []string{
			ColorBlue + "       /\\" + ColorReset,
			ColorBlue + "      /  \\" + ColorReset,
			ColorBlue + "     / /\\ \\" + ColorReset,
			ColorBlue + "    / /  \\ \\" + ColorReset,
			ColorBlue + "   / /    \\ \\" + ColorReset,
			ColorBlue + "  / /      \\ \\" + ColorReset,
			ColorBlue + " / /   --   \\ \\" + ColorReset,
			ColorBlue + "/_/          \\_\\" + ColorReset,
		}
	}

	userHost := fmt.Sprintf("%s%s\uf007 %s%s@%s%s", primaryColor, ColorBold, username, ColorReset, ColorBold, hostname)
	separator := strings.Repeat("─", len(username)+len(hostname)+4)

	// Системна інформація з іконками JetBrains Mono Nerd Font
	info := []string{
		userHost,
		primaryColor + separator + ColorReset,
		fmt.Sprintf("%s\uf303 OS:%s     %s", primaryColor, ColorReset, prettyOS),
		fmt.Sprintf("%s\uf17c Kernel:%s %s", primaryColor, ColorReset, runtime.GOOS),
		fmt.Sprintf("%s\uf4bc Arch:%s   %s", primaryColor, ColorReset, runtime.GOARCH),
		fmt.Sprintf("%s\uf4bc CPUs:%s   %d cores", primaryColor, ColorReset, runtime.NumCPU()),
		fmt.Sprintf("%s\ue626 Go Ver:%s %s", primaryColor, ColorReset, runtime.Version()),
	}

	maxLines := len(logo)
	if len(info) > maxLines {
		maxLines = len(info)
	}

	for i := 0; i < maxLines; i++ {
		logoLine := "                                      "
		if i < len(logo) {
			logoLine = logo[i]
		}

		infoLine := ""
		if i < len(info) {
			infoLine = info[i]
		}

		fmt.Printf("%-42s %s\n", logoLine, infoLine)
	}
}

