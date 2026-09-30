package main
import "fmt"

func main() {
	const pi = 3.1415926535

	var r int
	fmt.Print("Jejari = ")
	fmt.Scan(&r)

	rf := float64(r)
	volume := 4.0 / 3.0 * pi * rf * rf * rf
	luas := 4 * pi * rf * rf

	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}
