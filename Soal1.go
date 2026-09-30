package main
import "fmt"

func main() {
	var masukan float64
	fmt.Scan(&masukan)

	hasil := 2/(masukan+5) + 5
	fmt.Println(hasil)
