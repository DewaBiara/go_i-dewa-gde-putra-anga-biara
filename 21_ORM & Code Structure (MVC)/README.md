# **ORM and Code Structure (MVC)**
## Oleh: I Dewa Gde Putra Anga Biara

# Resume

* **MVC (Model-View-Controller)** adalah sebuah pola dalam pengembangan perangkat lunak yang memisahkan atau membagi struktur code menjadi tiga bagian utama yaitu Model, View, dan Controller. `Model` merupakan bagian dari MVC yang bertanggung jawab dalam pengaturan, manipulasi, penyimpanan dan pengorganisasian data yang berhubungan dengan database. `View` merupakan bagian dari MVC yang bertanggung jawab dalam menampilkan data atau informasi dalam format tertentu untuk user. Sedangkan `Controller` merupakan bagian dari MVC yang berfungsi menghubungkan dan mengatur Model dan View agar dapat berinteraksi satu sama lain. Selain itu, Controller juga berisi logika bisnis dari aplikasi atau sistem yang dibuat untuk mengolah data.

* **ORM** atau *Objec Relationship Mapping* adalah teknik atau metode yang digunakan untuk menghubungkan antara tabel dalam database relational dengan object dalam pemrograman. Tujuan dari penggunaan ORM adalah untuk mempermudah proses query database yang seringkali membutuhkan waktu yang lama. Dengan ORM, developer dapat melakukan query dari object yang sudah dibuat, sehingga lebih efisien dalam mengakses data dalam database. Untuk membuat atau melakukan perubahan pada tabel dalam database, developer dapat menggunakan fitur `migration`. Dengan fitur ini, developer dapat melihat history tabel yang sudah dibuat dan dimodifikasi, sehingga memudahkan dalam manajemen tabel dalam database. Di bahasa pemrograman Golang, terdapat library ORM yang sangat populer dan sering digunakan, yaitu [gorm.io](https://gorm.io/).

* **Gorm.io** adalah sebuah library ORM yang dapat digunakan dalam pengembangan aplikasi dengan bahasa pemrograman Golang. Library ini menyediakan fitur lengkap ORM seperti CRUD, relasi antar tabel, auto migration, dan lain-lain. Selain itu, Gorm juga menyediakan berbagai jenis database dan konfigurasi koneksi yang berbeda-beda. Untuk menginstal Gorm, pengguna dapat menggunakan perintah `go get -u gorm.io/gorm` dan diperlukan driver dari tipe database yang digunakan, yang dapat diinstal dengan perintah `go get -u gorm.io/driver/nama-database`. Gorm sangat populer dan banyak digunakan dalam pengembangan aplikasi dengan Golang yang membutuhkan implementasi ORM.

 ---


# Latihan
## Praktikum - ORM and Code Structure (MVC)

-   Problem 1 - API CRUD User Using Database

    Full code can be found [here](/21_ORM%20%26%20Code%20Structure%20(MVC)/praktikum/01_API-CRUD-User/main.go)

    1. Get single user by id

        - Code

        ```go
           // Get user by id
           func GetUserController(c echo.Context) error {
               id, err := strconv.Atoi(c.Param("id"))
               if err != nil {
                   return echo.NewHTTPError(http.StatusBadRequest, err.Error())
               }

               var user User
               if err := DB.First(&user, id).Error; err != nil {
                   if err == gorm.ErrRecordNotFound {
                       return echo.NewHTTPError(http.StatusNotFound, "user not found")
                   }

                   return echo.NewHTTPError(http.StatusBadRequest, err.Error())
               }

               return c.JSON(http.StatusOK, map[string]interface{}{
                   "message": "success get user by id",
                   "user":    user,
               })
           }
        ```

        - Api Test

            - Success

                ![image](./screenshots/1_Get%20all%20users.png)

    2. Delete user by id

        - Code

        ```go
           // Delete user by id
           func DeleteUserController(c echo.Context) error {
               id, err := strconv.Atoi(c.Param("id"))
               if err != nil {
                   return echo.NewHTTPError(http.StatusBadRequest, err.Error())
               }

               result := DB.Delete(&User{}, id)
               if result.Error != nil {
                   return echo.NewHTTPError(http.StatusBadRequest, err.Error())
               }

               if result.RowsAffected == 0 {
                   return echo.NewHTTPError(http.StatusNotFound, "user not found")
               }

               return c.JSON(http.StatusOK, map[string]interface{}{
                   "message": "success delete user",
               })
           }
        ```

        - Api Test

            - Success

                ![image](./screenshots/1_Delete%20user.png)

    3. Update user by id

        - Code

            ```go
                // update user by id
                func UpdateUserController(c echo.Context) error {
                    id, err := strconv.Atoi(c.Param("id"))
                    if err != nil {
                        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
                    }

                    user := User{}
                    c.Bind(&user)

                    result := DB.Model(&User{}).Where("id = ?", id).Updates(user)
                    if result.Error != nil {
                        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
                    }

                    if result.RowsAffected == 0 {
                        return echo.NewHTTPError(http.StatusNotFound, "user not found")
                    }

                    return c.JSON(http.StatusOK, map[string]interface{}{
                        "message": "success update user",
                        "user":    user,
                    })
                }
            ```

        - Api Test

            - Success

                ![image](./screenshots/1_Put%20user.png)

-   Problem 2 - Structuring Project with Layered Architecture

    Full code can be found [here](./praktikum/02_layered_api)

    -   Code Structure

        ![image](./screenshots/02_Structured%20model.png)

    -   API Endpoint

        -   Part 1 - User routes

            1.  Get single user by id

                -   Success

                    ![image](./screenshots/1_Get%20user%20by%20id.png)

            2.  Create new user

                -   Success

                    ![image](./screenshots/1_Post%20user.png)

            3.  Delete user by id

                -   Success

                    ![image](./screenshots/1_Delete%20user.png)

            4.  Update user information

                -   Success

                    ![image](./screenshots/1_Put%20user.png)

        -   Part 2 - Book routes

            1. Get all book data

                - Success

                    ![image](./screenshots/2_Get%20all%20books.png)

            2. Get single book by id

                - Success

                    ![image](./screenshots/2_Get%20book%20by%20id.png)

            3. Create new book

                - Success

                    ![image](./screenshots/2_Post%20books.png)

            4. Delete book by id

                - Success

                    ![image](./screenshots/2_Delete%20books.png)

            5. Update book information

                - Success

                    ![image](./screenshots/2_Put%20books.png)

-   Problem 3 - Structuring Project with Layered Architecture With Relation

    Full code can be found [here](./praktikum/03_layered_api_blog/)

    -   Code Structure

        ![image](./screenshots/03_Structured%20model.png)

    -   API Endpoint

        -   Blogs Routes

            1. Get all blogs data

                - Success

                    ![image](./screenshots/3_Get%20all%20blogs.png)

            2. Get single blogs by id

                - Success

                    ![image](./screenshots/3_Get%20blogs%20by%20id.png)

            3. Create new blogs

                - Success

                    ![image](./screenshots/3_Post%20blogs.png)

            4. Delete blogs by id

                - Success

                    ![image](./screenshots/3_Delete%20blogs.png)

            5. Update blogs information

                - Success

                    ![image](./screenshots/3_Put%20blogs.png)