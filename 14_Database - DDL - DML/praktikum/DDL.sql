-- 1. Create database 
CREATE DATABASE alta_online_shop;

-- Use database
USE alta_online_shop;

-- 2. Create users, products, product_types, opeartors, payment_methods, transactions and transaction_details table
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

-- 3. Create kurir table
CREATE TABLE kurir (
	id_kurir INT AUTO_INCREMENT,
	name VARCHAR(50),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	CONSTRAINT PK_KURIR PRIMARY KEY(id_kurir)
);

-- 4. Add ongkos_dasar field
ALTER TABLE kurir
	ADD COLUMN ongkos_dasar INT;

-- 5. Change kurir table to shipping
ALTER TABLE kurir
	RENAME TO shipping;

-- 6. Drop shipping table
DROP TABLE shipping;

-- 7. Add new realtion
--    a. 1-to-1 payment method to description
CREATE TABLE payment_description (
	payment_desc_id INT AUTO_INCREMENT,
	description TEXT,
	CONSTRAINT PK_PAYMENT_DETAIL PRIMARY KEY(payment_desc_id)
);

ALTER TABLE payment_methods
	ADD COLUMN payment_desc_id INT,
	ADD CONSTRAINT FK_DESCID FOREIGN KEY (payment_desc_id) REFERENCES payment_description(payment_desc_id);


--    b. 1-to-many users to address
CREATE TABLE address (
	address_id INT AUTO_INCREMENT,
	address VARCHAR(100),
	zip VARCHAR(8),
	CONSTRAINT PK_ADDRESS_ID PRIMARY KEY(address_id)
);

ALTER TABLE users
	MODIFY address INT,
	ADD CONSTRAINT FK_ADDRESSID FOREIGN KEY (address) REFERENCES address(address_id);


--    c. many-to-many user to payment_method
CREATE TABLE user_payment_method_detail(
	user_id INT,
	payment_id INT,
	CONSTRAINT FK_USER_2 FOREIGN KEY (user_id) REFERENCES users(user_id),
	CONSTRAINT FK_PAYMENT FOREIGN KEY (payment_id) REFERENCES payment_methods(payment_id)
);