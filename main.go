package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// ini komentar untuk satu baris
	fmt.Println("Hello, World!")
	// output dari ini adalah Hello, World!
	fmt.Println("hello", "worlds", "this is it", "I will Learn Golang")
	/* ini Komentar Multi line
	   output dari code di atas adalah hello worlds this is it I will Learn Golang
	*/

	// ini variable
	var firstName = "ayub"
	var lastName string
	lastName = "Subagiya"
	fmt.Printf("Halo %s %s!\n", firstName, lastName)
	fmt.Printf("halo %s %s!\n", lastName, firstName)
	fmt.Println("Hello!", firstName, lastName+"!")

	// deklarasi variable tanpa tipe data
	var firstName1 = "ayub subagiya"
	lastName1 := "pardjo"
	fmt.Printf("Hello %s %s!\n", firstName1, lastName1)

	// menggunakan var, tanpa tipe data, menggunakan perantara =
	//var firstName = "john"
	// tanpa var tanpa tipe data, menggunakan perantara :=
	//lstName := "doe"
	// tanda := hanya digunakan sekali di awal deklarasi variable
	/* lastname := "ayub"
	lastname = "subagiya"
	lastname = "pardjo"
	*/

	// deklarasi multiple variable
	// dalam golang deklarasi banyak variable secara bersamaan harus dengan tanda ","
	var ini, contoh, variable, multiple string
	ini, contoh, variable, multiple = "ini", "contoh", "variable", "multiple"
	fmt.Println(ini, contoh, variable, multiple)

	// pengisian nilai juga bisa dilakukan bersamaan saat deklarasi  misal
	//var satu, dua, tiga string= "satu", "dua", "tiga"
	// atau
	satu, dua, tiga := "satu", "dua", "tiga"
	fmt.Println(satu, dua, tiga)

	// dengan teknik type interface deklarasi multi variable bisa dilakukan untuk
	// berbagai tipe data misal
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	fmt.Printf("%d %t %f %s\n", one, isFriday, twoPointTwo, say)

	// jadi di go tidak boleh ada variable yang tidak digunakan artinya variable yang dideklarasikan harus digunakan
	// jika tidak digunakan maka akan error
	// Underscore atau (_) adalah reserverd variable yang bisa digunakan untuk menanmpung nilai yang tidak digunakan
	// bisa disebut variable dummy atau keranjang sampah
	// (_) biasanya digunakan untk menampung nilai balik fungsi yang tidak digunakan
	// jadi setiap daya yang masuk ke dalam variable _ akan diabaikan dan hilang layaknya ditelan black hole
	_ = "Ini variable dummy"
	_ = "sampah disini:"
	name, _ := "ayub", "subagiya"
	fmt.Println(name)

	// Deklarasi variable dengan keyword new
	// Fungsi new() digunakan untuk membuat variable pointer dengan tipe data tertentu,
	// Nilai data defaultnya sesuai dengan tipe data yang dideklarasikan
	// misal
	name1 := new(string)
	name2 := new(int)

	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(*name1)
	/* variable name1 menampung data bertipe pointer string dan name2 menampung data bertipe pointer int
	outputnya berupa alamat memori nilai tersebut dalam bentuk notasi heksadesimal, untuk menampilkan nilai aslinya
	variable tersebut harus di dereference terlebih dahulu dengan operator * seperti *name1 sebelum variable name1
	*/
	// Fungsi Make() digunakan untuk membnuat beberapa jenis variable saja seperti
	// channel, slice, dan map

	fmt.Println("")

	// tipe data di golang
	// tipe data dasar di golang terdiri dari 4 kelompok yaitu
	// 1. tipe data numerik (desimal dan nondesimal)
	// 2. tipe data string
	// 3. tipe data boolean
	// 4. tipe data pointer, array, slice, map, channel, struct, interface, dan function

	// Tipe data numerik non-desimal terbagi menjadi dua yaitu
	// Uint tipe data untuk bilangan cacah (bilangan bulat positif)
	// Int tipe data untuk bilangan bulat (bilangan bulat positif dan negatif)
	/*
			Tipe data	Cakupan bilangan
			uint8	0 ↔ 255
			uint16	0 ↔ 65535
			uint32	0 ↔ 4294967295
			uint64	0 ↔ 18446744073709551615
			uint	sama dengan uint32 atau uint64 (tergantung nilai)
			byte	sama dengan uint8
			int8	-128 ↔ 127
			int16	-32768 ↔ 32767
			int32	-2147483648 ↔ 2147483647
			int64	-9223372036854775808 ↔ 9223372036854775807
			int	sama dengan int32 atau int64 (tergantung nilai)
			rune	sama dengan int32
		Dianjurkan untuk tidak sembarangan dalam menentukan tipe data variable, karena akan mempengaruhi penggunaan memori
		pemilihan untuk pemakaian tipe data yang sesuai dengan kebutuhan agar optimal dalam penggunaan memori tidak terlalu besar
	*/

	var positveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positveNumber)
	fmt.Printf("Bilangan negatif: %d\n", negativeNumber)

	// tipe data numerik desimal
	fmt.Println("")
	fmt.Println("tipe data numerik desimal")
	// terdapat dua tipe data numerik desimal yaitu float32 dan float64
	var decimalNumber = 2.3321

	fmt.Printf("bilangan desimal %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)
	fmt.Println("")

	// tipe data bool (Boolean) tipe data ini berisikan 2 variasi nilai yaitu true dan false, tipe data
	// ini dimanfaatkan untuk mengecek kondisi dalam program

	var exist = true
	exist1 := false

	fmt.Println(exist)
	fmt.Println(exist1)
	fmt.Println("")

	// tipe data string ciri khas dari tipe data ini adalah di apit
	// oleh tanda kutip dua (") atau tanda kutip satu (')
	// misal
	message := "Ayub Subagiya welcome"
	fmt.Printf("Hello %s!\n", message)
	// selain menggunakan tanda kutip dua (") atau tanda kutip satu (') bisa juga menggunakan backtick (`) untuk
	// membuat string multiline
	// misal
	pesan := `Hello "ini penting"!, Nama saya "Ayub Subagiya", saya tinggal di "Bekasi"`
	fmt.Printf("ini pesan %s\n", pesan)
	fmt.Println("")

	// nilai nil dan zero Value
	// nil bukan merupakan tipe data melainkan sebuah nilai yang digunakan untuk merepresentasikan nilai nol atau kosong
	// semua tipe memiliki nilai nil masing-masing (default value)

	// Konstanta adalah jenis variable yang nilainya tidak bisa diubah setelah pertama dideklarasikan
	// Konstanta dideklarasikan dengan kata kunci const
	fmt.Println("")
	const firstName3 = "Ayub"
	const lastName3 = "Subagiya"
	//firstName3 := "Ayub"
	//lastName3 := "Subagiya"
	// firstName3 = "Ayub" // error karena nilai konstanta tidak bisa diubah
	// lastName3 = "Subagiya" // error karena nilai konstanta tidak bisa diubah
	fmt.Printf("hello %s %s!\n", firstName3, lastName3)

	// Deklarasi Multi konstanta
	// seperti variable, konstanta juga bisa dideklarasikan secara bersamaan
	// contoh
	const (
		square           = "kotak"
		circle           = "lingkaran"
		triangle         = "segitiga"
		numeric  uint8   = 100
		floatNum float64 = 3.14
		hari             = "senin"
		besok
		sekarang
	)
	const satu1, dua1, tiga1 = 1, 2, 3
	fmt.Println(square, circle, triangle, numeric, floatNum, hari, besok, sekarang)
	fmt.Println(satu, dua, tiga)
	fmt.Println("")

	// Operator Aritmetika
	/*
		Tanda	Penjelasan
		+	penjumlahan
		-	pengurangan
		*	perkalian
		/	pembagian
		%	modulus / sisa hasil pembagian
	*/
	// misal
	value := (((2+6)%3)*4 - 2) / 3
	fmt.Println(value)
	// Operator Perbandingan
	/*
		Tanda	Penjelasan
		==	apakah nilai kiri sama dengan nilai kanan
		!=	apakah nilai kiri tidak sama dengan nilai kanan
		<	apakah nilai kiri lebih kecil daripada nilai kanan
		<=	apakah nilai kiri lebih kecil atau sama dengan nilai kanan
		>	apakah nilai kiri lebih besar dari nilai kanan
		>=	apakah nilai kiri lebih besar atau sama dengan nilai kanan
	*/
	// misal
	value1 := (((2+6)%3)*4 - 2) / 3
	isEqual := value1 == 2
	fmt.Printf("nilai %d (%t) \n", value1, isEqual)

	// Operator Logika
	/*
		Tanda	Penjelasan
			&&	kiri dan kanan
			||	kiri atau kanan
			!	negasi / nilai kebalikan
	*/
	// misal
	left := false
	right := true

	leftAndRight := left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)
	leftOrRight := left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)
	leftReverse := !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
	rightReverse := !right
	fmt.Printf("!right \t\t(%t) \n", rightReverse)

	fmt.Println("")
	// seleksi kondisi digunakan untuk mengontrol alur eksekusi flow program
	// analogi seperti fungi rambu lalu lintas, jika kondisi tertentu terpenuhi maka jalur akan berubah
	// acuan dari seleksi kondisi adalah nilai boolean bisa berasal dari variable ataupun hasil operasi perbandingan
	// nilai atau value menentukan blok kode mana yang di eksekusi

	// GO memiliki 2 maca keyword yaitu if else dan switch case

	fmt.Println("")
	// seleksi kondisi menggunakan if, else if dan else
	// penggunaan ini dalam golang sama seperti bahasa lainnya yang membedakan
	// hanya pada kurungnya (parengthesis) yang tidak digunakan misal

	point := 99

	if point == 100 {
		fmt.Println("mendapatkan nilai sempurna")
	} else if point > 70 {
		fmt.Println("lulus")
	} else if point > 50 {
		fmt.Println("remidi")
	} else {
		fmt.Println("tidak lulus")
	}

	// temporary variable pada if else adalah variable yang bisa digunakan pada deretan blok seleksi kondisi dimana
	// variable ditempatkan, berikut manfaat variable temporary if else
	// scope atau cakupan variable jelas, hanya bisa digunakan pada blok seleksi kondisi itu saja
	// kode menjadi lebih rapih
	// ketika nilai variable tersebut didapat dari sebuah koputasi, perhitungan tidak perlu dilakukan didalam blok masing masing kondisi

	// contoh
	point1 := 231098231.2312

	if percent := point1 / 100; percent >= 100 {
		fmt.Printf("nilai %.2f (%.2f%%) sempurna\n", point1, percent)
	} else if percent >= 70 {
		fmt.Printf("nilai %.2f (%.2f%%) lulus\n", point1, percent)
	} else {
		fmt.Printf("nilai %.2f (%.2f%%) tidak lulus\n", point1, percent)
	}

	point2 := 10000.21231
	if nilainya := point2 / 2; nilainya >= 100 {
		fmt.Printf("nilai %.2f (%.2f%%) ganjil\n", point2, nilainya)
	} else if nilainya >= 70 {
		fmt.Printf("nilai %.2f (%.2f%%) genap\n", point2, nilainya)
	} else if nilainya >= 50 {
		fmt.Printf("nilai %.2f (%.2f%%) prima\n", point2, nilainya)
	} else {
		fmt.Printf("nilai %.2f (%.2f%%) bukan prima\n", point2, nilainya)

	}

	fmt.Println("")
	// seleksi kondisi menggunakan switch case
	// switch merupakan seleksi kondisi yang sifatnya fokus pada satu variable
	// kemudian di cek nilainya contoh sederhanaya seperti ini
	point = 90

	switch point {
	case 100:
		fmt.Println("mendapatkan nilai sempurna")
	case 90:
		fmt.Println("nilai sangat bagus")
	case 80:
		fmt.Println("nilai bagus")
	case 70:
		fmt.Println("nilai cukup")
	default:
		fmt.Println("nilai tidak terdefinisi")
	}

	// pemanfaatan case untuk banyak kondisi
	// caranya dengan menuliskan nilai permbanding-pembanding variable yang di switch setelah keyword case
	// dengan (,) misal
	fmt.Println("")
	point = 19

	switch point {
	case 20:
		fmt.Println("Perfect")
	case 19, 18, 17, 16, 15:
		fmt.Println("Good")
	case 14, 13, 12, 11, 10:
		fmt.Println("Not Bad")
	default:
		fmt.Println("Bad")
	}

	// Kurung Kurawal pada switch case and default
	// tanda kurung kurawal ({}) pada blok kode case dan default bersifat opsional
	// bagus jika dipakai pada blok kondisi yang didalamnya ada banyak statement, agar terlihat rapih
	// nisal

	point = 50

	switch point {
	case 100:
		fmt.Println("GG bro")
	case 90, 80, 70:
		fmt.Println("not bad brother")
	default:
		{
			fmt.Println("bad")
			fmt.Println("try again")
		}

	}

	// switch case dengan gaya if else
	/*
		uniknya GO, switch bisa digunakan dengan gaya ala if-else nilai
		yang dibandingkan tidak tuliskan setelah ketword switch, melainkan akan ditulis langusng
		dalam bentuk perbandingan dalam keyword case

		contoh kode di bawah ini, kode program switch case
		diubah ke dalam gaya if else, Variable point dihilangkan dari keyword switch
		lalu kondisinya dituliskan setiap case
	*/

	point = 20

	switch {
	case point == 30:
		fmt.Println("Perfect")
	case (point < 19) && (point > 10):
		fmt.Println("Good")
	default:
		{
			fmt.Println("Bad")
			fmt.Println("Try again")
			fmt.Println("Nice try DIdiy")
		}
	}

	// switch case dengan fallthrough
	// uniquenya Go dimana jika case terpenuhi pengecekan kondisi tidak di teruskan ke case selanjutnya
	// keyword fallthrough digunakan untuk memaksa pengecekan kondisi ke case selanjutnya
	// tanpa menghiraukan nilai kondisinya, efeknya adalah case di pengecekan selanjutnya selalu dianggap true( meskipun aslinya bisa saja false akantetap dianggap true)
	// misal
	fmt.Println("")

	point = 6

	switch {
	case point == 8:
		fmt.Println("Perfect")
	case (point < 8) && (point > 2):
		fmt.Println("Good")
		fallthrough
	case point < 5:
		fmt.Println("Not Bad")
	default:
		{
			fmt.Println("Bad")
			fmt.Println("Try again")
			fmt.Println("Nice try DIdiy")
		}
	}
	/*
			seharusnya setelah pengecekan case (point < 8) && (point > 2): selesai, dilanjut ke pengecekan
			case point < 5, karena adanya fallthrough, maka pengecekan langsung dilanjutkan ke case point < 5
			akan selalu dianggap true meskipun sebenarnya false

		pada case dalam switch diperbolehkan terdapat lebih dari satu fallthrough
	*/
	fmt.Println("")

	// Seleksi kondisi bersarang
	/*
		nested condition adalah seleksi kondisi yang berada dalam seleksi kondisi, yang mungkin juga berada dalam seleksi kondisi dan seterusnya seleksi kondisi bersarang bisa dilakukan pada if - else maupun switch case
		atau pun kombinasi keduanya
	*/
	// misal

	point = 10

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("Perfect")
		default:
			{
				fmt.Println("Nice")
			}
		}
	} else {
		if point == 5 {
			fmt.Println("Good")
		} else if point == 3 {
			fmt.Println("Not Bad")
		} else {
			fmt.Println("Bad")
		}
	}
	fmt.Println("")

	// tugas Khan Academhy dari seleksi kondisi bersarang
	score_kedisiplinan := 90
	score_kerjasamaTim := 80
	score_productivity := 85

	switch {
	case score_kedisiplinan >= 85 && score_productivity >= 85 && score_kerjasamaTim >= 85:
		fmt.Println("Sangat baik")
	case score_kedisiplinan >= 80 && (score_productivity >= 70 || score_kerjasamaTim >= 70):
		fmt.Println("Baik")
	case (score_kedisiplinan >= 60 && score_kerjasamaTim >= 60) && score_productivity >= 50:
		fmt.Println("Cukup")
	default:
		fmt.Println("Perlu Perbaikan")
	}

	// perulangan atau looping
	/*
		 perulangan adalah proses mengulang-ulang eksekusi blok kode tanpa henti, selama kondisi
		kondisi yang dijadikan acuan terpenuhi, biasanya disiapkan variable untuk iterasu atau
		variable penanda kapan perulanagn harus berhenti

		di GO keyword perulangan hanya for saja,
		meski hanya for kemampuannya merupakan gabungan for, foreach dan while
	*/

	// Perulanngan dengan keyword "FOR"
	// misal
	for i := 0; i < 5; i++ {
		fmt.Println("Perulangan ke ", i)
	}

	fmt.Println("")

	//Penggunaan Keyword for dengan argumen hanya kondisi
	// dengan menuliskan kondisi setelah keyword for deklarasikan dan iterasi variable counter
	// tidak dituliskan setekah keyword for, hanya kondisi perulangan saja , konsepnya mirip seperti while
	// misal contoh for dengan argumen hanya kondisi seperti if output yang dihasilkan sama seperti penerapan for

	i := 0

	for i < 5 {
		fmt.Println("Perulangan ke ", i)
		i++
	}

	fmt.Println("")

	// Penggunaan Keyword for tanpa argumen
	/*
		cara ketiga for dituliskan tanpa kondisi dengan ini perulangan tanpa henti sama dengan for true
		berhenti dengan menggunakan keyword break
	*/

	i = 0

	for {
		fmt.Println("Perulangan ke ", i)
		i++

		if i == 5 {

			break
		}
	}

	fmt.Println("")

	// penggunaan keyword for - range
	/*
		care ke 4 perulangan dengan kombinasi keyword for dan range
		digunakan untuk me-loopng data gabungan(misalnya string, array, slice dan map)

	*/

	// misal
	fmt.Println("Perulangan dengan keyword for - range")
	var xs = "ayub subagiya" // string
	for i, v := range xs {
		fmt.Printf("index %d karakter %c\n", i, v)
	}

	var ys = [5]string{"ayub", "subagiya", "sumira", "subagiya", "budi"} // array
	for _, v := range ys {
		fmt.Printf("nilai %s\n", v)
	}

	var zs = ys[0:2] // slice
	for _, v := range zs {
		fmt.Printf("nilai %s\n", v)
	}

	var kvs = map[byte]int{'a': 1, 'b': 2, 'c': 3} // map
	for k, v := range kvs {
		fmt.Printf("key %c value %d\n", k, v)
	}

	for range kvs {
		fmt.Println("done")
	}

	for i := range 5 {
		fmt.Println(i)
	}

	// Penggunaan Keyword break dan continue
	/*
			break digunakan untuk menghentikan secara paksa sebuah perulangan,
		sedangkan continue digunakan untuk menghentikan iterasi saat ini dan melanjutkan ke iterasi berikutnya
	*/
	// misal
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Println("Perulangan ke ", i)
	}
	/*
		Kode di atas akan lebih mudah dicerna jika dijelaskan secara berurutan. Berikut adalah penjelasannya.

		Dilakukan perulangan mulai angka 1 hingga 10 dengan i sebagai variabel iterasi.
		Ketika i adalah ganjil (dapat diketahui dari i % 2, jika hasilnya 1, berarti ganjil), maka akan dipaksa lanjut ke perulangan berikutnya.
		Ketika i lebih besar dari 8, maka perulangan akan berhenti.
		Nilai i ditampilkan.
	*/

	// Perulangan bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, "")
		}
		fmt.Println()
	}
	/*
		Pada kode di atas, untuk pertama kalinya fungsi fmt.Println() dipanggil tanpa disisipkan parameter. Cara seperti ini bisa digunakan untuk menampilkan baris baru. Kegunaannya sama seperti output dari statement fmt.Print("\n").
	*/

	// pemanfaatan label dalam perulangan
	/*
		di perulangan bersarang break dan continue akan berlaku pada blok perulangan dimana ia digunakan saja
		cara kedua keyword ini bisa tertuju pada perulangan terluar atau perulangan tertentu yaitu dengan memanfaatkan teknik pemberian label
	*/

selesaiLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break selesaiLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

	// array merupakan kumpulan data bertipe sama yang disimpan dalam satu variabel
	/*
		array memiliki kapasitas yang nilainya ditentukan pada saaat pembuatan
		sehingga array tidak bisa menambah atau mengurangi elemen
	*/

	var names [5]string
	names[0] = "ayubs"
	names[1] = "subagiya"
	names[2] = "sumira"
	names[3] = "subagiyaa"
	names[4] = "zxiell"

	fmt.Println(names[2])

	// inisialisasi nilai awal dari array
	/*
		pengisian elemen array bisa dilakukan pada saat deklarasi variable
	*/
	// misal
	var fruits = [4]int{12, 232, 2112, 23121}
	fmt.Println("jumlah elemen array\t \t", len(fruits))
	fmt.Println("isi semua elemen array\t", fruits)

	// inisialisasi nilai array dengan gaya vertikal
	/*
		elemen array bisa dituliskan dalam bentuk horizontal maupun vertikal
	*/
	// contoh

	var fruits1 [4]string

	fruits1 = [4]string{"apple", "grape", "banana", "melon"}

	// vertikal
	fruits1 = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}
	fmt.Println(fruits1)
	fmt.Println("")

	// inisialisasi nilai array tanpa jumlah elemen
	/*
		deklarasi array yang nilainua tidak dituliskan jumlah lebar arraynya cukup
		ganti dengan (...) metode ini membuat kapasitas array otomatis dihitung dari jumlah array

	*/

	var numberss = [...]int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numberss)
	fmt.Println("jumlah elemen array\t", len(numberss))

	// array multidimensi
	/*
		merupalan array yang tiap elementnya juga merupakan array
		level kedalaman array ini tidak terbatas

		cara deklarasi dengan array biasa, bedanya adalah pada array biasa tiap elemen berisi satu nilai
		sedangkan pada array multidimensi tiap elemen berisi array lagi

		array subdimensi/elemen boleh atau tidak dituliskan jumlah datanya
	*/
	// contoh

	var numbers11 = [2][3]int{{3, 2, 4}, {3, 1, 2}}
	var numbers22 = [2][3]int{{3, 2, 4}, {3, 1, 2}}

	fmt.Println(numbers11)
	fmt.Println(numbers22)

	// perulangan elemen array dengan keyword for

	var misal = [10]string{"ayub", "subagiya", "sumira", "subagiya", "budi", "ayub", "subagiya", "sumira", "subagiya", "budi"}
	for i = 0; i < len(misal); i++ {
		fmt.Printf("elemen %d %s \n", i, misal[i])
	}

	// perulangan elemen array dengan keyword for - range

	var misal11 = [5]string{"ayub", "subagiya", "sumira", "subagiya", "budi"}

	for i, misal12 := range misal11 {
		fmt.Printf("elemen %d %s \n", i, misal12)
	}

	var points = [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	for i, point1 := range points {
		fmt.Printf("elemen %d %d \n", i, point1)
	}

	// variable _ dalam for - range
	/*
		dalam loopong menggunakan for range ada ke butuhan untuk elemennya saja sedangkan indeks nya tidak

	*/

	var buah = [4]string{"apple", "Mangga", "jeruk", "pisang"}
	for i, buah := range buah {
		fmt.Printf("elemen %d %s \n", i, buah)
	}

	var fruitss = [4]string{"apple", "grape", "banana", "melon"}

	for _, fruit := range fruitss {
		fmt.Printf("nama buah : %s\n", fruit)
	}
	// variable _ digunakan untuk menampung nilai yang tidak digunakan
	/*
		( _ ) sering disebut variable pengangguran menampung nilai yang tidak digunakan

	*/

	var kipas = [3]string{"Baling-baling", "Kipas Angin", "Kipas Udara"}
	for _, kipas := range kipas {
		fmt.Printf("nama kipas : %s\n", kipas)
	}

	// alokasi elemen array menggunakan keyword make
	/*
		deklarasi dan sekaligus alokasi kapasitas array lewat keyword make
	*/

	var keyboard = [3]string{"Button", "Keyboard", "Mouse"}
	for _, keyboard := range keyboard {
		fmt.Printf("nama keyboard : %s\n", keyboard)
	}

	var buahbuahan = make([]string, 2)
	buahbuahan[0] = "apple"
	buahbuahan[1] = "mangga"
	fmt.Println(buahbuahan)

	// latihan membuat dan menampilkan array
	var tugas = [5]int{2, 4, 6, 8, 10}
	for i, tugas := range tugas {
		fmt.Println("Elemen ke-", i, ":", tugas)
	}
	// menjumlahkan dua array
	var array1 = [3]int{1, 2, 3}
	var array2 = [3]int{4, 5, 6}
	var arraytambahan = [3]int{}

	for i := 0; i < len(array1); i++ {
		arraytambahan[i] = array1[i] + array2[i]
	}
	fmt.Println("Hasil penjumlaham array :", arraytambahan)

	var cari = [4]int{3, 7, 9, 11}
	var target = 7
	var ditemukan = false
	for _, cari := range cari {
		if cari == target {
			fmt.Println("Angka Ditemukan")
			ditemukan = true
			break
		}

	}
	if !ditemukan {
		fmt.Println("Angka tidak ditemukan")
	}

	// panjang array
	var panjang1 = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Jumlah elemen array:", len(panjang1))

	// slice
	/*
		merupakan reference elemen array, slice dapat di buat, manipulasi dari array
		karena slice merupakan data reference, menjadikan perubahan data di tiap elemen
		slice akan berdampak pada slice lain yang memiliki referensi yang sama dan alamat
		memori yang sama
	*/

	// inisialisasi slice
	/*
		cara membuat slice mirip seperti array bedanya tidap perlu
		mendefinisikan jumlah elemen ketika di awal deklarasi, pengaksesan nya juga beda
	*/
	// misal
	var lagu = []string{"Houdini", "Eminem", "Lose Yourself", "Eminem", "Rap God"}
	fmt.Println(lagu[2])

	// berikut perbedaanya slice dan array
	//var laguA = []string{"Houdini", "Eminem", "Lose Yourself", "Eminem", "Rap God"} // slice
	//var laguB = [5]string{"Houdini", "Eminem", "Lose Yourself", "Eminem", "Rap God"} // array
	//var laguC = [...]string{"Houdini", "Eminem", "Lose Yourself", "Eminem", "Rap God"} // array

	/*
		Hubungan antara slice dengan array dan operasi slice
		Perbedaany hanya di penentuan alokasi pada saat inisialisasi.
		sebenarnya slice dan array tidak bisa dibedakan karena merupakan sebuah kesatuan, array
		adalah kumpulan nilai atau elemen, sedangkan slice adalah refrensi tiap elemen

		slice bisa dibentuk dari array yang sudah didefinisikan,
		caranya dengan teknik 2 index mengambil elemen-nya
	*/
	// misal

	var laguD = []string{"Houdini", "Eminem", "Lose Yourself", "Eminem", "Rap God"}
	var newLaguD = laguD[3:5]
	/*
				kode laguD[3:5} adalah pengaksesan elemen dalam slice laguD,
				yang dimulai dari index ke-3 sampai sebelum index ke-5, elemen
				yang memnuhi kriteria akan didapat kemudian disimpan pada variable
				lain sebagai slice baru

				contohya newLaguD adalah slice baru yang tercetak dari slice laguD
				dengan isi elemen "Eminem" dan "Rap God"

		ketika mengakses elemen array menggunakan satu baris indeks (misal laguD[3]),
		maka nilai yang didapat adalah hasil copy dari refrerensi aslinya, berbeda pengaksesan
		menggunakan 2 index seperti data[3:5] yang menghasilkan slice baru mendapat nilai dari
		referensi aslinya atau slice
	*/
	fmt.Println(newLaguD)
	/*
		misal
		Kode		Output							Penjelasan
		fruits[0:2]	[apple, grape]					semua elemen mulai indeks ke-0, hingga sebelum indeks ke-2
		fruits[0:4]	[apple, grape, banana, melon]	semua elemen mulai indeks ke-0, hingga sebelum indeks ke-4
		fruits[0:0]	[]								menghasilkan slice kosong, karena tidak ada elemen sebelum indeks ke-0
		fruits[4:4]	[]								menghasilkan slice kosong, karena tidak ada elemen yang dimulai dari indeks ke-4
		fruits[4:0]	[]								error, pada penulisan fruits[a:b] nilai a harus lebih kecil atau sama dengan b
		fruits[:]	[apple, grape, banana, melon]	semua elemen
		fruits[2:]	[banana, melon]					semua elemen mulai indeks ke-2
		fruits[:2]	[apple, grape]					semua elemen hingga sebelum indeks ke-2
	*/
	// slice merupakan tipe data reference
	/*
			artinya jika ada slice baru yang terbentuk dari slice lama, maka
		slice yang baru  itu akan memiliki referensi dan alamat memori yang sama dengan elemen slice lama
		setiap perubahan yang terjadi di elemen slice baru akan berdampak pada elemen slice lama

	*/

	// misal

	var starboy = []string{"The Weeknd", "Starboy", "The Hills", "Can't Feel My Face", "I Feel It Coming"}

	var aStarboy = starboy[2:5]
	var bStarboy = starboy[0:2]

	var aaStarboy = aStarboy[0:1]
	var bbStarboy = bStarboy[0:3]

	fmt.Println(starboy)
	fmt.Println(aStarboy)
	fmt.Println(bStarboy)
	fmt.Println(aaStarboy)
	fmt.Println(bbStarboy)

	bbStarboy[0] = "Ayub SUbagiya"

	fmt.Println(starboy)
	fmt.Println(aStarboy)
	fmt.Println(bStarboy)
	fmt.Println(aaStarboy)
	fmt.Println(bbStarboy)
	/*
		var fruits = []string{"apple", "grape", "banana", "melon"}

		var aFruits = fruits[0:3]
		var bFruits = fruits[1:4]

		var aaFruits = aFruits[1:2]
		var baFruits = bFruits[0:1]

		fmt.Println(fruits)   // [apple grape banana melon]
		fmt.Println(aFruits)  // [apple grape banana]
		fmt.Println(bFruits)  // [grape banana melon]
		fmt.Println(aaFruits) // [grape]
		fmt.Println(baFruits) // [grape]

		// Buah "grape" diubah menjadi "pinnaple"
		baFruits[0] = "pinnaple"

		fmt.Println(fruits)   // [apple pinnaple banana melon]
		fmt.Println(aFruits)  // [apple pinnaple banana]
		fmt.Println(bFruits)  // [pinnaple banana melon]
		fmt.Println(aaFruits) // [pinnaple]
		fmt.Println(baFruits) // [pinnaple]

	*/
	/*
		dipahami bahwa setelah slice yang isinya data adalah grape diubah menjadi pinnaple
		maka slice yang lain yang memiliki referensi dan alamat memori yang sama akan ikut berubah
		variable aFruits, bFruits, aaFruits, dan baFruits akan berubah sesuai dengan perubahan yang terjadi pada slice yang memiliki referensi dan alamat memori yang sama

		selanjutnya nilai baFruits[0] diubah menjadi "pinnaple" maka slice yang lain yang memiliki referensi dan alamat memori yang sama akan ikut berubah

	*/

	// Fungsi len()
	/*
		fungsi len() digunakan untuk menghitung jumlah elemen slice yang ada, sebagai contoh jika sebuah variable adalah slice dengan data 4 buah maka
		fungski ini pada variable tersebut akan mengenbalikan nilai 4
	*/

	var stan = []string{"Stan", "Eminem", "Lose Yourself", "Eminem", "Rap God"}
	fmt.Println(len(stan))

	// fungsi cap()
	/*
		fungsi cap() digunakan untuk menghitung lebar atau kapasitas maksimum slice
		nilai kembalian fungsi ini untuk slice yang baru dibuat pasti sama dengan len
		tapi bisa berubah seiring operasi slice dilakukan
	*/
	// misal
	var song = []string{"Stan", "Lose Yourself", "Rap God", "Till i collapse", "guilty conscience"}
	fmt.Println("ini jumlah awal dengan fungsi len", len(song))
	fmt.Println("ini jumlah awala dengan fungsi cap", cap(song))

	var aSong = song[0:3]
	fmt.Println(len(aSong))
	fmt.Println(cap(aSong))

	var bSong = song[0:3]
	fmt.Println(len(bSong))
	fmt.Println(cap(bSong))
	/*
			contoh lain
		var fruits = []string{"apple", "grape", "banana", "melon"}
		fmt.Println(len(fruits))  // len: 4
		fmt.Println(cap(fruits))  // cap: 4

		var aFruits = fruits[0:3]
		fmt.Println(len(aFruits)) // len: 3
		fmt.Println(cap(aFruits)) // cap: 4

		var bFruits = fruits[1:4]
		fmt.Println(len(bFruits)) // len: 3
		fmt.Println(cap(bFruits)) // cap: 3
	*/

	// penjelasan
	/*
				variable fruits disiapkan di awal dengan jumlah elemen 4, fungsi len(fruits) dan cap(fruits) akan menghasilkan nilai 4

				variable aFruits dan bFruits merupakan slice baru berisikan 3 buah elemen dari slice fruits,
				variable aFruits mengambil index 0, 1, 2 sedangkan bFruits mengambil index 1, 2, 3
				Fungsi len() menghasilkan angka 3, karena jumlah elemen kedua slice ini adalah 3. Tetapi cap(aFruits) menghasilkan angka yang berbeda, yaitu 4 untuk aFruits dan 3 untuk bFruits. kenapa? jawabannya bisa dilihat pada tabel berikut.

			Kode			Output					len()	cap()
			fruits[0:4]		[buah buah buah buah]		4	4
			aFruits[0:3]	[buah buah buah ----]		3	4
			bFruits[1:4]	---- [buah buah buah]		3	3

		slicing 2 index di analogikan menggunakan x dan y
		maka akan menjadi fruits[x:y]
		slicing yang dimulai dari indeks 0 hingga y akan mengembalikan elemen mulai dari indeks 0 sebelum indeks y dengan lebar kapasitas adalah sama dengan slice aslinya
		sedangkan slicing yang dimulai dari ideks x, dimana x adalah lebih dari 0 ini menjadi elemen mulai dari indeks membuat elemen ke x slice yang diambil menjadi elemen ke 0 slice baru
		ini yang membuat kapasitas slice berubah
	*/

	// Fungsi Append()
	/*
		digunakan untuk menambah elemen pada slice, elemen baru diposisikan setelah indeks paling akhir, nilai baik fungsi ini adalah slice yang sudah di tambahakan nilai barunya
	*/
	// misal
	var music = []string{"Stan", "Lose Yourself", "Rap God", "Till i collapse", "guilty conscience"}
	var cMusic = append(music, "Shingeki no Kyojin")

	fmt.Println(music)
	fmt.Println(cMusic)

	// hal yang perlu diperhatikan
	/*
		1. ketika jumlah elemen dan lebar kapasitas adalah sama (len(fruits) == cap(fruits)), maka penambahan elemen baru hasil append() akan membuat slice baru atau refrerensi baru
		2. jika jumlah elemen lebih kecil dibanding kapasitas (len(fruits) < cap(fruits)), maka elemen baru akan ditambahkan pada slice lama
	*/
	// misal

	var favorite = []string{"Stan", "Lose Yourself", "Rap God", "Till i collapse", "guilty conscience"}
	var aFavorite = favorite[0:3]

	fmt.Println(len(aFavorite)) // len: 3
	fmt.Println(cap(aFavorite)) // cap: 5

	fmt.Println(favorite)  // [Stan Lose Yourself Rap God Till i collapse guilty conscience]
	fmt.Println(aFavorite) // [Stan Lose Yourself Rap God]

	var cFavorite = append(aFavorite, "Shingeki no Kyojin")
	fmt.Println(favorite)  // [Stan Lose Yourself Rap God Shingeki no Kyojin guilty conscience]
	fmt.Println(aFavorite) // [Stan Lose Yourself Rap God]
	fmt.Println(cFavorite) // [Stan Lose Yourself Rap God Shingeki no Kyojin]
	/*
		contoh lain
		var fruits = []string{"apple", "grape", "banana"}
		var bFruits = fruits[0:2]

		fmt.Println(cap(bFruits)) // 3
		fmt.Println(len(bFruits)) // 2

		fmt.Println(fruits)  // ["apple", "grape", "banana"]
		fmt.Println(bFruits) // ["apple", "grape"]

		var cFruits = append(bFruits, "papaya")

		fmt.Println(fruits)  // ["apple", "grape", "papaya"]
		fmt.Println(bFruits) // ["apple", "grape"]
		fmt.Println(cFruits) // ["apple", "grape", "papaya"]

		Pada contoh di atas bisa dilihat, elemen indeks ke-2 slice fruits
		nilainya berubah setelah ada penggunaan keyword append() pada bFruits.
		Slice bFruits kapasitasnya adalah 3 sedang jumlah datanya hanya
		2. Karena len(bFruits) < cap(bFruits), maka elemen baru yang dihasilkan,
		terdeteksi sebagai perubahan nilai pada referensi yang lama
		(referensi elemen indeks ke-2 slice fruits), membuat elemen yang
		referensinya sama, nilainya berubah.
	*/

	// Fungsi copy()
	/*
		fungsi copy() digunakan untuk menyalin elements slice pada src(slice asal) ke dst(slice tujuan)
		copy(dst, src)

		jumlah element yang dicopy dari src adalah sejumlah lebar slice dst atau len(dst) jika jumlah slice pada src lebih kecil dari dst maka akan tercopy sejumlah elemen yang ada pada src semua

	*/
	// misal

	dst := make([]string, 3)
	src := []string{"apple", "grape", "banana", "melon"}
	n := copy(dst, src)

	fmt.Println(dst) // [apple grape banana]
	fmt.Println(src) // [apple grape banana melon]
	fmt.Println(n)   // 3

	//penjelasan
	/*
			pada kode variabel slide dst dipersiapkan dengan lebar 3 elemen,
			slice src yang isinya 4 element di copy ke dst menjadikan isi slice dst sekarang adalah 3 buah elements yang sama dengan 3 buah element src
			hasil dari operasi copy() adalah jumlah element yang berhasil di copy dari src ke dst

			yang tercopy hanya 3 buah (meski src memiliki 4 elements( hal ini karena copy hanya menyalin elements sebanyak len dari dst len(dst)

		berikut contoh kedua dst merupakan slice yang sudah ada isinya 3 buah elements variabale src yang juga merupakan slice dengan isi dua elements dicopy ke dst karena operasi copy() akan
		menyalin sejumlah len(dst) maka semua elements src akan tercopy karena jumlahnya di bawah atau sama dengan lebar dst
	*/
	// misal

	dst1 := []string{"apple", "grape", "banana"}
	src1 := []string{"watermelon", "pinnaple"}
	n1 := copy(dst1, src1)

	fmt.Println(dst1) // [watermelon pinnaple banana]
	fmt.Println(src1) // [watermelon pinnaple]
	fmt.Println(n1)   // 2
	/*
		Jika dilihat pada kode di atas, isi dst masih tetap 3 elements, tapi dua elements pertama adalah sama dengan src. Element terakhir dst isinya tidak berubah, tetap potato, hal ini karena proses copy hanya memutasi element ke-1 dan ke-2 milik dst, karena memang pada src hanya dua itu elements-nya.
	*/

	// pengaksesan elemen slice dengan 3 indeks
	/*
		3 index adalah teknik slicing untuk pengaksesan elemen yang sekaligus menentukan kapasitas cara penggunaanya yaitu dengan menyisipkan angka kapasitas di belakang seperti fruits[0:2:3] angka kapasitas yang diisilan tidak boleh melebihi dari kapasitas slice yang di slicing

	*/
	// misal
	var marshal = []string{"Stan", "Lose Yourself", "Rap God", "Till i collapse", "guilty conscience"}
	var aMarshal = marshal[0:2:3]
	var bMarshal = marshal[0:4]

	fmt.Println(marshal)
	fmt.Println(len(marshal))
	fmt.Println(cap(marshal))
	fmt.Println("")
	fmt.Println(aMarshal)
	fmt.Println(len(aMarshal))
	fmt.Println(cap(aMarshal))
	fmt.Println("")
	fmt.Println(bMarshal)
	fmt.Println(len(bMarshal))
	fmt.Println(cap(bMarshal))

	var t bool
	t = true
	f := !t
	or := f || t
	and := f && or
	fmt.Println(and)

	// Map
	/*
		map adalah tipe data asosiatif dari Go yang terbentuk key-value pair,
		Data/Value yang disimpan di map selalu di sertai dengan key, keynya juga harus unik,
		karena digunakan sebagai penanda atau identifier untuk pengaksesan value yang disimpan di map

		kalau dilihat map seperti slice namun identifier yang digunakan untuk pengaksesan bukanlah index numerik
		melainkan bisa dalam tipe data apapun sesuai yang di ingin kan
	*/
	// penggunaan map, pengaplikasiannya mudah dengan menuliskan keyword map diikuti tipe data key dan valuenya

	// misal
	var phone = map[string]int{}
	phone = map[string]int{}

	phone["Samsung"] = 10000000
	phone["Iphone"] = 20000000

	fmt.Println("harga Samsung", phone["Samsung"])
	fmt.Println("harga Iphone", phone["Iphone"])

	var song133 = map[string]string{}
	song133 = map[string]string{}

	song133["Stan"] = "Eminem"
	song133["Gudako"] = "Fate Grand Order"

	fmt.Println("ini lagu stan artist", song133["Stan"])
	fmt.Println("ini lagu gudako artist", song133["Gudako"])

	// contoh penggunaan map
	/*
		var chicken map[string]int
		chicken = map[string]int{}

		chicken["januari"] = 50
		chicken["februari"] = 40

		fmt.Println("januari", chicken["januari"]) // januari 50
		fmt.Println("mei",     chicken["mei"])     // mei 0

		Variabel chicken dideklarasikan bertipe data map, dengan key ditentukan tipenya adalah string dan tipe value-nya int. Dari kode tersebut bisa dilihat bagaimana cara penerapan keyword map untuk pembuatan variabel.

		Kode map[string]int merepresentasikan tipe data map dengan key bertipe string dan value bertipe int.

		Zero value atau nilai default variabel map adalah nil. Dari sini maka penting untuk menginisialisasi nilai awal map agar tidak nil. Jika dibiarkan nil, ketika map digunakan untuk menampung data pasti memunculkan error.

		Cara untuk inisialisasi map dengan menambahkan kurung kurawal buka tutup di akhir penulisan map, contoh: map[string]int{}.

		Cara menambahkan item pada map adalah dengan menuliskan variabel-nya, kemudian diikuti dengan key pada kurung siku variabel (mirip seperti cara pengaksesan elemen slice), lalu operator =, kemudian nilai/data yang ingin disimpan. Contohnya seperti chicken["februari"] = 40. Sedangkan cara mengakses item map dengan cukup dengan menuliskan nama variabel diikuti kurung siku dan key.

		Pengisian data pada map bersifat overwrite, artinya variabel sudah memiliki item dengan key yang sama, maka value item yang lama (dengan key sama) akan ditimpa dengan value baru.
	*/

	// inisialisasi nilai map
	/*
		zero value dari map adalah nil, disarankan menginisialisasi secara explisit nilai awalanuya agar tidak nil

	*/
	// misal
	var data map[string]int
	//data["one"] = 1 ini error

	data = map[string]int{}
	data["one"] = 1

	/*
		nilai variable bertipe map didefinisikan di awal, caranya dengan menambahkan kurung kurawal
		setelah tipe data, kemudian menuliskan key dan valuenya dalam kurawa tersebut

		cara ini mirip dengan definisi array/slice hanya saja dalam bentuk key-value

	*/
	// misal
	// ini horizontal
	var eminem = map[string]int{"stan": 20, "lose yourself": 30, "rap god": 40}
	fmt.Println(eminem)
	fmt.Println("eminem stan", eminem["stan"])
	fmt.Println("eminem lose yourself", eminem["lose yourself"])

	// ini vertikal
	var eminem1 = map[string]int{
		"stan":          20,
		"lose yourself": 30,
		"rap god":       40,
	}
	fmt.Println(eminem1)

	var harga = map[string]int{
		"Samsung": 10000000,
		"Iphone":  20000000,
		"Infinix": 3000000,
		"Oppo":    4000000,
		"Macbook": 20000000,
		"Tuf":     15000000,
	}
	fmt.Println("harga dari samsung", harga["Samsung"])
	fmt.Println("harga dari Iphone", harga["Iphone"])
	fmt.Println("harga dari Infinix", harga["Infinix"])
	fmt.Println("harga dari Oppo", harga["Oppo"])
	fmt.Println(harga)

	/*
		key dan value dituliskan dengan pembatas titik dua (:) sedangkan tiap itemnya dituliskan dengan pembatas tanda koma ( , )
		khusus deklarasi dengan gaya vertikal, tanda koma perlu di tuliskan setelah item akhir

		variable map dapat dinisialisasi dengan tanpa nilai awal, caranya dengan menggunakan tanda kurung kurawal
		contohnya map[string]int{}
		atau dengan keyword make dan new

	*/
	// misal
	var woman1 = map[string]int{}
	var woman2 = make(map[string]int)
	var woman3 = *new(map[string]int)
	fmt.Println(woman1)
	fmt.Println(woman2)
	fmt.Println(woman3)

	// iterasi atau oengulangan item map menggunakan for - range
	/*
		item variable map bisa di iterasi menggunakan for - range
		cara penerapannya masih sama seperti pada slice
		dengan perbedaan pada map data yang dikembalikan di tiap perulangannya adalah key dan value
		bukan indels dan elemen
	*/
	// misal
	var tws = map[string]int{
		"Baseus encok":  200000,
		"SoundCore 50i": 300000,
		"JBL 100":       400000,
	}
	for key, val := range tws {
		fmt.Println(key, " :\t", val)
	}

	// menghapus item map
	/*
		fungsi delete() digunakan untuk menghapus item map dengan key tertentu pada variable map
		dengan memasukkan object map dan key item yang ingin di hapus sebagai argument pemanggil fungsi
		delete()
	*/
	// misal
	var bottle = map[string]int{"tupperware": 20000, "locknlock": 30000, "zujirushi": 40000}
	fmt.Println(len(bottle))
	fmt.Println(bottle)

	delete(bottle, "tupperware")

	fmt.Println("")
	fmt.Println(len(bottle))
	fmt.Println(bottle)

	// deteksi keberadaan item dengan key tertentu

	/*
		untuk mengetahui apakah dalam variable map terdapat item dengan key tertentu
		atau tidak yaitu dengan memanfaatkan 2 variable sebagai penampung nilai kembalian pengaskesan item,
		return value ke-2 sifatnya opsional boleh ditulis boleh juga tidak
		isi nilainua bool jika berisi true menendalan bahwa item yang dicari ada di map jika false maka tidak add

	*/
	// misal

	var coffe = map[string]int{
		"arabica": 20000,
		"robusta": 30000,
		"luwak":   40000,
	}
	var value111, ituAda = coffe["starbucks"]

	if ituAda {
		fmt.Println(value111)
	} else {
		fmt.Println("item tidak ada")
	}

	// kombinasi slice dan map
	/*
		slice dan map bisa dikombinasikan, pada praktiknya cukup sering digunakan
		contohnya untuk keperluan penyimpan data array yang beriskan map

		cara penerapannya dengan []map[string]int{} tipe tersebut artinya sebuah slice yang bertipe
		slice yang tipe elemennya adalah map[string]int
	*/
	// misal

	var laguFav = []map[string]string{
		{"judul": "Stan", "penyanyi": "Eminem"},
		{"judul": "Lose Yourself", "penyanyi": "Eminem"},
		{"judul": "Rap God", "penyanyi": "Eminem"},
	}

	for _, laguFav := range laguFav {
		fmt.Println(laguFav["judul"], "oleh", laguFav["penyanyi"])
	}

	var bestBand = []map[string]string{
		{"band": "Linkin Park", "vocalist": "Chester Bennington"},
		{"band": "Queen", "vocalist": "Freddie Mercury"},
		{"band": "Cold Play", "vocalist": "Chris Martin"},
	}
	for _, bestBand := range bestBand {
		fmt.Println(bestBand["band"], "oleh", bestBand["vocalist"])
	}
	/*
		variable laguFav berisikan 3 item bertipe map[string]string, setiap item map berisikan 2 key yaitu judul dan penyanyi
		penulisan tipe data tiap item adalah opsional bisa ditulis bisa juga tidak, karena go bisa mengetahui tipe data dari item map yang diisikan
		dalam []map[string]string{} tiap elemen bisa saja memiliki key yang berbeda-beda, tidak harus sama
	*/

	var paper = []map[string]int{
		{"hvs": 10000, "art paper": 20000, "kertas koran": 30000},
		{"penjualan": 100000, "pembelian": 200000, "laba": 300000},
		{"pajak": 900000, "PPh": 70000, "perkara": 300000},
	}
	for _, paper := range paper {
		fmt.Println(paper["hvs"], "oleh", paper["art paper"])
		fmt.Println(paper["penjualan"], "oleh", paper["pembelian"])
	}

	// Fungsi
	/*
		 dalam konteks pemrograman fungsi adalah sekumpulan blok kode yang di bungkus denagn nama tertentu
		penerapan fungsi yang tepat akan menjadikan kode lebih modular dan juga dry (don't repeat yourself)
		yang artinya tidak perlu menuliskan bnyakkode untuk kegunaan yang sama
		cukup deklarasikan satu satu blok kode sebagai suatu fungsi lalu panggil sesuai kebutuhan

	*/

	// peneraoan fungsi
	/*
		pengimplementasian fungsi banyak pada praktek sebelumnya yaitu fungsi main()
		Fungsi main() sendiri merupakan fungsi utama pada program GO

		selain fungsi main() kita juga bisa membuat fungsi lainnya dan caranya cukup mudah yaitu degna menuliskan keyword
		func kemudian di ikuti nama fungsi lalu kurung () yang bisa diisi parameter
		dan di akhiri dengan kurung kurawal untuk membungkus blok kode

		Parameter merupakan variael yan menempel fungsi yang nilainya ditentukan saat pemanggilan fungsi tersebut
		parameter sifatnya opsional suatu fungsi bisa tidak memiliki parameter, atau bisa saja memiliki satu atau banyak parameter
		tergantung kebutuhan
	*/
	// misal fungsi

	beriSalam("ayub", "programmer", "backend")
	printMessage("halo", []string{"ayub", "subagiya", "sumira"})
	/*
		pada kode di atas dibuat sebuah fungsi baru dengan nama printMessage() memiliki 2 buah parameter string message dan slice string arr
		fungsi tersebut dipanggil dalam mainI() dalam pemanggilannya disisipkan dua buah argument paramewter
		fungsi tersebut dipanngil dalam main

		1. argument parameter pertama adalah string "helo" yang ditampung oleh parameter message
		2. Argument parameter kedua adalah slice string yang berisikan 3 buah string yang ditampung oleh parameter arr

		di dalam printMessage(), nilai arr yang merupakan slice string di gabungkan menjadi sebuah string dengan pembatas adalah karakter spasi
		dengan memanfaatkan fungsi strings.Join() ini harus di import

	*/
	// fungsi dengan return value / nilai balik
	/*
		selain parameter, fungsi juga bisa memiliki attribute return value atau nilai balik
		fungsi yang memiliki return value saat deklarasinya harus ditentukan terlebih dahulu tipe data dari nilai baliknya

		fungsi yang tidak mengembalikan nilai apapun contohnya seperti fungsi main()
		dan printMessage() bisa disebut void function
	*/
	// misal
	var random = randomWithRange(2, 9999999999999999)
	fmt.Println("nilai random", random)

	var randomValue int

	randomValue = randomWithRange(2, 9999999999999999)
	fmt.Println("nilai random", randomValue)

	randomValue = randomWithRange(10, 100)
	fmt.Println("nilai random", randomValue)

	/*
		jadi kalau untuk return value itu dengan disisipkan juga tipe data yang akan dikembalikan
	*/

	// contoh lagi
	var best string
	best = bestSong("Lose Yourself", "Eminem", "The Eminem Show")
	fmt.Println(best)

	// penggunaan fungsi ran.New()
	// contog fungsi untuk return value yang menghentikan
	divideNumber(10, 5)

	// fungsi multiple return
	/*
			dalam Go suatu fungsi bisa saja mengembalikan nilai balik lebih
		dari 1 buah, Teknik ini untuk alternatif selain menggunakan tipe
		data kolektif seperti map, slice atau struct
	*/
	// penerapan fungsi multiple return
	/*
		cara membuat fungsi agar memiliki banyak nilai balik tidaklah sulit
		caranya saat melakukan deklarasi fungsi semua tipe data nilai balik yang
		ingin dikembalikan kemudian dalam body fungsi, pada penggunaan
		keyword return diikuti dengan nilai yang ingin dikembalikan
	*/
	// misal
	fmt.Println(calculate(10))
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Println("luas lingkaran", area)
	fmt.Println("keliling lingkaran", circumference)

	fmt.Println("")
	var jari float64 = 99
	var area1, circumference1 = calculate1(jari)
	fmt.Println("luas lingkaran", area1)
	fmt.Println("keliling lingkaran", circumference1)

	result := song11("Eminem", "Stan", "Lose Yourself", "Rap God")
	fmt.Println(result)

	var avg = menghitung(1, 2, 3, 4, 5, 6, 7, 8, 9, 22222)
	fmt.Printf("rata-rata: %.2f\n", avg)

	// penjelasan tambahan
	/*
		pengunaan fungsi fmt.Sprintf() dasarnya sama seperti fmt.Printf()
		yang membedakan adalah fmt.Sprintf() mengembalikan nilai string yang sudah diformat
	*/
	/*
		fungsi float64() digunakan untuk mengkonversi tipe data numerik ke float64

	*/

	// pengisian parameter fungsi cariadic menggunakan data slice
	/*
		slice bisa digunakan sebagai argument pada fungsi variadic, caranya penerapannya
		tulis nama variable disertai tanda titik tiga (...) setelah nama variable yang dijadikan parameter
	*/

	var kangka = []int{2, 3, 4, 5, 6, 123, 231}
	var avg1 = menghitung(kangka...) // bisa dilihat pada kode di atas
	fmt.Printf("rata-rata: %.2f\n", avg1)

	/*
		pada kode diatas, variable numbers bertipe data slice int, disisipkan pada pemanggilan fungsi menghitung() dengan menambahkan tanda titik tiga (...) di belakang nama variable
		teknik ini sangat berguna pada case dimana sebuah data slice perlu untuk digunakan untuk argument parameter variadic
	*/
	var kingki = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var avg2 = menghitung(kingki...)
	fmt.Printf("rata-rata: %.2f\n", avg2)
	// pada deklarasi parameter fungsi variadic tanda 3 titik (...) di tuliskan sebelum tipe data parameter sedangkan pada pemanggilan fungsi dengan menyisipkan parameter array, tanda tersebut dituliskan di belakang variablenysa

	// fungsi dengan parameter Biasa dan Variadic
	/*
		para meter variadic bisa dikombinasikan dengan parameter biasa, dengan syarat parameter variadicnya harus diposisikan di akhir
	*/
	// misal
	yourHobbies("ayub", "coding", "reading", "swimming", "gaming")

	/*
		nilai parameter pertama dari fungsi yourHobbies() akan ditampung oleh name, sedangkan nilai
		parameter kedua dan seterusnya akan ditampung oleh hobbies sebagai slice

		jika parameter kedua dan seterusnya ingin diisi dengan data dan slice maka gunnakan tanda titik tiga (...) di belakang nama parameter\

	*/
	// misal
	var hobbies = []string{"coding", "reading", "swimming", "gaming"}
	yourHobbies("ayub", hobbies...)

	var song111 = []string{"Stan", "Lose Yourself", "Rap God", "Till i collapse", "guilty conscience"}
	yourHobbies("Eminem", song111...)

	// Fungsi Closure
	/*
		merupakan fungsi tanpa nama bisa disimpan dalam variable closure memiliki sifat seperti fungsi yang disimpan
		untuk pencarian nilai terendah dan tertinggi dari data array logika pencariannya dibungkus dalam closure yang ditampung oleh variable
	*/
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	var kecil, besar = getMinMax(numbers)
	fmt.Printf("data: %v\nmin : %v\nmax : %v\n", numbers, kecil, besar)

	//pointer
	/*
		Pointer adalah refrensi atau alamat memori, variable pointer berarti variable yang berisi alamat memori suatu nilai sebagai contoh sebuah variable bertipe integer
		memiliki nilai 4, maksud pointeer adalah alamat memori dimana nilai 4 disimpan, bukan nilai 4 itu sendiri

		variable-variable yang memiliki reference atau alamat pointer yang sama dan saling berhubungan satu sama lain dan nilainya pasti sama
		ketika ada perubahan nilai, maka akan memberikan efek kepada varibale lain (yang preferensi sama) yaitu nilainya ikut berubah

	*/
	// Penerapan Pointer, menggunakan tanda asterisk (*) sebelum tipe data variable
	// misal

	//var number22 *int
	//var name22 *string
	// nilai default pointer adalah nill (kosong), variable Pointer tidak bisa menampung nilai bukan pointer dan sebaliknya variable biasa tidak bisa menampung nilai pointer
	// DUA HAL PENTING POINTER
	/*
		1. Variable biasa bisa diambil nilai pointernya, caranya dengan menambahkan tanda ampersand (&) tepat sebelum nama variable, metode ini disebut dengan referencing
		2. Dan sebaliknya, nilai asli variable pointer juga bisa diambil, dengan cara menambahkan tanda asterisk (*) tepat sebelum nama variable, dapat disebut dereferencing

	*/

	// misal
	var nomorA int = 10
	var nomorB *int = &nomorA
	fmt.Println("number A (value) :", nomorA)
	fmt.Println("number A (address) :", &nomorA)
	fmt.Println("number B (value) :", *nomorB)
	fmt.Println("number B (address) :", nomorB)
	/*
		variable numberB dideklarasikan bertipe pointer int dengan nilai awal adalah referensi variable nomorA
		dengan ini variable nomorA dan nomorB menampung data dengan refrensi alamat memori yang sama

		varibale pointer jika di-print akan menghasilkan string alamat memori (dalam notasi Heksadesimal)
		contohnya seperti numberB yang diprint menghasilkan alamat memori dari variable nomorA

		nilai asli sebuah variable pointer bisa didapatkan dengan cara di-dereference terlebih dahulu
		misal pada kode *nomorB, artinya mengambil nilai asli dari variable nomorA yang ditampung oleh variable pointer nomorB

	*/

	var harga11 = 100000
	var afterDiscount *int = &harga11
	fmt.Println("harga asli", harga11)
	fmt.Println("harga setelah diskon", *afterDiscount)

	// EFEK PERUBAHAN NILAI POINTER
	/*
		ketika salah satu variable pointer diubah nilainya, sedang ada variable lain yang memiliki refrensi memori yang sama maka
		nilai variable lain tersebut juga akan berubah
	*/
	// misal
	var nomor1 = 10
	var nomor2 *int = &nomor1

	fmt.Println("nomor 1 (value) :", nomor1)
	fmt.Println("nomor 1 (address) :", &nomor1)

	fmt.Println("nomor 2 (value) :", *nomor2)
	fmt.Println("nomor 2 (address) :", nomor2)

	fmt.Println("")
	nomor1 = 20
	fmt.Println("nomor 1 (value) :", nomor1)
	fmt.Println("nomor 1 (address) :", &nomor1)
	fmt.Println("nomor 2 (value) :", *nomor2)
	fmt.Println("nomor 2 (address) :", nomor2)
	/*
		variable nomor1 dan nomor2 memiliki refrensi yang sama, perubahan pada salah satu nilai variable tersebut akan memberikan efek pada variable lainnya yang terikat pointer
		misal jika nomor1 nilainya di ubah maka membuat nilai nomor2 juga berubah
	*/

	fmt.Println(nomor1)

	change(&nomor1, 100)
	fmt.Println(nomor1)
	fmt.Println(&nomor2)

	/*
		fungsi change() memiliki 2 parameter. yaitu original yang tipenya adalah pointer int dan valuenya yang bertipe int
		didalam fungsi tersebut nilai asli parameter pointer original diubah
		fungsi change() kemudian diimplementasikan di main, Variable number yang nilai awalnya adalah 4 diambil referensinya lalu digunakan sebagai
		parameter pada pemanggilan fungsi change()

		nilai variable number berubah menjadi 10 karena perubahan yang terjadi pada fungsi change() yang mengubah nilai asli dari variable number
	*/

}
func change(original *int, value int) {
	*original = value
}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")
	fmt.Printf("Hello my name is : %s\n", name)
	fmt.Printf("My hobbies are : %s\n", hobbiesAsString)
}

