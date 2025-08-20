# Temperature Converter CLI

A simple command-line application written in Go that performs temperature conversions between three major temperature scales: Celsius, Fahrenheit, and Kelvin.

## ✨ Features

- **6 complete conversion types**: All possible conversions between Celsius ↔ Fahrenheit ↔ Kelvin
- **User-friendly interface** with interactive menu system
- **Attractive UI** with emojis and ASCII art borders
- **Input validation** and menu loop for repeated use
- **Precise calculations** with proper decimal formatting

## 🎯 Available Conversions

1. Celsius to Fahrenheit
2. Fahrenheit to Celsius  
3. Celsius to Kelvin
4. Kelvin to Celsius
5. Fahrenheit to Kelvin
6. Kelvin to Fahrenheit

## 🔧 Prerequisites

- Go 1.16 or higher installed on your system
- Terminal/Command prompt access

## 🚀 How to Run

### Clone the repository:
```bash
git clone https://github.com/yourusername/temperature-converter-go.git
cd temperature-converter-go
```

### Run the program:
```bash
go run main.go
```

### Or build and run:
```bash
go build -o temp-converter main.go
./temp-converter
```

## 📖 Usage

1. Run the program
2. Choose conversion type from the menu (1-7)
3. Enter temperature value when prompted
4. View the converted result
5. Choose another conversion or exit (option 7)

## 📋 Example Output

```
========================================
||         KONVERSI SUHU v1.0         ||
||     Celsius ⇄ Fahrenheit ⇄ Kelvin  ||
========================================

Pilih jenis konversi suhu:
1. Celsius ke Fahrenheit
2. Fahrenheit ke Celsius
3. Celsius ke Kelvin
4. Kelvin ke Celsius
5. Fahrenheit ke Kelvin
6. Kelvin ke Fahrenheit
7. Keluar

Masukkan pilihan (1-7): 1

==========================================
  🌡️ Selamat Datang di Konversi Suhu 🌡️  
         Celsius ➡️ Fahrenheit            
==========================================

Input suhu(°C): 25
Hasil konversi suhu = 77°F
```

## 🛠️ Technology Stack

- **Language**: Go (Golang)
- **Standard Library**: fmt package for I/O operations
- **Platform**: Cross-platform CLI application

## 📚 Learning Objectives

This project demonstrates:
- Basic Go syntax and structure
- Function implementation and organization
- Switch-case statements
- Loop structures and flow control
- User input handling with `fmt.Scan()`
- String formatting with `fmt.Printf()`
- Type conversion (int to float64)

## 🤝 Contributing

Feel free to fork this project and submit pull requests for any improvements:
- Code optimization
- Additional temperature scales (Rankine, Réaumur, etc.)
- Better error handling
- Unit tests
- GUI version

---

⭐ If you found this project helpful, please consider giving it a star!
