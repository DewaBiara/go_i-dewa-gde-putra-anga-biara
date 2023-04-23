# **Clean and Hexagonal Architecture**
## Oleh: I Dewa Gde Putra Anga Biara

# Resume

* Saat merancang sebuah aplikasi, kita perlu memperhatikan pola perancangan dan pondasi yang kuat, mirip dengan merancang atau membangun sebuah gedung. Hal ini diperlukan untuk memastikan bahwa aplikasi dapat berjalan dengan stabil. Sebagaimana halnya dengan gedung, sebuah aplikasi juga memerlukan sebuah pondasi atau kerangka kerja (framework) yang kokoh, yang dalam hal ini disebut sebagai arsitektur. Arsitektur memiliki banyak manfaat, seperti memberikan standar yang baik untuk aplikasi, mendukung pengembangan aplikasi secara berkelanjutan, memudahkan pengujian unit, dan lain-lain. Ada banyak jenis arsitektur pengembangan aplikasi yang tersedia, seperti `clean architecture` dan `hexagonal architecture`.

* **Clean Architecture** merupakan sebuah solusi tambahan dari Design Pattern yang saat ini digunakan dalam pengembangan aplikasi. Hal ini membantu agar kode yang dibuat lebih mudah dibaca (readable), mudah di-maintain (maintainable), mudah diuji (testable), dan mudah digunakan kembali (reusable). Dalam clean architecture, setiap komponen aplikasi berfungsi secara independen dan dikelompokkan atau dipisahkan berdasarkan fungsinya.

* **Hexagonal Architecture** adalah sebuah pola desain yang sering digunakan dalam pengembangan mikro-servis. Pada dasarnya, hexagonal architecture memisahkan input dan output dari aplikasi dengan logika yang berjalan di dalam aplikasi. Dalam arsitektur ini, untuk menghubungkan input dan output dengan logika aplikasi, biasanya digunakan port-port sebagai penghubung.

 ---

# Latihan

## Praktikum - Clean and Hexagonal Architecture

1.  Problem 1 - Rewrite Code ke Clean Architecture

    Kode akan diubah dengan menggunakan _clean architecture per feature module_. Hal ini dilakukan untuk memisahkan setiap domain dan memungkinkan setiap domain memiliki implementasi yang terpisah seperti _layer usecase_, _layer controller_, dan _layer driver_. Selain itu, struktur kode akan mengikuti _go standard layout_ dengan memisahkan package menjadi dua folder yaitu `internal` yang berisi _business logic_ dari aplikasi dan `pkg` yang berisi kode pendukung. Dengan demikian, struktur folder akan terlihat seperti contoh di bawah ini:

    ![Clean Architecture](./screenshot/1_rewrite%20code.png)

2.  Problem 2 - Implementasi JWT 
    
    Impleemntasi JWT dilakukan dengan menggunakan middleware JWT pada routing di masing-masing yang memerlukan. Contoh implementasi pada domain `user` pada fungsi `InitRoute()`;

    ```go

    // Routes with authentication
	secure := e.Group("")
	secure.Use(middleware.JWT([]byte(config.JWT_SECRET)))

	secure.GET("/users", u.GetAllUser)

    ```

3.  Problem 3 - Unit Testing

    Untuk melakukan _unit testing_ pada setiap _layer_ pada domain, akan dilakukan secara terpisah. Pada setiap _layer_ yang memerlukan implementasi dari _layer_ yang berada di bawahnya, akan dilakukan teknik _mocking_ untuk memisahkan dan mengisolasi _layer_ tersebut saat dilakukan _unit testing_. Hasil dari _unit testing_ dapat dilihat seperti contoh di bawah ini:

    ![Unit Testing](./screenshot/3_unit%20test.png)