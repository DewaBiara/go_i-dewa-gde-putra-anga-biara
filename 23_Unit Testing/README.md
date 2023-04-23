# **Unit Testing**
## Oleh: I Dewa Gde Putra Anga Biara

# Resume

* **Software** testing adalah sebuah teknik yang digunakan untuk melakukan verifikasi dan validasi pada sebuah aplikasi guna mengecek apakah sudah sesuai dengan persyaratan yang telah ditetapkan pada rancangan aplikasi. Tujuan dari testing adalah untuk memastikan produk yang dihasilkan tidak memiliki cacat atau bug. Ada berbagai jenis software testing seperti manual testing, automation testing, performance testing, regression testing, statistic testing, dan dynamic testing.

* **Unit testing** adalah level pengujian yang dilakukan pada bagian terkecil dari suatu aplikasi, seperti kode, fungsi, metode, prosedur, modul, atau objek tersendiri. Level testing lainnya adalah UI testing yang melibatkan interaksi pengguna dengan aplikasi dan integration testing yang melibatkan pengujian pada lingkup modular atau subsistem. Unit testing sangat penting karena merupakan tahap awal dalam memastikan fungsionalitas dari bagian terkecil aplikasi sebelum melakukan pengujian pada level yang lebih tinggi.

* **Golang** memiliki package testing yang menyediakan banyak tools untuk keperluan unit testing. Untuk membuat unit test di Golang, kita perlu membuat file dengan format namafile_test.go dan menggunakan package yang sama dengan nama file yang akan diuji. Kemudian, kita dapat menjalankan unit testnya dengan menggunakan perintah go test. Selain itu, package testing di Golang juga menyediakan banyak method yang dapat digunakan untuk mendapatkan hasil testing sesuai dengan data yang dibutuhkan.

 ---

# Latihan

## Praktikum - Unit Testing

-   Problem 1 - Simple Unit Testing

    1.  Function 
            
        ```go
        func Addition(a, b float64) float64 {
            return a + b
        }

        func Subtraction(a, b float64) float64 {
            return a - b
        }

        func Division(a, b float64) float64 {
            return a / b
        }

        func Multiplication(a, b float64) float64 {
            return a * b
        }
        ```

    2.  Test Cases

        Addtion:
        -   Name: `Test addition int values`
            Input: `a = 1, b = 2`
            Expected Output: `3`
        -   Name: `Test addition float and int values`
            Input: `a = 1.5, b = 2`
            Expected Output: `3.5`
        -   Name: `Test addition float values`
            Input: `a = 1.5, b = 2.5`
            Expected Output: `4`
        
        Subtraction:
        -   Name: `Test subtraction int values`
            Input: `a = 1, b = 2`
            Expected Output: `-1`
        -   Name: `Test subtraction float and int values`
            Input: `a = 1.5, b = 2`
            Expected Output: `-0.5`
        -   Name: `Test subtraction float values`
            Input: `a = 1.5, b = 2.5`
            Expected Output: `-1`
        
        Division:
        -   Name: `Test division int values`
            Input: `a = 1, b = 2`
            Expected Output: `0.5`
        -   Name: `Test division float and int values`
            Input: `a = 1.5, b = 2`
            Expected Output: `0.75`
        -   Name: `Test division float values`
            Input: `a = 1.5, b = 2.5`
            Expected Output: `0.6`

        Multiplication:
        -   Name: `Test multiplication int values`
            Input: `a = 1, b = 2`
            Expected Output: `2`
        -   Name: `Test multiplication float and int values`
            Input: `a = 1.5, b = 2`
            Expected Output: `3`
        -   Name: `Test multiplication float values`
            Input: `a = 1.5, b = 2.5`
            Expected Output: `3.75`

    3.  Test Code

        Test code dapat dilihat pada bagian [ini](./praktikum/1_calculator/lib/calculate/calculate_test.go)

    4.  Test Result

        -   Report
            
            ![Test Report](./screenshot/1_test%20calculator.png)

-   Problem 2 - RESTful API Unit Testing

    -   Report

        ![Test Report](./screenshot/2_test%20rest%20api.png)
