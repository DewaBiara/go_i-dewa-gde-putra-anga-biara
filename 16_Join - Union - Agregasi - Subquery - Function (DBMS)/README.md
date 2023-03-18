# **Join - Union - Agregasi - Subquery - Function (DBMS)**
## Oleh: I Dewa Gde Putra Anga Biara

# Resume

* **JOIN** adalah perintah yang dapat digunakan untuk menggabungkan data dari dua atau lebih tabel. Join memiliki beberapa jenis, seperti INNER JOIN, LEFT JOIN, dan RIGHT JOIN. INNER JOIN digunakan untuk mengambil record yang memenuhi kondisi pada kedua tabel yang dijoin. LEFT JOIN digunakan untuk mengambil record dari tabel kiri dan menampilkan data dari tabel kanan yang sesuai dengan kondisi join, namun record dari tabel kiri yang tidak sesuai dengan kondisi tetap akan ditampilkan. Sedangkan RIGHT JOIN adalah kebalikan dari LEFT JOIN, di mana record dari tabel kanan akan tetap ditampilkan meskipun tidak sesuai dengan kondisi join, namun record dari tabel kiri yang tidak sesuai dengan kondisi akan diabaikan.

* Dalam beberapa situasi, kita memerlukan perhitungan kelompok data seperti menghitung total jumlah baris dalam sebuah tabel, mencari nilai maksimum dari sebuah kolom pada tabel, atau bahkan menemukan nilai rata-rata dari sebuah kolom pada tabel. Untuk melakukan hal tersebut, kita dapat menggunakan perintah SQL yang disebut dengan **Fungsi Agregasi**. Fungsi agregasi digunakan untuk melakukan operasi pada kelompok-kelompok baris data, dan akan menghasilkan satu baris data untuk setiap kelompok baris data yang ada. Contoh fungsi agregasi yang sering digunakan adalah MIN, MAX, AVG, SUM, dan COUNT.

* Dalam database, konsep **Function** mirip dengan function dalam bahasa pemrograman, dimana Function dapat menjalankan suatu proses tertentu dan mengembalikan nilai sebagai hasil dari function. Sedangkan, **Trigger** hampir sama dengan Function, dengan perbedaan bahwa Trigger akan otomatis dijalankan ketika suatu kondisi terpenuhi, seperti sebelum atau sesudah melakukan operasi seperti insert, update, atau delete pada tabel. Trigger ini dapat melakukan proses yang diinginkan seperti menambah atau memanipulasi data dalam tabel yang bersangkutan.

 ---

# Latihan

## Praktikum – Join - Union - Agregasi - Subquery - Function (DBMS)

- Part 1 - DDL

  SQL file can be found [here](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/praktikum/DDL.sql) or open this;

    <details>
    <summary><b>Table DDL</b></summary>

    ```sql
    -- DDL
    CREATE TABLE users (
        user_id INT AUTO_INCREMENT,
        name VARCHAR(50),
        address VARCHAR(50),
        dob DATE,
        status_user SMALLINT,
        gender CHAR,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        CONSTRAINT PK_USER PRIMARY KEY(user_id)
    );

    CREATE TABLE product_types (
        type_id INT AUTO_INCREMENT,
        detail VARCHAR(50),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        CONSTRAINT PK_TYPE PRIMARY KEY(type_id)

    );

    CREATE TABLE product_description (
        product_id INT(11),
        description TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        CONSTRAINT product_description_fk FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE ON UPDATE CASCADE
    );


    CREATE TABLE operators (
        operator_id INT AUTO_INCREMENT,
        name VARCHAR(50),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        CONSTRAINT PK_OPERATOR PRIMARY KEY(operator_id)
    );

    CREATE TABLE payment_methods (
        payment_id INT AUTO_INCREMENT,
        name VARCHAR(50),
        status_type SMALLINT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        CONSTRAINT PK_PAYMENT PRIMARY KEY(payment_id)
    );

    CREATE TABLE products (
        product_id INT AUTO_INCREMENT,
        product_name VARCHAR(50),
        product_type INT,
        product_description TEXT,
        operator INT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        CONSTRAINT PK_PRODUCT PRIMARY KEY(product_id),
        CONSTRAINT FK_TYPE FOREIGN KEY (product_type) REFERENCES product_types(type_id) ON DELETE CASCADE ON UPDATE CASCADE,
        CONSTRAINT FK_OPERATOR FOREIGN KEY (operator) REFERENCES operators(operator_id) ON DELETE CASCADE ON UPDATE CASCADE
    );

    CREATE TABLE transactions (
        transaction_id INT AUTO_INCREMENT,
        user_id INT,
        payment_id INT,
        status VARCHAR(10),
        total INT,
        total_qty INT(11),
        total_price NUMERIC(25,2),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        CONSTRAINT PK_TX PRIMARY KEY(transaction_id),
        CONSTRAINT FK_USER FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
        CONSTRAINT FK_PAYID FOREIGN KEY (payment_id) REFERENCES payment_methods(payment_id) ON DELETE CASCADE ON UPDATE CASCADE
    );

    CREATE TABLE transaction_details (
        transaction_id INT, 
        product_id INT,
        status VARCHAR(10),
        qty INT,,
        price NUMERIC(25,2),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        CONSTRAINT FK_TXID FOREIGN KEY (transaction_id) REFERENCES transactions(transaction_id) ON DELETE CASCADE ON UPDATE CASCADE,
        CONSTRAINT FK_PRODUCTID FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE ON UPDATE CASCADE
    );

    ```

    </details>

