package main

import "fmt"

func convertCelsiusToFahrenheit(){
fmt.Print("\n")
fmt.Println("==========================================")
fmt.Println("  🌡️ Selamat Datang di Konversi Suhu 🌡️  ")
fmt.Println("         Celsius ➡️ Fahrenheit            ")
fmt.Println("==========================================")
fmt.Print("\nInput suhu(°C): ")
var input int
fmt.Scan(&input)
fahrenheit := (input * 9/5) + 32
fmt.Printf("Hasil konversi suhu = %d°F", fahrenheit)
}

func convertFahrenheitToCelsius(){
fmt.Print("\n")
fmt.Println("==========================================")
fmt.Println("  🌡️ Selamat Datang di Konversi Suhu 🌡️  ")
fmt.Println("         Fahrenheit ➡️ Celsius            ")
fmt.Println("==========================================")
fmt.Print("\nInput suhu(°F): ")
var input int
fmt.Scan(&input)
celcius := (input - 32) * 5/9
fmt.Printf("Hasil konversi suhu = %d°C", celcius)
}

func convertCelsiusToKelvin(){
fmt.Print("\n")
fmt.Println("==========================================")
fmt.Println("  ❄️ Konversi Suhu: Celsius ke Kelvin ❄️  ")
fmt.Println("     Menjaga suhu tetap ilmiah dan akurat ")
fmt.Println("==========================================")
fmt.Print("\nInput suhu(°C): ")
var input int
fmt.Scan(&input)
kelvin := float64(input) + 273.15
fmt.Printf("Hasil konversi suhu = %.2f°K", kelvin)
}

func convertKelvinToCelsius(){
fmt.Print("\n")
fmt.Println("==========================================")
fmt.Println("  🧊 Konversi Suhu: Kelvin ke Celsius 🧊  ")
fmt.Println("     Dari skala absolut ke keseharian     ")
fmt.Println("==========================================")
fmt.Print("\nInput suhu(°K): ")
var input int
fmt.Scan(&input)
celcius := float64(input) - 273.15
fmt.Printf("Hasil konversi suhu = %.2f°C", celcius)
}

func convertFahrenheitToKelvin(){
fmt.Print("\n")
fmt.Println("==========================================")
fmt.Println("  🔥 Konversi Suhu: Fahrenheit ke Kelvin 🔥")
fmt.Println("     Saat panas Amerika bertemu sains     ")
fmt.Println("==========================================")
fmt.Print("\nInput suhu(°F): ")
var input int
fmt.Scan(&input)
kelvin := float64(input - 32) * 5/9 + 273.15
fmt.Printf("Hasil konversi suhu = %.2f°K", kelvin)
}

func convertKelvintoFahrenheit(){
fmt.Print("\n")
fmt.Println("==========================================")
fmt.Println("  ❄️ Konversi Suhu: Kelvin ke Fahrenheit ❄️")
fmt.Println("     Dari suhu mutlak ke gaya barat       ")
fmt.Println("==========================================")
fmt.Print("\nInput suhu(°K): ")
var input int
fmt.Scan(&input)
fahrenheit := (float64(input) - 273.15) * 9/5 + 32
fmt.Printf("Hasil konversi suhu = %.2f°F", fahrenheit)
}

func mainMenu(){
	fmt.Println("Pilih jenis konversi suhu:")
	fmt.Println("1. Celsius ke Fahrenheit")
    fmt.Println("2. Fahrenheit ke Celsius")
    fmt.Println("3. Celsius ke Kelvin")
    fmt.Println("4. Kelvin ke Celsius")
    fmt.Println("5. Fahrenheit ke Kelvin")
    fmt.Println("6. Kelvin ke Fahrenheit")
    fmt.Println("7. Keluar")
	for {
	fmt.Print("\nMasukkan pilihan (1-7): ")
	var pilihan int
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
        convertCelsiusToFahrenheit()
	case 2:
        convertFahrenheitToCelsius()
	case 3:
        convertCelsiusToKelvin()
	case 4:
        convertKelvinToCelsius()
	case 5:
        convertFahrenheitToKelvin()
	case 6:
        convertKelvintoFahrenheit()
	case 7:
		fmt.Println("Terimakasih sudah menggunakan program.")
		return

	default:
		fmt.Println("Invalid! Pilihan tidak ada di menu. Pilih (1-7).")
	}
}
}

func main() {
fmt.Println("========================================")
fmt.Println("||         KONVERSI SUHU v1.0         ||")
fmt.Println("||     Celsius ⇄ Fahrenheit ⇄ Kelvin  ||")
fmt.Println("========================================")
fmt.Println("")
mainMenu()
}