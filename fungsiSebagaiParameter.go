package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{"wick", "jason", "ethan"}
	var dataContains0 = saring(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLength5 = saring(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t:", data)

	fmt.Println("saring ada huruf \"0\"\t:", dataContains0)
	fmt.Println("jumlah huruf \"5\"\t:", dataLength5)

	// pointer meerupaka refrence atau alamat memori variable

}

// penerapan fungsi sebagai parameter
/*
cara membuat parameter fungsi adalah dengan langsung menuliskan skema fungsi
sebagai tipe data
*/
// misal
func saring(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}

	}
	return result
}

/*
parameter callback merupakan sebuah closure yang dideklarasikan bertipe func(string) bool
Closure tersebut dipanggil di tiap perulangan dalam fungsi saring()

Fungsi saring() dibuat untuk filtering data array yang datanya didapat dari parameter pertama)
dengan kondisi saring bisa ditentukan sendiri Di bawah ini merupakan contoh pemanfaatan fungsi

ada banyak yang terjadi dari pemanggilan fungsi saring()
1. Data array yang didapat dari parameter pertama akan diloooping
2. tiap perulangan closure callback dipanggil dengan disisipkan data tiap elemen
perulangan sebagai parameter
3. closure callback berisikan kondisi filtering, dengan hasil yang bertipe bool yang kemudian dijadikan nilai balik dikembalikan
dalam fungsi filter() sendiri ada proses seleksi kondisi yang nilainya didapat dari hasil eksekusi closure callback
ketika kondisinya bernilai true maka data elemen yang sedang diulang dinyatakan lolos proses filtering
5. data yang lolos ditampung variable result variable tersebut dijadikan sebagai nilai balik fungsi filter()

pada
*/
