# **Data Structure**
## Oleh: I Dewa Gde Putra Anga Biara

# Resume

* **Array** adalah salah satu jenis struktur data yang sering digunakan. Array dapat menyimpan sejumlah elemen dalam satu wadah, dan ukuran array pada saat deklarasi tidak dapat diubah. Dalam bahasa pemrograman Golang, array dapat dideklarasikan dengan menggunakan sintaks ***var <nama_array> [ukuran_array] <tipe_variabel>***, dimana indeks pada array dimulai dari 0.

* **Slice** merupakan jenis struktur data yang mirip dengan array, namun lebih fleksibel dan kuat karena ukurannya dapat diubah sesuai kebutuhan. Dalam Golang, slice dapat dideklarasikan dengan menggunakan sintaks ***nama_slice := []tipe_data{values}***, dan juga dapat dibuat dari array yang sudah ada dengan sintaks ***nama_slice := nama_array[index_mulai:index_akhir]***. Selain itu, slice juga dapat dibuat dengan menggunakan sintaks make(), contohnya ***nama_slice := make([]tipe_data, panjang, kapasitas)***.

* **Map** adalah struktur data yang terdiri dari pasangan key:value. Setiap elemen dalam map terdiri dari kunci dan nilai yang dimilikinya. Map tidak memiliki urutan tertentu dan nilai-nilainya dapat diubah. Selain itu, map juga tidak mendukung duplikasi data. Dalam Golang, map dapat dideklarasikan dengan menggunakan sintaks ***var nama_map := map[tipe_key]tipe_value{key : value, ...}*** atau dengan menggunakan sintaks make(), contohnya ***nama_map := make(map[tipe_key]tipe_value)***.

 ---

# Latihan

## Praktikum – Data Structure