# **Database - DDL - DML**
## Oleh: I Dewa Gde Putra Anga Biara

# Resume

* **Database** atau basis data adalah kumpulan data yang disusun berdasarkan aturan tertentu dan saling berhubungan agar mudah dalam pengelolaannya. Dengan menggunakan pengelolaan yang tepat, pengguna dapat dengan mudah mencari, menyimpan, memproses, dan menghapus data. DBMS atau Database Management System adalah sistem yang digunakan untuk mengelola basis data, dan MySQL merupakan salah satu jenis DBMS yang populer digunakan.

* **Data Definition Language (DDL)** merupakan serangkaian perintah untuk membuat atau menghapus database dan objek-objek yang terkait seperti tabel, view, procedure, trigger, dan tipe data. DDL digunakan untuk mengatur struktur dan skema dari database. Beberapa contoh perintah DDL adalah CREATE untuk membuat objek, ALTER untuk mengubah struktur objek, dan DROP untuk menghapus objek.

* **Data Manipulation Language (DML)** adalah kumpulan perintah yang digunakan untuk memanipulasi atau mengubah data yang tersimpan di dalam tabel. DML dirancang untuk meningkatkan interaksi pengguna dengan sistem dengan bahasa yang mudah dimengerti seperti bahasa Inggris sederhana. Perintah DML yang paling umum digunakan adalah SELECT untuk memilih data, UPDATE untuk melakukan perubahan data, INSERT INTO untuk menambahkan data ke dalam tabel, dan DELETE FROM untuk menghapus data dari tabel.

 ---

# Latihan

## Praktikum – Database - DDL - DML

- Part 1 - Schema Database

  - ERD Screenshot

  ![ERD](/12_Database%20-%20DDL%20-%20DML/screenshots/ERD_Digital%20Outlet%20Pulsa.jpg)

- Part 2 - Data Definition Language
  1. Create database alta_online_shop
     
    ![1](/14_Database%20-%20DDL%20-%20DML/screenshots/1_create_database.png)

  2. Implement ERD Schema
     
    ![2](/14_Database%20-%20DDL%20-%20DML/screenshots/2_Implementasi%20ERD_1.png)
    ![2](/14_Database%20-%20DDL%20-%20DML/screenshots/2_Implementasi%20ERD_2.png)

  3. Create `kurir` table

    ![3](/14_Database%20-%20DDL%20-%20DML/screenshots/3_Buat%20tabel%20kurir.png)


  4. Add `ongkos_dasar` field

    ![4](/14_Database%20-%20DDL%20-%20DML/screenshots/4_Tambah%20field%20ongkos%20dasar.png)

  5. Rename `kurir` to `shipping`

    ![5](/14_Database%20-%20DDL%20-%20DML/screenshots/5_Ganti%20nama%20tabel%20kurir%20menjadi%20shipping.png)

  6. Drop `shipping` table

    ![6](/14_Database%20-%20DDL%20-%20DML/screenshots/6_Hapus%20tabel%20shipping.png)

  7. Add new entity

     a. 1-to-1: payment method to description 
        
        ![7a](/14_Database%20-%20DDL%20-%20DML/screenshots/7_1-to-1.png)

     b. 1-to-many: user to address

        ![7b](/14_Database%20-%20DDL%20-%20DML/screenshots/7_1-to-many.png)

     c. many-to-many: user to payment method

        ![7c](/14_Database%20-%20DDL%20-%20DML/screenshots/7_many-to-many.png)