# Go + Vite + Tailwind CLI Tool

A simple **CLI tool** written in Go that scaffolds a full-stack project with a **Go backend** and a **Vite + React + Tailwind frontend**.

---

## 🚀 What It Does

* Creates a new project folder with the given name
* Initializes a **Go backend**

  * Runs `go mod init`
  * Creates common folders: `api/`, `database/`, `utils/`
  * Adds starter files: `main.go`, `.env`, `.gitignore`, `README.md`
* Sets up a **frontend**

  * Runs `npm create vite@latest` with the React template
  * Installs and configures **Tailwind CSS** + PostCSS + Autoprefixer
  * Updates `tailwind.config.js` and CSS files automatically

---
## ⚙️ Installation

You can install the CLI in two ways:

1. Download the binary (Windows)

Download the latest prebuilt binary from GitHub:

`curl -L -o govite.exe https://github.com/Not-Dhananjay-Mishra/FastVite-Go/raw/main/govite.exe`


Then run it directly:

./govite.exe myapp

2. Install with Go

If you have Go installed, you can install it globally:

`go get github.com/Not-Dhananjay-Mishra/FastVite-Go`
`cd FastVite-Go`
`go mod tidy`
`go install`


This will place the govite binary in your $GOPATH/bin (make sure it’s in your PATH).
---

## 📦 Example Project Structure

```
<project-name>/
│── api/             # API endpoints
│── database/        # Database layer
│── utils/           # Utility functions
│── main.go          # Go entry point
│── .env             # Environment variables
│── .gitignore
│── README.md
│── frontend/        # React + Vite + Tailwind app
```

---

## ⚙️ Installation

Clone this repo and build the CLI:

```bash
git clone <your-repo-url>
cd <your-repo-name>
go build -o create-app
```

---

## 🎮 Usage

Run the CLI and pass your project name:

```bash
./govite
```

This will:

1. Create a `myapp/` directory
2. Scaffold Go backend files
3. Scaffold `frontend/` with Vite + React
4. Install & configure Tailwind automatically

---

## 📖 Example

```bash
./govite
```

Output:

```
Making Go backend...
Making JS frontend...
Adding Tailwind CSS...
✅ Project setup complete with Go backend + Vite + Tailwind frontend!
```

---

## 🛠 Requirements

* Go 1.20+
* Node.js + npm
* Git (optional but recommended)

---

## 🤝 Contributing

Pull requests are welcome! If you find a bug or have a feature request, open an issue.

