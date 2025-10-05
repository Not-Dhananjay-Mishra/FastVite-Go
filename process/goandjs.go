package process

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func GoAndJS(name string) {
	// Create main project folder
	err := os.Mkdir(name, 0755)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	os.Chdir(name)

	// Backend
	fmt.Println("Making Go backend...")
	cmd := exec.Command("cmd.exe", "/K", "go mod init "+name)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Backend structure
	os.Create("main.go")
	os.Mkdir("api", 0755)
	os.Mkdir("database", 0755)
	os.Mkdir("utils", 0755)
	os.Create(".env")
	os.Create("README.md")
	os.Create(".gitignore")

	// Frontend
	fmt.Println("Making JS frontend...")
	viteCmd := exec.Command("cmd.exe", "/K", "npm create vite@latest frontend -- --template react")
	if err := viteCmd.Run(); err != nil {
		fmt.Println("Error:", err)
		return
	}
	npmi := exec.Command("cmd.exe", "/K", "cd frontend && npm install")
	if err := npmi.Run(); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Adding Tailwind CSS...")
	tailwindCmd := exec.Command("cmd.exe", "/K",
		"cd frontend && npm install -D tailwindcss postcss autoprefixer && npx tailwindcss init -p")
	if err := tailwindCmd.Run(); err != nil {
		fmt.Println("Error installing Tailwind:", err)
		return
	}

	// Update tailwind.config.js
	configContent := `/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}`
	err = os.WriteFile("frontend/tailwind.config.js", []byte(configContent), 0644)
	if err != nil {
		fmt.Println("Error writing tailwind.config.js:", err)
	}

	// Update index.css (or main.css)
	cssContent := `@tailwind base;
@tailwind components;
@tailwind utilities;`
	err = os.WriteFile("frontend/src/index.css", []byte(cssContent), 0644)
	if err != nil {
		// sometimes Vite uses main.css instead of index.css
		_ = os.WriteFile("frontend/src/main.css", []byte(cssContent), 0644)
	}

	fmt.Println("Project setup complete with Go backend + Vite + Tailwind frontend!")
	fmt.Println("Starting VS code...")
	time.Sleep(2 * time.Second)
	vscodeCmd := exec.Command("cmd.exe", "/K", "code .")
	if err := vscodeCmd.Run(); err != nil {
		fmt.Println("Error:", err)
		return
	}
}
