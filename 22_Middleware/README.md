# **ORM and Code Structure (MVC)**
## Oleh: I Dewa Gde Putra Anga Biara

# Resume

* **Middleware** adalah sebuah komponen yang terhubung dengan request dan response yang masuk ke dan keluar dari server. Dengan menggunakan middleware, kita dapat mengeksekusi fungsi atau metode tertentu sebelum request http masuk ke server atau setelah response http keluar dari server. Beberapa contoh middleware yang sering digunakan antara lain autentikasi, logging, dan sebagainya. Dengan menggunakan middleware, kita dapat menambahkan fitur tambahan pada aplikasi web tanpa harus memodifikasi kode inti dari aplikasi itu sendiri.

* **Implementasi Middleware** dapat di ilustrasikan oleh gambar berikut.

    ![ilustration](/22_Middleware/screenshot/middleware_ilustration.png)

    Dalam ilustrasi tersebut middleware akan berada diantara request dan respons yang dari aplication atau server. Middleware disini dapat berupa autentikasi, logging, session, dan middleware lain sesuai kebutuhan.

* **Echo Middleware** merupakan middleware-middleware yang sudah disediakan oleh echo yang dapat kita manfaatkan tanpa perlu membuatnya dari awal. Echo sendiri menyiadakan middleware untuk root level yaitu itu before router dan after router, dalam gorup level, dan route level.

 ---


# Latihan
## Praktikum - Middleware

-   Problem 1 - Logging & JWT Authentication

    Code can be found [here](./praktikum/1_JWT_Auth/)

    1.  Endpoint :

        -   Users

            | Route        | Method | Access            |
            | ------------ | ------ | ----------------- |
            | `/login`     | POST   | Not Authenticated |
            | `/users`     | GET    | Authenticated     |
            | `/users/:id` | GET    | Authenticated     |
            | `/users`     | POST   | Not Authenticated |
            | `/users/:id` | DELETE | Authenticated     |
            | `/users/:id` | PUT    | Authenticated     |

        -   Books

            | Route        | Method | Access            |
            | ------------ | ------ | ----------------- |
            | `/books`     | GET    | Authenticated     |
            | `/books/:id` | GET    | Authenticated     |
            | `/books`     | POST   | Authenticated     |
            | `/books/:id` | DELETE | Authenticated     |
            | `/books/:id` | PUT    | Authenticated     |

    2.  Common Error Response

        -   No JWT Token while accessing secured endpoint

            Response :

            ```json
            {
            	"message": "missing or malformed JWT"
            }
            ```

            Screenshot :

            ![no jwt token](./screenshot/1_No%20token.png)

        -   Invalid JWT Token while accessing secured endpoint

            Response :

            ```json
            {
            	"message": "invalid or expired JWT"
            }
            ```

            Screenshot :

            ![invalid jwt token](./screenshot/2_invalid%20jwt.png)

    3.  Response

        -   Users route

            1.  Login

                -   Success

                    ![login success](./screenshot/3_Login%20sucess.png)

                -   Failed (email or password is wrong)

                    ![login failed](./screenshot/4_Login%20failed.png)

            2.  Get all users

                ![get all users](./screenshot/5_Get%20all%20users.png)

            3.  Get user by ID

                -   Success

                    ![get user by id success](./screenshot/6_Get%20user%20success.png)

                -   Failed (requested user id is not found)

                    ![get user by id failed](./screenshot/7_Get%20user%20by%20id%20failed.png)

            4.  Create user

                ![create user success](./screenshot/8_Create%20user.png)

            5.  Delete user

                -   Success

                    ![delete user success](./screenshot/11_Delete%20user%20success.png)

                -   Failed (requested user id is not found)

                    ![delete user failed](./screenshot/12_Delete%20user%20failed.png)

            6.  Update user

                -   Success

                    ![update user success](./screenshot/9_Update%20user%20success.png)

                -   Failed (requested user id is not found)

                    ![update user failed](./screenshot/10_Update%20user%20failed.png)

        -   Books route

            1.  Get all books

                - Success

                ![get all books success](./screenshot/13_Get%20all%20book%20success.png)

                - Failed

                ![get all books failed](./screenshot/13_Get%20all%20book%20failed.png)

            2.  Get book by ID

                -   Success

                    ![get book by id success](./screenshot/14_Get%20book%20by%20id%20success.png)

                -   Failed (requested book id is not found)

                    ![get book by id failed](./screenshot/15_Get%20book%20by%20id%20failed.png)

            3.  Create book

                -   Success

                ![create book success](./screenshot/16_Create%20book%20success.png)

                -   Failed

                ![create book failed](./screenshot/17_Create%20book%20failed.png)

            4.  Delete book

                -   Success

                    ![delete book success](./screenshot/20_Delete%20book%20success.png)

                -   Failed (requested book id is not found)

                    ![delete book failed](./screenshot/21_Delete%20book%20failed.png)

            5.  Update book

                -   Success

                    ![update book success](./screenshot/18_Update%20book%20success.png)

                -   Failed (requested book id is not found)

                    ![update book failed](./screenshot/19_Update%20book%20failed.png)
