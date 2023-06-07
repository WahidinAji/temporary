package main

func main() {
	NewCar()
	NewPerson()
	NewAnu()
}

func NewCar() *Car { return &Car{} }
func NewPerson() *Person { return &Person{} }
func NewAnu() *Anu { return &Anu{} }

type Car struct {}
type Person struct{}
type Anu struct{}




// func Obj() Car {
// 	return Car{data: 1}
// }

// func CreateObj() *Car {
// 	return &Car{data: 1}
// }

// func ManyObj() []Car {
// 	return []Car{{data: 1}, {data: 2}, {data: 3}}
// }

// func OneMilConcurrent() {
// 	var wg sync.WaitGroup
// 	var mu sync.Mutex
// 	wg.Add(500)
// 	for i := 0; i < 500; i++ {
// 		mu.Lock()
// 		go func() {
// 			defer wg.Done()
// 			fmt.Println("Hello")
// 		}()
// 		mu.Unlock()
// 	}
// 	wg.Wait()
// }

// // Namun, jika objek yang Anda buat relatif besar dalam ukuran atau kompleksitasnya, atau objek tersebut perlu diakses atau dimodifikasi oleh pihak lain setelah fungsi selesai, maka mengembalikan objek sebagai pointer akan lebih efisien dalam hal penggunaan memori
// func CreateObj() *Car {
// 	return &Car{data: 1}
// }
