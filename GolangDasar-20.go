package main

import "fmt"

type person struct {
	student1
	name string
	age  int
}
type student1 struct {
	name  string
	grade int
}

func main() {
	// pointer
	/*
			pointer adalah refrence atau alamat memori variable pointer yang berisi alamat memory
			penerapan pointer ditandai dengan adanya astterisk (*) sebelum tipe data

			misal
			var number *int
			var name *string

			nila default variable pointer adalah nil
		    pointer tidak bisa menampung nilai yang bukan pointer
			dan sebaliknya variable biasa tidak bisa menampung nilai pointer

			point penting
			1. variable biasa diambil nilai pointernya caranya dengan menambahkan ampersand (&) sebelum nama variable ini disebut referencing
			2. dan untuk nilai asli variable pointer juga bisa diambil dengan cara menambahkan tanda asteerisk (*) sebelum nama variable pointer ini disebut dereferencing

	*/
	// contoh penerapan pointer

	var numberA int = 1000000
	var numberB *int = &numberA

	fmt.Println("numberA (value) ", numberA)
	fmt.Println("numberA (address) ", &numberA)
	fmt.Println("numberB (value) ", *numberB)

	// contoh pointer perubahan nilai
	numberA = 2000000
	fmt.Println("numberA (value) ", numberA)
	fmt.Println("numberA (address) ", &numberA)
	fmt.Println("numberB (value) ", *numberB)

	ganti(&numberA, 3000000)
	fmt.Println("numberA after (value) ", numberA)
	fmt.Println("numberA after (address) ", &numberA)

	// Struct
	/*
		Go tidak memiliki konsep OOP, namun Go memiliki tipe data struktur data yaitu struct
		struct merupakan kumpulan definisi variabel (atau property) dan fungsi (atau method) yang dibungkus sebagai satu tipe data baru dengan nama tertentu
		Property dalam struct ini tipe datanya dapat bervariasi artinya menampung semmua tipe data, seperti map
		hanya saja keynya sudah di definisikan di awal dan tipe data tiap itemnya bisa berbeda

		Struct ini dapat memiliki atribut sesuai skema struct tersebut
		kita sepakati bahwa variable disebut object atau variable object

		memanfaatkan struct, ini melakukan penyimpanan data yang sifatnya kolektif menjadi lebih mudah, rapih dan mudah untuk dikelola

		Cara Deklarasi Struct dengan kombinasi keyword type dan struct ini digunakan untuk deklarasi struct
	*/
	type student struct {
		name    string
		grade   int
		address string
		jurusan string
	}
	// jadi dari struct student memiliki 4 property yaitu name, grade, address, jurusan, property juga merupakan istilah untuk variable yang menempel ke struct
	// membuat penerapan struct
	var s1 student
	s1.name = "Rizky"
	s1.grade = 4.0
	s1.address = "Jakarta"
	s1.jurusan = "Teknik Informatika"

	s1.name = "Ayubs"

	fmt.Println("name : ", s1.name)
	fmt.Println("grade : ", s1.grade)
	fmt.Println("address : ", s1.address)
	fmt.Println("jurusan : ", s1.jurusan)
	// cara membuat variable object ini sama seperti pembuatan variable biasa, tinggal tulis saja nama variable yang diikuti nama struct

	// inisialisasi Object Struct
	/*
		caranya dengan menuliskan nama struct yang telah dibuat mengikuti dengan kurung kurawal, nilai masing masing property diisi pada saat inisialisasi
		contoh 3 buah object yang dideklarasikan berbeda
	*/
	var s2 = student{}
	s2.name = "Subagiya"
	s2.grade = 3.0

	var s3 = student{"Ayubs", 3, "Jakarta", "Teknik Informatika"}
	var s4 = student{name: "Riyan"}

	fmt.Println("student 1 : ", s2.name)
	fmt.Println("student 2 : ", s3.name)
	fmt.Println("student 3 : ", s4.name)

	/*
		s1 ini menampung object cetakan student, variable tersebut kemudian diset nilai propertynya

		variable object s2 dideklarasikan dengan metode sama seperti s1, bedanya di s2 nilai properynya di isi langsung ketika deklarasi, nilai pertama akan menjadi nilai property pertama (yaitu name)
		dan selanjutnya

		pada deklarasi s3 hanya name saja, ini efektif jika digunakan untuk membuat objek baru yang nilai propertynya tidak semua harus disiapkan di awal

		keistimewahan lainnya adalah penentuan nilai property
	*/
	var s5 = student{name: "wyne", grade: 2}
	var s6 = student{name: "Ayub Subagiya", grade: 4, address: "Bekasi", jurusan: "Teknik Informatika"}
	fmt.Println("student 5 : ", s5.name, s5.grade, s5.address, s5.jurusan)
	fmt.Println("student 6 : ", s6.name, s6.grade, s6.address, s6.jurusan)

	// Variabel Object Pointer
	/*
		dibuat dari tipe struct bisa diambil pointernya, dan bisa disimpan pada variable objek yang bertipe struct pointer
		contoh penerapannya
	*/

	var s7 = student{name: "SUbgaiya", grade: 3, address: "Jakarta", jurusan: "Teknik Informatika"}
	var s8 *student = &s7

	fmt.Println("student 7, name : ", s7.name, s7.grade, s7.address, s7.jurusan)
	fmt.Println("student 8, name : ", s8.name, s8.grade, s8.address, s8.jurusan)

	s7.name = "Ethan"
	fmt.Println("student 7, name : ", s7.name, s7.grade, s7.address, s7.jurusan)
	fmt.Println("student 8, name : ", s8.name, s8.grade, s8.address, s8.jurusan)
	/*
		s8 adalah varibale pointer dari hasil struct student, s8 menampung refrensi s7
		menjadikan setiap perubahan pada property variable tersebut, ini berpengaruh pada variable objek s7

		meskipun s8 bukan variable aslli, property tetap dapat di akses seperti biasa
		inilah keistimewahan property dalam object property tersebut tanpa deferencing
		nilai asli dari property tetap bisa diakses

		pengisian nilai dan property juga bisa langsung menggunakan nilai asli contohnya s8.name = "Ethan"

	*/

	// Embedded Struct
	/*
		Embedded struct adalah mekanisme untuk menempelkan sebuah struct sebagai properti struct lain, agar lebih musah dipahami

	*/
	var s9 = person{}
	var s10 = student1{}
	s9.name = "Naruto"
	s9.age = 21
	s9.student1.name = "Sasuke"
	s10.name = "Sakura"

	fmt.Println("name : ", s9.name)
	fmt.Println("age : ", s9.age)
	fmt.Println("name : ", s9.student1.name)
	fmt.Println("name : ", s10.name)

	// pengisian nilai sub-struct
	/*
		Sub-struct dilakukan dengan langsung memasukkan variable object yang tercetak dari struct yang sama
	*/
	var p1 = person{name: "Subagiya", age: 20}
	var p2 = student{
		name:    "Ayubs",
		grade:   20,
		address: "Bekasi",
		jurusan: "Teknik Informatika",
	}
	fmt.Println("name : ", p1.name)
	fmt.Println("age : ", p1.age)
	fmt.Println("name : ", p2.name)
	fmt.Println("grade : ", p2.grade)

	// anonymous struct yang tidak dideklarasikan di awal sebagai tipe data baru, melainkan langsung ketika pembuatan object, ini cukup efisien digunakan pada use case pembuatan variable object stuctnya dipakai sekali
	var s12 = struct {
		person
		student
		user string
		id   int
	}{}
	s12.person = person{name: "Subagiya", age: 20}
	s12.person = person{name: "Ayubs", age: 20}
	s12.student = student{name: "Ayubs", grade: 20, address: "Bekasi", jurusan: "Teknik Informatika"}

	fmt.Println(s12.person.age)
	fmt.Println(s12.person.name)
	/*
		s12 langsung diisi objek anonymous struct yang memiliki 4 property, person, student, user, id
		salah satu aturan yang diingat dalam pembuatan anonymous struct adalah, deklarasi harus diikuti dengan inisialisasi, bisa dilihat s12
		setelah deklarasi struktur struct, terdapat kurung kurawal untuk inisialisasi object
	*/

	// anonymous struct tanpa pengisian property
	var s13 = struct {
		person
		student
		nilai int
	}{}

	// anonymous struct dengan pengisian property
	var s14 = struct {
		person
		student
		student1
		nilai int
	}{
		person: person{
			student1: student1{},
			name:     "Ayubs",
			age:      20,
		},
		student: student{
			name:  "Ayubs",
			grade: 20,
		},
	}

	fmt.Println(s13)
	fmt.Println(s14)

	// kombinasi Slice dan Struct
	/*
		slice dan struct dikombinasikan seperti slice dan map,
		caranya cukup tambahkan tanda [] sebelum tipe data saat deklarasi
	*/
	var allStudents = []student{
		{name: "Subagiya", grade: 3},
		{name: "Ayubs", grade: 4},
		{name: "Rizky", grade: 3},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, " - ", student.grade)
	}

	// inisialisasi slice anonymous struct
	/*
		dapat dijadikan sebagai tipe sebuah slice, dan nilai awalnya juga diinisialisasi langsung saat deklarasi

	*/

	var allPerson = []struct {
		person
		student
	}{
		{person: person{name: "Subagiya", age: 20}, student: student{name: "Ayubs", grade: 3}},
		{person: person{name: "Rizky", age: 20}, student: student{name: "Ayubs", grade: 3}},
		{person: person{name: "Ayubs", age: 20}, student: student{name: "Ayubs", grade: 3}},
	}
	for _, person := range allPerson {
		fmt.Println(person.person.name, " - ", person.student.name, " - ", person.student.grade)
	}

	// deklarasi Anonymous Struct dengan Keyword var
	/*
		dengan menggunakan keyword var
	*/
	var student15 struct {
		person
		student
		grade int
	}
	student15.person = person{name: "Subagiya", age: 20}
	student15.grade = 21
	/*
		statement type student struct adalah contoh cara deklarasi struct, maknanya akan berbedaketika keyword type diganti var
		seperti pada contoh diatas, deklarasi struct diawali dengan var, kemudian diikuti dengan nama variable dan tipe data struct yang diinginkan
		sebuah object dari anonymous struct kemudian di simpan pada variable bernama student15

		deklarasi anonymous struct dengan keyword var ini bisa digunakan jika object tersebut akan digunakan lebih dari satu kali
	*/
	// hanya deklarasi struct
	var student2 struct {
		person
		student
		grade int
	}
	var student3 = struct {
		name string
		age  int
	}{
		"ayub Subagiya", 20,
	}
	fmt.Println(student3)
	fmt.Println(student2)

	// Nested Struct
	/*
		ini merupakan anonymous struct yang dideklarasikan atau di embed di dalam struct lain
		struct peng-embed disebut sebagai nested struct
	*/
	type student16 struct {
		person12 struct {
			name string
			age  int
		}
		grade   int
		hobbies []string
	}
	// digunakan ketika decoding data JSON yang struktur datanya cukup

	// Deklarasi dan inisiaisasi struct secara horizontal
	type person13 struct {
		name    string
		age     int
		hobbies []string
	}
	// menggunakan tanda semicolon (;) untuk memisahkan property struct
	var p3 = struct {
		name    string
		age     int
		hobbies []string
	}{name: "Subagiya", age: 20, hobbies: []string{"Game", "Coding"}}
	var p4 = struct {
		name    string
		age     int
		hobbies []string
	}{"Subagiya", 20, []string{"Game", "Coding"}}

	fmt.Println(p3)
	fmt.Println(p4)

	// tag property dalam struct
	/*
		berupa informasi opsional ayng ditambahkan pada property struct
	*/
	type person struct {
		name string `tag1`
		age  int    `tag2`
		// ini untuk keperluan encode/decode data, informasi tag juga bisa diakses lewat reflect

	}

	// Type Alias
	/*
		tipe data seperti struct dibuatkan alias baru
	*/
	type smartPhone struct {
		merk string
		ram  int
	}
	type phone = smartPhone
	// phone dan smartPhone memiliki struktur yang sama
	var p5 = smartPhone{
		merk: "Infinix",
		ram:  16,
	}
	fmt.Println(p5)
	var p6 = smartPhone{"Samsung", 32}
	fmt.Println(p6)

}

func ganti(original *int, value int) {
	*original = value
}
