package main

import "fmt"

func main() {
	// fungsi closure
	/*
		adalah suatu fungsi anonymous function (fungsi tanpa nama)
		yang disimpan dalam variable, dengan adanya closure, biasanya
		digunakan untuk membuat fungsi di dalam fungsi atau bahkan membuat fungsi
		yang mengembalikan fungsi,

		closure dimanfaatkan untuk membungkus suatu proses yang hanya dijalankan sekali saja atau hanya dipakai pada blok
		tertentu saja
	*/
	// Closure disimpan sebagai variable
	/*
		sebuah fungsi tanpa nama bisa disimpan dalam variable closure yang memiliki sifat seperti fungsi
		yang disimpannya

		berikut contoh program sederhana yang menerappkan closure untuk pencarian nilai terendah dan tinggi dari data
		arra, logika pencarian dibungkus dalam closure yang ditampung oleh variable
	*/
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				min, max = e, e
			case e < min:
				min = e
			case e > max:
				max = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)

	// closure jenis IIFE (Immediately Invoked Function Expression)
	/*
		eksekusinya adalah langsung saat deklarasi, Teknik ini biasa diterapkan untuk membungkus proses
		yang hanya dilakukan sekali IIFE bisa memiliki nilai baik atau juga tidak
	*/
	var numbers2 = []int{2, 3, 4, 3, 4, 2, 3, 1, 3, 2, 1, 5, 6, 7, 543, 34, 21}

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers2 {
			if e < min {
				continue
			}
			r = append(r, e)

		}
		return r
	}(5) // IIFE
	// IIFE itu untuk memberikan output yang dimulai dari minimal 5

	fmt.Println("original number :", numbers2)
	fmt.Println("filtered number :", newNumbers)

	// ciri khas dari penulisan IIFE adalah adanya tanda kurung parameter yang ditulis di akhir deklarasi closure
	//jika IIFE memiliki parameter, maka argumentnya juga ditulis
	/*
		misal var newNumbers = func(min int) []int {
			// ,,,,,
		}(5)
	*/
	/*
			menghasilkan nilai balik yang ditampung variable
			newNumber, Perlu diperhatikan bahwa yang ditampung adalah nilai kembaliannya bukan body fungsi atau closurenya
			Closure bisa juga dengan gaya manifest typing, caranya dengan menuliskan skema closure-nya sebagai tipe data. Contoh:
			var closure (func (string, int, []string) int)
			closure = func (a string, b int, c []string) int {
		    // ..
		}
	*/

	// Closure sebagai Nilai Kembalian
	/*
		salah satu keunikan lain dari closure adalah
		closure bisa dijadikan sebagai nilai balik fungsi
		sedikit aneh, tapi pada kondisi tertentu ini bisa sangat berguna
		contoh findMax() yang salah satu nilai kembaliannya adalah berupa
		closure
	*/
	var max11 = 3
	var numbers11 = []int{2, 3, 4, 3, 4, 2, 3, 1, 3, 2, 1, 5, 6, 7, 543, 34, 21}
	var howMany, getNumbers = findMax(numbers11, max11)
	var theNumbers = getNumbers()

	fmt.Println("numbers :", numbers11)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("how many \t:", howMany)
	fmt.Println("the numbers \t:", theNumbers)

	// fungsi sebagai parameter
	/*
		sebuah fungsi yang memiliki parameter sebuah fungsi

		pada Go fungsi bisa dijadikan sebagai tipe data variable maka sangat memungkinkan
		untuk menjadikannya sebagai parameter

	*/
	// Penerapan fungsi sebagai parameter
	// dengan menuliskan skema fungsi nya sebagai tipe data

}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

// CALLBACK merupakan sebuah closure yang dideklarasikan bertipe func(string) bool
// closure tersebut dipanggil di tiap perulangan dalam fungsi filter()
/*
fungsi filter()
sendiri dibuat untuk filtering data array ( yang datanya didapat dari parameter pertama)
dengan kondisi
*/

// contoh fungsi closure sebagai nilai kembalian
func findMax(numbers []int, max int) (int, func() []int) {
	var result []int
	for _, e := range numbers {
		if e <= max {
			result = append(result, e)
		}
	}
	return len(result), func() []int {
		return result
	}
}

/*
nilai kembalian ke-2 pada fungsi di atas adalah closure dengan skema func() []int
Bisa dilihat di bagian akhir
return len(res), func() []int{
	return res
}
Fungsi tanpa nama yang akan dikembalikan boleh disimpan pada variabel terlebih dahulu. Contohnya:
var getNumbers = func() []int {
    return res
}
return len(res), getNumbers

Tentang fungsi findMax() sendiri, fungsi ini dibuat untuk mempermudah pencarian angka-angka yang nilainya di bawah atau sama dengan angka tertentu. Fungsi ini mengembalikan dua buah nilai balik:

Nilai balik pertama adalah jumlah angkanya.
Nilai balik kedua berupa closure yang mengembalikan angka-angka yang dicari.

*/
