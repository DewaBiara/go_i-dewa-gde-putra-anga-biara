# **Intro RESTful API**
## Oleh: I Dewa Gde Putra Anga Biara

# Resume

* **API** merupakan singkatan dari `Application Programming Interface` yang merujuk pada kumpulan fungsi atau prosedur yang memungkinkan pengguna untuk mengakses fitur atau data dari sistem operasi, aplikasi, atau layanan tertentu. Cara kerja API dapat diilustrasikan seperti memesan makanan di restoran, di mana pelanggan (client) melakukan permintaan melalui pelayan (API), yang kemudian mengirimkan permintaan tersebut ke dapur (server) untuk diproses. Setelah selesai, dapur akan mengirimkan makanan (data) yang diminta kembali ke pelayan untuk disajikan kepada pelanggan sebagai respons. Salah satu jenis API yang umum digunakan adalah REST atau RESTful API.

* **REST** adalah singkatan dari `Representational State Transfer` dan merupakan suatu arsitektur perangkat lunak yang mengatur cara kerja API. Awalnya REST diciptakan sebagai panduan dalam mengatur komunikasi pada jaringan yang kompleks seperti internet. REST API menggunakan protokol HTTP sebagai cara komunikasi, yang seringkali melalui web browser. Selain itu, interface API biasanya berupa URL, path, atau endpoint yang dirancang agar mudah dimengerti dan digunakan ketika menggunakan API. Selanjutnya, dalam melakukan request di HTTP, terdapat berbagai jenis method seperti GET, POST, PUT, DELETE, dan lain-lain. Setelah melakukan request, API akan memberikan respons atau tanggapan dari request tersebut yang biasanya memiliki format sendiri seperti XML, SOAP, dan JSON yang merupakan format yang paling sering digunakan.

* **JSON** adalah format data yang digunakan untuk menyimpan atau mentransfer data dan singkatan dari `JavaScript Object Notation`. JSON memiliki kelebihan ukurannya yang lebih ringan dan struktur kodenya yang sederhana dan mudah dipahami. Dalam syntax JSON terdiri dari `key` dan `value`, dan terdapat enam jenis data yang dapat digunakan yaitu `Object, String, Array, Boolean, Number, dan NULL`. Dalam respons API saat ini, format JSON merupakan yang paling banyak digunakan karena ringan, sederhana, mudah diolah, dan memiliki format yang jelas dan mudah dipahami.

 ---

# Latihan

## Praktikum – Intro RESTful API

  - **Environment Variable**
    
    ![env](/19_Intro%20RESTful%20API/screenshots/Environment.png)
    
    > Exported Postman environment variable can be found [here](/19_Intro%20RESTful%20API/praktikum/Alterra%20Test.postman_environment.json)

  - **API Request**
    
    > Exported Postman collection can be found [here](/19_Intro%20RESTful%20API/praktikum/Alterra%20Task.postman_collection.json)
    1. JsonPlaceHolder API

       - [JsonPlaceHolder] Get All Data
         
         ![JsonPlaceHolder](/19_Intro%20RESTful%20API/screenshots/Get%20data%20API.png)

       - [JsonPlaceHolder] Get Data by ID 3
         
         ![JsonPlaceHolder](/19_Intro%20RESTful%20API/screenshots/Get%20data%20ID%203.png)

       - [JsonPlaceHolder] Add New Data
         
         ![JsonPlaceHolder](/19_Intro%20RESTful%20API/screenshots/Post%20Data.png)

       - [JsonPlaceHolder] Delete Data by ID
         
         ![JsonPlaceHolder](/19_Intro%20RESTful%20API/screenshots/Delete%20Data.png)

    2. RentABook API

       - [RentABook] Get Client by ID
         
         ![RentABook](/19_Intro%20RESTful%20API/screenshots/RentABook_Get%20Client%20By%20ID.png)

       - [RentABook] Get All Client
         
         ![RentABook](/19_Intro%20RESTful%20API/screenshots/RentABook_Get%20All%20Client.png)

       - [RentABook] Add New Client
         
         ![RentABook](/19_Intro%20RESTful%20API/screenshots/RentABook_Add%20New%20Client.png)

       - [RentABook] Update Client by ID
         
         ![RentABook](/19_Intro%20RESTful%20API/screenshots/RentABook_Update%20Client%20by%20ID.png)

       - [RentABook] Delete Client by ID
         
         ![RentABook](/19_Intro%20RESTful%20API/screenshots/RentABook_Delete%20Client%20by%20ID.png)