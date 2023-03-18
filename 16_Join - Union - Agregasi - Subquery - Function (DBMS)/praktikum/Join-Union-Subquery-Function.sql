-- 1. Gabungkan data transaksi user id 1 dan user id 2
SELECT * FROM transactions WHERE user_id = 1 OR user_id = 2;


-- 2. Tampilkan jumlah harga transaksi user id 1 
SELECT SUM(td.price * td.qty) FROM transactions t
JOIN transaction_details td ON t.transaction_id = td.transaction_id
WHERE t.user_id = 1;


-- 3. Tampilkan total transaksi dengan product type 2
SELECT SUM(td.qty * td.price) FROM transaction_details td
JOIN products p ON td.product_id = p.product_id
WHERE p.product_type_id = 2;


-- 4. Tampilkan semua field table product dan field name table product type yang saling berhubungan
SELECT p.*, pt.detail FROM products p
JOIN product_types pt ON p.product_type_id = pt.type_id;


-- 5. Tampilkan semua field table transactions, field name table product dan field name table users
SELECT t.*, p.product_name, u.name FROM transactions t
JOIN transaction_details td ON t.transaction_id = td.transaction_id
JOIN products p ON td.product_id = p.product_id
JOIN users u ON t.user_id = u.user_id;


-- 6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud
DELIMITER $$
$$
CREATE TRIGGER delete_transaction_details
AFTER DELETE ON transactions
FOR EACH ROW
BEGIN
	DECLARE v_transaction_id INT;
	SET v_transaction_id = OLD.transaction_id;
	DELETE FROM transaction_details WHERE transaction_id = v_transaction_id;
END

$$
DELIMITER ;


-- 7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus
DELIMITER $$
$$
CREATE TRIGGER update_total_qty
BEFORE DELETE ON transaction_details
FOR EACH ROW
BEGIN
    UPDATE transactions
    SET total_qty = total_qty - OLD.qty
    WHERE transaction_id = OLD.transaction_id;
END$$
DELIMITER ;


-- 8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan subquery
SELECT * FROM products WHERE product_id NOT IN (SELECT product_id FROM transaction_details);