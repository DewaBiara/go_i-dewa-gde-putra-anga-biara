# **System Design**
## Oleh: I Dewa Gde Putra Anga Biara

# Resume

* **Diagram** adalah simbol atau bentuk yang digunakan untuk merepresentasikan informasi dengan tujuan memberikan pemahaman atau penjelasan dalam suatu konteks tertentu dalam desain sistem. Beberapa contoh diagram yang umum digunakan dalam sistem adalah flowchart, use case, ERD, dan HLA. **Flowchart** adalah diagram yang digunakan untuk menggambarkan alur kerja sistem secara visual. Use case menggambarkan interaksi antara pengguna atau user dengan sistem. ERD, singkatan dari Entity Relationship Diagram, adalah diagram yang digunakan untuk menggambarkan struktur atau desain dari database dalam menyimpan data. Sedangkan HLA, singkatan dari High Level Architecture, adalah standar untuk simulasi terdistribusi yang digunakan ketika membangun simulasi untuk tujuan yang lebih besar dengan menggabungkan beberapa simulasi.

* **Monolithic architecture** adalah sebuah pendekatan tradisional dalam arsitektur aplikasi yang menggabungkan semua kode menjadi satu kesatuan yang terikat dengan kode lainnya. Pendekatan ini memiliki kekurangan jika terjadi perubahan pada sistem, maka seluruh sistem harus diperbarui secara bersamaan. Selain itu, jika terjadi error pada salah satu menu atau modul, maka seluruh sistem akan mengalami masalah atau down. Untuk mengatasi kekurangan tersebut, dapat digunakan pendekatan Microservices architecture. Berbeda dengan monolithic, pendekatan microservices membagi aplikasi menjadi unit-unit yang lebih kecil dan spesifik. Setiap unitnya terpisah dan memiliki sistem beserta database sendiri untuk beroperasi, dan menggunakan mekanisme API untuk terhubung dengan unit lainnya. Kelebihan dari pendekatan microservices adalah mampu menutupi kekurangan dari monolithic. Namun, kelemahannya adalah tingkat kompleksitas dalam membangun dan melakukan pengujian pada aplikasi yang menggunakan pendekatan ini lebih tinggi.

* Terdapat dua jenis desain database yaitu SQL dan NoSQL. SQL adalah kependekan dari Structured Query Language yang sebenarnya digunakan sebagai bahasa untuk mengelola database relasional. Database relasional memiliki struktur yang konkret berbentuk tabel dengan field, tipe data, dan nilai di dalamnya. Yang paling penting dari database relasional adalah setiap tabel dapat membentuk relasi atau hubungan dengan tabel lainnya. Contoh database relasional adalah MySQL, PosgreSQL, Oracle, dan lain-lain. Sedangkan NoSQL adalah kebalikan dari SQL yang tidak memiliki relasi yang menghubungkan satu schema atau tabel dengan yang lain. Bentuk dari NoSQL untuk menyimpan data tidak berupa tabel, melainkan dapat berupa dokumen dalam bentuk JSON yang hanya berisi key dan value. Kelebihan dari NoSQL adalah performa yang cepat dan tidak memiliki relasi sehingga dapat menyimpan berbagai model data dalam satu dokumen. Contoh NoSQL seperti MongoDB, CloudDB, dan lain-lain.

 ---

# Latihan

## Praktikum – System Design

- Problem 1 - Diagram

  - ERD

  ![ERD](/18_System%20Design/screenshots/ERD.png)

  - Use Case Diagram
  
  ![Use Case Diagram](/18_System%20Design/screenshots/Use%20Case.png)

- Problem 2 - Query

  - SQL 

    `SELECT * FROM users;`

  - Redis
    
    - Storing in Set

      `SMEMBERS users`

    - Storing in List

      `LRANGE users 0 -1`
    
  - Neo4j

    ```
    MATCH (u:users)
    RETURN u;
    ```

  - Cassandra

    `SELECT * FROM users;`
  