- Part 2 - DML

  SQL file can be found [here](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/praktikum/DML.sql) or open this;

  1. **Insert**
    
     <details>
     <summary><b>Details</b></summary>

     - a. Insert 5 operator pada tabel operator

       - Source
          
         ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Insert_1_Insert%205%20operator%20pada%20tabel%20operator_code.png)

       - After run

         ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Insert_1_Insert%205%20operator%20pada%20tabel%20operator_run.png)
          
     - b. Insert 3 product type
        
       - Source
          
         ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Insert_2_Insert%203%20product%20type_code.png)

       - After run

         ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Insert_2_Insert%203%20product%20type_run.png)
          
     - c. Insert 2 product dengan product type id 1 dan operator id 3
        
       - Source
          
         ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Insert_3_Insert%202%20product%20dengan%20product%20type%20id%201%20dan%20operator%20id%203_code.png)

       - After run

         ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Insert_3_Insert%202%20product%20dengan%20product%20type%20id%201%20dan%20operator%20id%203_run.png)
          
     - d. Insert 3 product dengan product type id 2 dan operator id 1
    
       - Source
          
         ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Insert_4_Insert%203%20product%20dengan%20product%20type%20id%202%20dan%20operator%20id%201_code.png)

       - After run

         ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Insert_4_Insert%203%20product%20dengan%20product%20type%20id%202%20dan%20operator%20id%201_run.png)
          
     - e. Insert 3 product dengan product type id 3 dan operator id 4
  
       - Source
          
         ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Insert_5_Insert%203%20product%20dengan%20product%20type%20id%203%20dan%20operator%20id%204_code.png)

       - After run

         ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Insert_5_Insert%203%20product%20dengan%20product%20type%20id%203%20dan%20operator%20id%204_run.png)
          
     - f. Insert product description untuk masing-masing product
  
       - Source
          
         ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Insert_6_Insert%20product%20description%20untuk%20masing-masing%20product_code.png)

       - After run

         ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Insert_6_Insert%20product%20description%20untuk%20masing-masing%20product_run.png)
          
     - g. Insert 3 payment methods
  
       - Source
          
         ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Insert_7_Insert%203%20payment%20methods_code.png)

       - After run

         ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Insert_7_Insert%203%20payment%20methods_run.png)
          
     - h. Insert 5 user pada tabel users
  
       - Source
          
         ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Insert_8_Insert%205%20user%20pada%20tabel%20users_code.png)

       - After run

         ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Insert_8_Insert%205%20user%20pada%20tabel%20users_run.png)
          
     - i. Insert 3 transaksi masing-masing user
  
       - Source
          
         ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Insert_9_Insert%203%20transaksi%20masing-masing%20user_code.png)

       - After run

         ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Insert_9_Insert%203%20transaksi%20masing-masing%20user_run.png)
          
     - j. Insert 3 transaksi detail masing-masing transaksi
  
       - Source
          
         ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Insert_10_Insert%203%20product%20masing-masing%20transaksi_code.png)

       - After run

         ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Insert_10_Insert%203%20product%20masing-masing%20transaksi_run.png)
          
     </details>

  2. **Select**
    
     <details>
     <summary><b>Details</b></summary>

     - a. Tampilkan user dengan gender "M"
    
       - Source
          
         ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Select_1_Tampilkan%20user%20dengan%20gender%20M_code.png)

       - After run

         ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Select_1_Tampilkan%20user%20dengan%20gender%20M_run.png)

     - b. Tampilkan product dengan id = 3
      
       - Source
          
         ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Select_2_Tampilkan%20product%20dengan%20id%20%3D%203_code.png)

       - After run

         ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Select_2_Tampilkan%20product%20dengan%20id%20%3D%203_run.png)

     - c. Tampilkan pelanggan yang created_at dakan range 7 hari terakhir dan nama mengandung kata 'a'
        
       - Source
          
         ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Select_3_Tampilkan%20pelanggan%20yang%20created_at_code.png)

       - After run

         ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Select_3_Tampilkan%20pelanggan%20yang%20created_at_run.png)
         
     - d. Hitung jumlah user dengan gender "F"
        
       - Source
          
         ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Select_4_Hitung%20jumlah%20user%20dengan%20gender%20F_code.png)

       - After run

         ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Select_4_Hitung%20jumlah%20user%20dengan%20gender%20F_run.png)
         
     - e. Tampilkan data pelanggan dengan urutan sesuai nama abjad
        
       - Source
          
         ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Select_5_Tampilkan%20data%20pelanggan%20dengan%20urutan%20sesuai%20nama%20abjad_code.png)

       - After run

         ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Select_5_Tampilkan%20data%20pelanggan%20dengan%20urutan%20sesuai%20nama%20abjad_run.png)
         
     - f. Tampilkan 5 data pada data product
        
       - Source
          
         ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Select_5_Tampilkan%205%20data%20pada%20data%20product_code.png)

       - After run

         ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Select_6_Tampilkan%205%20data%20pada%20data%20product_run.png)
         

      </details>

  3. **Update**
      
     <details>
     <summary><b>Details</b></summary>

     - a. Ubah product id 1 menjadi 'product dummy'
        
       - Source
          
         ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Update_1_Ubah%20product%20id%201%20menjadi%20product%20dummy_code.png)

       - After run

         ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Update_1_Ubah%20product%20id%201%20menjadi%20product%20dummy_run.png)  

     - b. Update qty = 3 pada transaction detail dengan product id 1
          
       - Source
          
         ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Update_2_Update%20qty%20%3D%203%20pada%20transaction%20detail_code.png)

       - After run

         ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Update_2_Update%20qty%20%3D%203%20pada%20transaction%20detail_run.png)  
         
     </details>
  
  4. **Delete**
    
     <details>
     <summary><b>Details</b></summary>

     - a. Hapus product dengan id 1
            
       - Source
          
         ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Delete_1_Hapus%20product%20dengan%20id%201_code.png)

       - After run

         ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Delete_1_Hapus%20product%20dengan%20id%201_run.png) 

     - b. Hapus product dengan product type id 1
            
       - Source
          
         ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Delete_2_Hapus%20product%20dengan%20product%20type%20id%201_code.png)

       - After run

         ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/Delete_2_Hapus%20product%20dengan%20product%20type%20id%201_run.png) 

  </details>

