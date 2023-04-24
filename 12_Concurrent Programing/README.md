# **Concurrent Programing**
## Oleh: I Dewa Gde Putra Anga Biara

# Resume

* **Concurrent programming** adalah teknik untuk menyelesaikan masalah dengan melakukan banyak request atau proses kecil secara bersamaan. Dalam concurrency, satu sumber daya tidak dapat menjalankan banyak proses secara bersamaan, tetapi dengan menggunakan Golang, semua proses ini tampak berjalan secara bersamaan karena kecepatannya. Golang sangat mendukung concurrent programming dan memiliki keunggulan dalam hal ini.

* **Goroutines** adalah fungsi yang berjalan secara bersamaan atau concurrent dengan alur program utama. Program utama sendiri adalah goroutine utama dan goroutine lainnya yang dibuat akan berjalan di bawah goroutine utama. Dengan goroutine, kita dapat menjalankan banyak proses secara concurrent dengan menggunakan sintaks "go" diikuti dengan nama fungsi atau proses yang akan dijalankan secara concurrent. Selain itu, goroutine juga menyediakan fitur untuk menentukan maksimum unit proses atau thread yang digunakan dalam program, yang disebut "GOMAXPROCS".

* **Channel** adalah objek komunikasi yang digunakan oleh goroutine. Melalui channel, setiap goroutine dapat berkomunikasi dengan goroutine lainnya. Channel sebenarnya adalah blok data yang dapat ditransfer antar goroutine dan dapat ditulis dan dibaca oleh goroutine. Ketika satu goroutine menulis data ke suatu channel, data tersebut dapat dibaca oleh goroutine lain melalui channel yang sama. Ada dua jenis channel, yaitu unbuffered dan buffered channel. Unbuffered channel adalah channel yang tidak memiliki kapasitas, sementara buffered channel memiliki kapasitas.

 ---

# Latihan

## Praktikum – Concurrent Programing

- Problem 1

    1. Fungsi yang mencetak angka kelipatan x

        - Screenshot running
        
        ![runnnig](./screenshots/Tugas%201_mencetak%20kelipatan.png)

    2. Program yang mencetak bilangan kelipatan 3 dengan menerapkan goroutine dan channel

        - Screenshot running
        
        ![runnnig](./screenshots/Tugas%201_mencetak%20kelipatan%203.png)

    3. Program yang mencetak bilangan kelipatan 3 dengan menerapkan goroutine dan buffer channel

        - Screenshot running
        
        ![runnnig](./screenshots/Tugas%201_mencetak%20kelipatan%203%20dengan%20buffer.png)

    4. Program yang yang menerapkan mutex

        - Screenshot running
        
        ![runnnig](./screenshots/Tugas%201_penerapan%20mutex.png)

- Problem - Letter Frequency

  Fungsi `countLetters()` akan membuat channel untuk menampung hasil penghitungan frekuensi ketika dipanggil, dan penghitungan akan dilakukan secara _concurrent_. Pada fungsi main, fungsi dipanggil per-kalimat yang ada di paragraf dan kemudian semua channel keluaran dari penghitung frekuensi akan disatukan dengan fungsi `mergeCountChannel()` yang nilai keluarannya dapat dilihat setelah semua _goroutine_ selesai menghitung huruf sesuai kalimat yang diberikan

  - Source
        
      ![source](/11_Concurrent%20Programing/screenshots/source.png)

  - Screenshot running
    
      ![runnnig](/11_Concurrent%20Programing/screenshots/running.png)