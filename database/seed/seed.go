package seed

import (
	"database/sql"
	"log"
)

func Seed(db *sql.DB) {
	queries := []string{
		`INSERT INTO categories (name, description)
		VALUES
		('Bilangan Bulat', "Berisi pertanyaan mengenai pengenalan bilangan bulat dan macam-macam operasinya."),
		('Statistika', "Berisi pertanyaan mengenai pengolahan data tunggal sederhana, seperti rata-rata, nilai tengah, dan nilai modus."),
		('FPB & KPK', "Berisi pertanyaan mengenai perhitungan FPB dan KPK antara 2 bilangan.");`,

		`INSERT INTO quizzes (category_id, question, correct_answer)
		VALUES
		(1, "Bilangan yang terletak antara -25 dan -47", "-36"),
		(1, "Pada garis bilangan, -12 terletak di ... bilangan -23", "sebelah kanan"),
		(1, "Ipung mula-mula berada pada titik nol (0). Setelah itu Ipung melangkah ke sebelah kiri sejauh 9 langkah. Lalu ia berjalan ke kiri lagi sebanyak 13 langkah. Apabila dinyatakan dalam kalimat matematika menjadi ...", "(-9) + (-13) = 22"),
		(1, "Suatu ruangan memiliki suhu -6 derajat Celcius. Sore harinya, suhu diturunkan 22 derajat Celcius. Pagi hari keesokannya, suhu ruangan dinaikkan kembali 18 derajat Celcius. Berapa suhu ruangan di pagi tersebut?", "-10 derajat Celcius"),
		(1, "Berapa hasil dari 1200 - 125 x 8 + 2?", "202"),
		(1, "Hasil dari 40 x 12 : (153 - 113) adalah ...", "12"),
		(1, "Hasil dari 298 + (-18) x 47 adalah ...", "-548"),
		(1, "Urutan bilangan bulat dari yang terkecil adalah ...", "-4, -3, -2, -1, 0, 1"),
		(1, "Urutan bilangan bulat dari yang terbesar adalah ...", "20, 15, 10, 0, -10, - 15"),
		(1, "Bilangan 32 terletak antara bilangan ... dan ...", "-24 dan 42");`,

		`INSERT INTO incorrect_answers (quiz_id, option_one, option_two)
		VALUES
		(1, "-20", "35"),
		(2, "sebelah kiri", "di bawah"),
		(3, "9 + 13 = 22", "9 + (-13) = -4"),
		(4, "10 derajat celcius", "12 derajat Celcius"),
		(5, "8602", "10750"),
		(6, "10", "11"),
		(7, "-846", "13160"),
		(8, "-4, -3, -2, -1, 1, 0", "-3, -2, -1, 0, 2, 1"),
		(9, "33, 28, 20, -10, -5, 0", "12, -11, 10, -9, 8, -7"),
		(10, "-32 dan 23", "-12 dan 23");`,

		`INSERT INTO quizzes (category_id, question, correct_answer)
		VALUES
		(2, "Pak Anton berjualan ayam potong di pasar setiap hari. Hasil penjualan ayam (dalam ekor) dari hari Senin sampai dengan hari Minggu adalah 10, 12, 9, 15, 18, 20, 21", "15 ekor"),
		(2, "Nilai ulangan Matematika siswa kelas V semester I adalah 6, 6, 8, 8, 8, 7, 7, 7, 8, 6, 6, 6, 8, 9, 10, 9, 9, 9, 9, 8 , 8, 6, 6, 8 , 7. Modul data tersebut adalah ...", "8"),
		(2, "Jumlah nilai siswa kelas VI adalah 160. Jika rata-rata nilai mereka adalah 8. Maka jumlah siswa kelas VI adalah ...", "20 orang"),
		(2, "Hasil panen kakek selama 5 bulan dalam ton adalah 10, 7, 7, 9, 8. Rata-rata hasil panen kakek tiap bulan adalah ... ton", "8,2"),
		(2, "Nilai median dari 10, 6, 7, 6, 8 adalah ...", "7"),
		(2, "Nilai median dari 3, 2, 4, 2, 1, 5", "2,5"),
		(2, "Bu Dewi membeli 1400 lembar kertas warna. Kemudian oleh Kepala Sekolah kertas tersebut ditambahkan 560 lembar. Kertas warna tersebut akan dibagikan kepada 35 muridnya sama banyak sebagai bahan kerajinan tangan. Banyak kertas warna yang diterima oleh masing-masing murid adalah ...", "56 lembar"),
		(2, "Data pekerjaan orang tua dari 300 siswa SD Jatisari adalah 80 orang POLRI, 160 orang wiraswasta, dan lainnya sebagai buruh. Modus jenis pekerjaan di atas adalah ...", "Wiraswasta"),
		(2, "Nilai ulangan IPA Ani adalah 6, 6, 7, 8, 9, 8, 6, 7. Modus nilai ulangan IPA Ani adalah ...", "Nilai 6"),
		(2, "Diketahui data: 50, 30, 70, 40, 80, 60, 60, 50, 60, 40. Nilai mediannya adalah ...", "55");`,

		`INSERT INTO incorrect_answers (quiz_id, option_one, option_two)
		VALUES
		(11, "12 ekor", "18 ekor"),
		(12, "6", "7"),
		(13, "15 orang", "25 orang"),
		(14, "7,5", "9"),
		(15, "10", "8"),
		(16, "3", "3,5"),
		(17, "57 lembar", "55 lembar"),
		(18, "Polri", "Buruh"),
		(19, "Nilai 8", "Nilai 7"),
		(20, "50", "60");`,

		`INSERT INTO quizzes (category_id, question, correct_answer)
		VALUES
		(3, "Nilai KPK dari 85, 90, 125 adalah ... ", "2250"),
		(3, "Tentukan KPK dari 10 dan 15!", "30"),
		(3, "Tentukan KPK dari 12, 15, dan 18!", "180"),
		(3, "Tentukan FPB dari 15 dan 25!", "5"),
		(3, "Tentukan FPB dari 40, 64, dan 100!", "4"),
		(3, "Adi berkunjung ke perpustakaan 3 hari sekali, Rina 5 hari sekali dan Riko 6 hari sekali. Jika mereka terkahir kali bersama-sama berkunjung pada tanggal 12 April 2017. Maka mereka akan berkunjung bersama-sama lagi pada tanggal ...", "12 Mei 2017"),
		(3, "Jam dinding di ruang makan berbunyi setiap 15 menit. Sedangkan Jam di ruang makan berbunyi setiap 20 menit. Kedua jam berbunyi bersamaan pertama kali pukul 12.30. Kedua jam berbunyi bersamaan lagi untuk kedua kali pada pukul ...", "13.30"),
		(3, "Ibu membeli 24 buah mangga, 30 buah apel dan 54 jeruk. Ibu ingin menyajikannya di ruang makan di atas piring. Berapa jumlah piring yang dibutuhkan jika ibu ingin membaginya dalam jumlah yang sama?", "6 piring"),
		(3, "Pak Guru telah membeli 60 buku pelajaran, 50 buku cerita dan 80 buku bergambar. Buku itu ingin dibuatkan rak buku. Pak guru nanti ingin meletakkan buku itu sama rata jumlahnya pada setiap rak. Berapakah rak buku yang dibutuhkan Pak Guru?", "10 rak"),
		(3, "Faktorisasi prima dari angka 100 adalah ...", "2 x 2 x 5 x 5");`,

		`INSERT INTO incorrect_answers (quiz_id, option_one, option_two)
		VALUES
		(21, "450", "1350"),
		(22, "15", "60"),
		(23, "150", "120"),
		(24, "1", "20"),
		(25, "8", "16"),
		(26, "28 Mei 2017", "21 Mei 2017"),
		(27, "13.15", "13.50"),
		(28, "12 piring", "18 piring"),
		(29, "20 rak", "15 rak"),
		(30, "10 x 10", "4 x 5 x 5");`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return
		}
	}

	log.Println("Successfully seeded all table")
}
