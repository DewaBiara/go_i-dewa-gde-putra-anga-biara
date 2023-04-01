# **Intro Echo Golang**
## Oleh: I Dewa Gde Putra Anga Biara

# Resume

* **Library atau Third Party** adalah kumpulan kode yang memiliki fungsi tertentu yang dapat dipanggil oleh program atau proyek lain. Menggunakan library mempermudah pengembang untuk menyelesaikan masalah atau proyek dengan memanfaatkan fungsi yang telah disediakan dalam library tersebut. Selain memberikan kemudahan, penggunaan library juga dapat mengurangi beban pikiran karena kita tidak perlu membuat fungsi-fungsi tersebut dari awal. Beberapa contoh third party pada bahasa pemrograman Golang antara lain `gorm.io`, `logrus`, `cobra`, `gin`, `echo`, dan lain-lain.I.

* **Echo** adalah salah satu third party Golang yang berguna untuk mengembangkan web atau REST API. Echo adalah sebuah web framework yang memiliki performa yang tinggi dibandingkan dengan framework lainnya, serta memiliki sifat yang extensible dan minimalist. Minimalist di sini berarti Echo tidak menyediakan fungsi untuk driver database atau ORM, sehingga dengan sifatnya yang extensible, pengguna dapat mengintegrasikannya dengan third party lain seperti `gorm.io`.

* **Echo** memiliki beberapa kelebihan yang dapat membantu para pengembang. **Optimized router** pada Echo mengacu pada kemampuannya untuk mengoptimalkan router sehingga lebih powerful. Echo juga menyediakan banyak **middleware** yang dapat disesuaikan dengan kebutuhan pengguna. **Data rendering** pada Echo dapat di-custom format atau bentuknya sesuai dengan kebutuhan respons dari permintaan (request). Echo dapat di-scaling baik dalam skala kecil maupun besar dengan metode yang relatif sama. Terakhir, Echo memungkinkan pengguna untuk dengan mudah menangkap atau mendapatkan data dari klien (client) yang dibawa ke dalam struktur (struct) yang sudah dibuat untuk memetakan data, disebut sebagai data binding.

 ---

# Latihan

## Praktikum – Intro Echo Golang

- **Create Static API CRUD User**

Full code can be found [here](/20_Intro%20Echo%20Golang/praktikum/API-CRUD-User/main.go)

  1. Get specific user by id

     - Code
      
          ![GET](/20_Intro%20Echo%20Golang/screenshots/Code_Get%20By%20ID.png)

     - Screenshot

          ![GET](/20_Intro%20Echo%20Golang/screenshots/GET%20BY%20ID.png)

  2. Update user by id

     - Code
      
          ![UPDATE](/20_Intro%20Echo%20Golang/screenshots/Code_Update%20User%20By%20ID.png)

     - Screenshot

          ![UPDATE](/20_Intro%20Echo%20Golang/screenshots/PUT.png)

  3. Delete user by id
  
     - Code
      
          ![DELETE](/20_Intro%20Echo%20Golang/screenshots/Code_Delete%20By%20ID.png)

     - Screenshot

          ![DELETE](/20_Intro%20Echo%20Golang/screenshots/DELETE.png)