//#########################################
//# function lainnya diatas adalah func main
//#########################################

func beriSalam(name, pekerjaan, divis string) {
	var salam = fmt.Sprintf("Halo %s, selamat datang sebagai %s di divisi %s", name, pekerjaan, divis)
	fmt.Println(salam)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

// contoh func dengan return value
var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func randomWithRange(min, max int) int {
	var value = randomizer.Int()%(max-min+1) + min
	return value
}

func bestSong(song, artist, album string) string {
	var best = fmt.Sprintf("lagu terbaik %s oleh %s dari album %s", song, artist, album)
	return best
}

// fungsi return untuk menghentikan proses eksekusi fungsi
func divideNumber(m, n int) {
	if n == 0 {
		fmt.Println("tidak bisa membagi dengan 0")
		return
	}
	var result = m / n
	fmt.Printf("hasil %d / %d = %d\n", m, n, result)
}

// contoh fungsi multiple return
func calculate(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference

}

/*
fungsi calculate() memiliki satu buah parameter yaitu d(diameter) didalam fungsi terdapat operasi perhitungan nilai luas dan keliling dari nilai d
kemudian hasilnya dijadikan sebagai return value

pendefinisiannya menggunakan banyak nilai balik bisa dilihat pada kode di atas, langsung tulis tipe data semua nilai balik dipisah tanda koma,
lalu ditambahkan kurung di antaranya
func calculate(d float64) (float64, float64)
pada bagian return harus dituliskan juga semmua data yang dijadikan nilai balik
return area, circumference
fungsi calculate memiliki banya nilai balik, maka dalam pemangilannya harus dissiapkan juga sejumlah variable untuk
menampung nilai balik fungsi (sesuai dengan jumlah nailai balik yang dideklarasikan)

var area, circumference = calculate(diameter)
*/

// Fungsi dengan Predefined Return Value
/*
keunikan lainnya Go varibale yang digunakan sebagai nilai balik bisa didefinisikan di awal

*/
func calculate1(d float64) (area, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d
	return
}

// fungsi calculate1() bisa di modifikasi menjadi lebih sederhana
/*
func calculate2(d float64) (area float64, circumference float64)
fungsi dideklarasikan dengan 2 buah tipe data, dan variable yang nantinya dijadikan nilai balik juga
dideklarasikan variable area yang bertipe float64 dan circumference yang bertipe float64

karena variable nilai balik suda ditentukan di awal, untuk menegembalikan nilai cukup dengan memanggil
return tanpa perlu diikuti variable apapun nilai terakhir area dan circumference sebelum oemanggil keyword return adalah hasil dari fungsi di atas
*/

// penjelasan tambahan
/*
fungsi math.Pow() digunakan untuk menghitung nilai pangkat, fungsi ini membutuhkan 2 buah parameter yaitu nilai yang akan dipangkatkan dan pangkatnya
fungsi math.Pi adalah konstanta yang merepresentasikan nilai phi dalam matematika
*/

// Fungsi Variadic
/*
Go mengadopsi konsep variadic function dengan parameter yang menampung nilai sejenis yang tidak terbatas jumlahnya

paramerer variadic memliki sifat yang mirip slice yaitu nilai dari parameter2x yang disisipkan bertipe data sama
dan kesemuanya cukup ditampung oleh satu variable saja

pengaksesannya juga menggunakan index seperti slice

*/

// contoh fungsi variadic
/*
deklrasikan parameter variadic dengan (...) di belakang tipe data parameter
sebelum tipe data nantinya nilai yang disisipkan sebagai parameter akan ditampung oleh variable

*/

// misal

func song11(artist string, songs ...string) string {
	var lyric = fmt.Sprintf("lagu yang dinyanyikan oleh %s adalah ", artist)
	for _, song := range songs {
		lyric += song + " "
	}
	return lyric
}

func menghitung(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	var avg = float64(total) / float64(len(numbers))
	return avg
}

/*
bisa dilihat pada fungsi calculate(), parameter numbers dideklarasikan dengan disisipkan tanda titik tiga (...)
menandakan bahwa numbers adalah sebuah parameter variadic dengan tipe data int

fun calculate(numbers ...int) float64
dipanggil fungsi dilakukan, jumlah parameter yang disisipkan bisa banyak
var avg = calculate(2, 4, 4, 5, 4, 43,412,3213,12312)
nilai tiap parameter bisa diakses seperti cara pengaksesan tiap elemen slice pada contoh di atas metode dipilih
adalah for - range
for _, number := range numbers {
*/