- Part 3 - Join, Union, Subquery, Function       
  
  SQL file can be found [here](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/praktikum/Join-Union-Subquery-Function.sql) or open this;

  <details>
  <summary><b>Details</b></summary>

  1. Gabungkan data transaksi user id 1 dan user id 2
              
     - Source
          
       ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/1_Gabungkan%20data%20transaksi%20user%20id%201%20dan%20user%20id%202_code.png)

     - After run

       ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/1_Gabungkan%20data%20transaksi%20user%20id%201%20dan%20user%20id%202_run.png) 

  2. Tampilkan jumlah harga transaksi user id 1 
                
     - Source
          
       ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/2_Tampilkan%20jumlah%20harga%20transaksi%20user%20id%201_code.png)

     - After run

       ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/2_Tampilkan%20jumlah%20harga%20transaksi%20user%20id%201_run.png) 

  3. Tampilkan total transaksi dengan product type 2
                
     - Source
          
       ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/3_Tampilkan%20total%20transaksi%20dengan%20product%20type%202_code.png)

     - After run

       ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/3_Tampilkan%20total%20transaksi%20dengan%20product%20type%202_run.png) 

  4. Tampilkan semua field table product dan field name table product type yang saling berhubungan
                
     - Source
          
       ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/4_Tampilkan%20semua%20field%20table%20product%20dan%20field%20name_code.png)

     - After run

       ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/4_Tampilkan%20semua%20field%20table%20product%20dan%20field%20name_run.png) 

  5. Tampilkan semua field table transactions, field name table product dan field name table users
                
     - Source
          
       ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/5_Tampilkan%20semua%20field%20table%20transactions_code.png)

     - After run

       ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/5_Tampilkan%20semua%20field%20table%20transactions_run.png) 

  6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud
                
     - Source
          
       ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/6_Buat%20function%20setelah%20data%20transaksi%20dihapus%20maka%20transaction%20detail%20terhapus.png)

     - Before Delete

       ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/6_Before.png) 

     - After Delete
       Delete with query `DELETE FROM transactions WHERE id = 1;` 

       ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/6_After.png) 

  7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus
                
     - Source
          
       ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/7_Buat%20function%20setelah%20data%20transaksi%20detail%20dihapus.png)

     - Before Delete

       ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/7_Before.png) 

     - After Delete
       Delete with query `DELETE FROM transaction_details WHERE transaction_id = 2 AND product_id = 3;`

       ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/7_After.png) 

  8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan subquery
                
     - Source
          
       ![source](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/8_Tampilkan%20data%20products%20yang%20tidak%20pernah%20ada_code.png)

     - After run

       ![screenshot](/16_Join%20-%20Union%20-%20Agregasi%20-%20Subquery%20-%20Function%20(DBMS)/screenshots/8_Tampilkan%20data%20products%20yang%20tidak%20pernah%20ada_run.png) 