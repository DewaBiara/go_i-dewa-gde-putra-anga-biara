-- 1. Insert
--    a. Insert 5 operator pada tabel operator
INSERT INTO operators (operator_id, name) 
VALUES 
    (1, 'Telkomsel'),
    (2, 'Indosat'),
    (3, 'XL'),
    (4, 'Axis'),
    (5, 'Tri');


--    b. Insert 3 product type
INSERT INTO product_types (type_id, detail)
VALUES
    (1, 'Pulsa'),
    (2, 'Paket Internet'),
    (3, 'Paket Games');


--    c. Insert 2 product dengan product type id 1 dan operator id 3
INSERT INTO products (product_type_id, operator_id, product_name)
VALUES
    (1, 3, 'Pulsa 12.000'),
    (1, 3, 'Pulsa 22.000');


--    d. Insert 3 product dengan product type id 2 dan operator id 1
INSERT INTO products (product_type_id, operator_id, product_name)
VALUES
    (2, 1, 'Paket Internet 1GB'),
    (2, 1, 'Paket Internet 5GB'),
    (2, 1, 'Paket Internet 10GB');


--    e. Insert 3 product dengan product type id 3 dan operator id 4
INSERT INTO products (product_type_id, operator_id, product_name)
VALUES
    (3, 4, 'Paket Games 1GB'),
    (3, 4, 'Paket Games 5GB'),
    (3, 4, 'Paket Games 10GB');


--    f. Insert product description untuk masing-masing product
INSERT INTO product_description (product_id, description)
VALUES
    (1, 'Pulsa 12.000'),
    (2, 'Pulsa 22.000'),
    (3, 'Paket Internet 1GB'),
    (4, 'Paket Internet 5GB'),
    (5, 'Paket Internet 10GB'),
    (6, 'Paket Games 1GB'),
    (7, 'Paket Games 5GB'),
    (8, 'Paket Games 10GB');


--    g. Insert 3 payment methods
INSERT INTO payment_methods (name)
VALUES
    ('BCA'),
    ('BNI'),
    ('BPD');


--    h. Insert 5 user pada tabel users
INSERT INTO users (user_id, name, dob, status_user, gender)
VALUES
    (1, 'Dewa', '2000-01-01', 1, 'M'),
    (2, 'Biara', '2000-01-01', 1, 'M'),
    (3, 'Budi', '2000-01-01', 1, 'M'),
    (4, 'Agus', '2000-01-01', 1, 'M'),
    (5, 'Rian', '2000-01-01', 1, 'M');
    


--    i. Insert 3 transaksi masing-masing user
INSERT INTO transactions (transaction_id, user_id, payment_id, status, total_qty, total_price)
VALUES
    (1,  1, 1, 1, 2, 20000),
    (2,  1, 1, 1, 5, 50000),
    (3,  1, 1, 1, 4, 40000),
    (4,  2, 2, 1, 1, 10000),
    (5,  2, 2, 1, 3, 50000),
    (6,  2, 2, 1, 4, 80000),
    (7,  3, 3, 1, 5, 50000),
    (8,  3, 3, 1, 6, 60000),
    (9,  3, 3, 1, 2, 30000),
    (10, 4, 1, 1, 3, 40000),
    (11, 4, 2, 1, 4, 40000),
    (12, 4, 2, 1, 4, 50000),
    (13, 5, 3, 1, 2, 20000),
    (14, 5, 3, 1, 4, 40000),
    (15, 5, 1, 1, 5, 50000);

   
--    j. Insert 3 product masing-masing transaksi
INSERT INTO transaction_details (transaction_id, product_id, qty, price)
VALUES
    (1,  1, 1, 10000),
    (1,  2, 2, 20000),
    (1,  5, 1, 10000),
    (2,  1, 2, 10000),
    (2,  3, 2, 10000),
    (2,  6, 2, 10000),
    (3,  1, 1, 10000),
    (3,  3, 1, 10000),
    (3,  6, 1, 10000),
    (4,  1, 1, 10000),
    (4,  3, 2, 10000),
    (4,  6, 1, 10000),
    (5,  1, 1, 10000),
    (5,  2, 1, 20000),
    (5,  6, 2, 10000),
    (6,  7, 1, 20000),
    (6,  2, 1, 20000),
    (6,  4, 1, 20000),
    (7,  1, 1, 10000),
    (7,  3, 1, 10000),
    (7,  6, 1, 10000),
    (8,  1, 1, 10000),
    (8,  3, 1, 10000),
    (8,  6, 1, 10000),
    (9,  1, 1, 10000),
    (9,  4, 1, 20000),
    (9,  6, 1, 10000),
    (10,  1, 1, 10000),
    (10,  3, 2, 10000),
    (10,  7, 1, 20000),
    (11,  1, 1, 10000),
    (11,  3, 1, 10000),
    (11,  6, 1, 10000),
    (12,  2, 1, 20000),
    (12,  3, 1, 10000),
    (12,  6, 1, 10000),
    (13,  1, 1, 10000),
    (13,  3, 1, 10000),
    (13,  6, 1, 10000),
    (14,  1, 1, 10000),
    (14,  3, 2, 10000),
    (14,  6, 1, 10000),
    (15,  1, 2, 10000),
    (15,  4, 1, 20000),
    (15,  6, 1, 10000);
    
   
-- 2. Select
--    a. Tampilkan user dengan gender "M"
SELECT * FROM users WHERE gender = 'M';


--    b. Tampilkan product dengan id = 3
SELECT * FROM products WHERE product_id = 3;


--    c. Tampilkan pelanggan yang created_at dakan range 7 hari terakhir dan nama mengandung kata 'a'
SELECT * FROM users WHERE created_at BETWEEN DATE_SUB(NOW(), INTERVAL 7 DAY) AND NOW() AND name LIKE '%a%';


--    d. Hitung jumlah user dengan gender "F"
SELECT COUNT(*) FROM users WHERE gender = 'F';


--    e. Tampilkan data pelanggan dengan urutan sesuai nama abjad
SELECT * FROM users ORDER BY name ASC;


--    f. Tampilkan 5 data pada data product
SELECT * FROM products LIMIT 5;


-- 3. Update
--    a. Ubah product id 1 menjadi 'product dummy'
UPDATE products SET product_name = 'product dummy' WHERE product_id = 1;


--    b. Update qty = 3 pada transaction detail dengan product id 1
UPDATE transaction_details SET qty = 3 WHERE product_id = 1;


-- 4. Delete
--    a. Hapus product dengan id 1
DELETE FROM products WHERE product_id = 1;


--    b. Hapus product dengan product type id 1
DELETE FROM products WHERE product_type_id = 1;   