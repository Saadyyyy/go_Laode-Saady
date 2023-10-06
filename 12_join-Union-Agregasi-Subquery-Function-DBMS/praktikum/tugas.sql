CREATE DATABASE Alta_shop;

USE Alta_shop;

CREATE TABLE product_types(
    product_type_id INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE operators(
    operator_id INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products(
    product_id INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    product_type_id INT(11),
    operator_id INT(11),
    code VARCHAR(50),
    name VARCHAR(100),
    status SMALLINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_type_id) REFERENCES product_types(product_type_id),
    FOREIGN KEY (operator_id) REFERENCES operators(operator_id)
);

CREATE TABLE product_descriptions(
    product_desc_id INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE payment_methods(
    payment_method_id INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    status SMALLINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users(
    user_id INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    status SMALLINT,
    dob DATE,
    gender CHAR(1),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions(
    transaction_id INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    user_id INT(11),
    payment_method_id INT(11),
    status VARCHAR(10),
    total_qty INT(11),
    total_price NUMERIC(25,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(payment_method_id)
);

CREATE TABLE transaction_details(
    transaction_id INT(11) NOT NULL,
    product_id INT(11),
    status VARCHAR(10),
    qty INT(11),
    price NUMERIC(25,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (transaction_id, product_id),
    FOREIGN KEY (transaction_id) REFERENCES transactions(transaction_id),
    FOREIGN KEY (product_id) REFERENCES products(product_id)
);

-- 1. Insert
-- a. Insert 5 operators pada table operators.
INSERT INTO operators (name) VALUES
    ('Operator 1'),
    ('Operator 2'),
    ('Operator 3'),
    ('Operator 4'),
    ('Operator 5');

SELECT * FROM operators;

-- b. Insert 3 product type.
INSERT INTO product_types (name) VALUES
    ('Lock'),
    ('Tshirt'),
    ('Wings');

SELECT * FROM product_types;

-- c. Insert 2 product dengan product type id = 1, dan operators id = 3.
INSERT INTO products (product_type_id, operator_id, code, name, status) VALUES
    (1, 3, 'XA76', 'World Lock', 2),
    (1, 3, 'XA77', 'small lock', 3);

-- d. Insert 3 product dengan product type id = 2, dan operators id = 1.
INSERT INTO products (product_type_id, operator_id, code, name, status) VALUES
    (2, 1, 'BA09', 'Black Tshirt', 4),
    (2, 1, 'BA10', 'Pagani Tshirt', 5),
    (2, 1, 'BA11', 'Saady Tshirt', 6);

-- e. Insert 3 product dengan product type id = 3, dan operators id = 4.
INSERT INTO products (product_type_id, operator_id, code, name, status) VALUES
    (3, 4, 'CV06', 'Davinci wings', 7),
    (3, 4, 'CV07', 'Burung Wings', 8),
    (3, 4, 'CV08', 'cacing Wings', 9);

SELECT * FROM products;

-- f. Insert product description pada setiap product.
INSERT INTO product_descriptions (description) VALUES 
    ('World Lock - lock ini digunakan untuk mengunci pintu rumah'),
    ('small lock - small lock di gunakan untuk mengunci rumah dengan sekala lebih kecil'),
    ('Black Tshirt - baju dengan warna hitam yang awet sampai 50 tahun'),
    ('Pagani Tshirt - baju Pagani adalah baju yang di dapatkan saat membeli mobil pagani'),
    ('Saady Tshirt- Baju saady the most expensive tshirt yang ada di dunia'),
    ('Davinci wings - davinci wing adalah sebuah sayap yang terinspirasi dari lord davinci'),
    ('Burung Wings- burung wings adalah sayap burung asli '),
    ('cacing Wings - cacing wings adalah cacing yang memiliki sayap yang besar melebihi tubuhnya');

SELECT * FROM product_descriptions;

ALTER TABLE products ADD COLUMN descriptions INT(11);

ALTER TABLE products ADD CONSTRAINT fk_product_descriptions
FOREIGN KEY (descriptions) REFERENCES product_descriptions(product_desc_id);

UPDATE products SET descriptions = 1 WHERE product_id = 1;
UPDATE products SET descriptions = 2 WHERE product_id = 2;
UPDATE products SET descriptions = 3 WHERE product_id = 3;
UPDATE products SET descriptions = 4 WHERE product_id = 4;
UPDATE products SET descriptions = 5 WHERE product_id = 5;
UPDATE products SET descriptions = 6 WHERE product_id = 6;
UPDATE products SET descriptions = 7 WHERE product_id = 7;
UPDATE products SET descriptions = 8 WHERE product_id = 8;

SELECT * FROM products JOIN product_descriptions on (products.descriptions = product_descriptions.product_desc_id);

-- g. Insert 3 payment methods
INSERT INTO payment_methods (name, status) VALUES
    ('QRIS', 1),
    ('Transfer Bank / Virtual Account', 2),
    ('Paylater', 3);

SELECT * FROM payment_methods;

-- h. Insert 5 user pada tabel user.
ALTER TABLE users ADD COLUMN name VARCHAR(255) NOT NULL;

INSERT INTO users (status, name, dob, gender) VALUES
    ('1', 'Laode saady', '2002-03-10', 'F'),
    ('1', 'saady Laode', '2001-11-12', 'F'),
    ('1', 'Laode Laode ', '1998-12-07', 'M'),
    ('1', 'saady saady', '1985-09-10', 'M'),
    ('1', 'saady Perdana', '2001-11-05', 'M');

SELECT * FROM users;

-- i. Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)
INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES
    (1, 1, 'Success', 2, 1000000),
    (1, 2, 'Pending', 1, 500000),
    (1, 3, 'Failed', 3, 1500000),
    --==========================--
    (2, 1, 'Success', 1, 500000),
    (2, 2, 'Success', 4, 2000000),
    (2, 3, 'Pending', 2, 1000000),
    --==========================--
    (3, 1, 'Failed', 1, 500000),
    (3, 2, 'Success', 2, 1000000),
    (3, 3, 'Success', 3, 1500000),
    --==========================--
    (4, 1, 'Success', 3, 1500000),
    (4, 2, 'Failed', 2, 1000000),
    (4, 3, 'Success', 1, 500000),
    --==========================--
    (5, 1, 'Pending', 4, 2000000),
    (5, 2, 'Success', 1, 500000),
    (5, 3, 'Success', 2, 1000000);

SELECT * FROM transactions

-- j. Insert 3 product di masing-masing transaksi.
-- Transaksi 1
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price) VALUES
    (1, 1, 'In Progress', 1, 500000),
    (1, 2, 'Completed', 2, 2000000),
    (1, 3, 'In Progress', 1, 1000000);

-- Transaksi 2
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price) VALUES
    (2, 4, 'Completed', 1, 500000),
    (2, 5, 'In Progress', 3, 1500000),
    (2, 6, 'Completed', 2, 1000000);

-- Transaksi 3
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price) VALUES
    (3, 6, 'In Progress', 2, 1000000),
    (3, 7, 'Completed', 1, 500000),
    (3, 8, 'Completed', 3, 1500000);

-- Transaksi 4
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price) VALUES
    (4, 7, 'Completed', 3, 1500000),
    (4, 8, 'In Progress', 2, 1000000),
    (4, 1, 'In Progress', 1, 500000);

-- Transaksi 5
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price) VALUES
    (5, 2, 'In Progress', 4, 2000000),
    (5, 3, 'Completed', 1, 500000),
    (5, 4, 'Completed', 2, 1000000);

SELECT * FROM transaction_details

-- 2. Select 
-- a. Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
SELECT name FROM users WHERE gender = 'M';

-- b. Tampilkan product dengan id = 3.
SELECT * FROM products WHERE product_id = 3;

-- c. Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.
SELECT * FROM users WHERE created_at >= NOW() - INTERVAL 7 DAY AND name LIKE '%a%';

-- d. Hitung jumlah user / pelanggan dengan status gender Perempuan.
SELECT COUNT(*) AS jumlah_perempuan FROM users WHERE gender = 'F';

-- e. Tampilkan data pelanggan dengan urutan sesuai nama abjad
SELECT * FROM users ORDER BY name ASC; 

-- f. Tampilkan 5 data pada data product
SELECT * FROM products LIMIT 5;

-- 3. Update
-- a. Ubah data product id 1 dengan nama ‘product dummy’.
UPDATE products SET name = 'product dummy' WHERE product_id = 1;
SELECT * FROM products;

-- b. Update qty = 3 pada transaction detail dengan product id = 1.
UPDATE transaction_details SET qty = 3 WHERE product_id = 1;
SELECT * FRoM transaction_details

-- 4. Delete
-- a. Delete data pada tabel product dengan id = 1.
DELETE FROM transaction_details WHERE product_id = 1;
DELETE FROM products WHERE product_id = 1;
SELECT * FROM products;

-- b. Delete pada pada tabel product dengan product type id 1.
DELETE FROM transaction_details WHERE product_id IN (SELECT product_id FROM products WHERE product_type_id = 1);
DELETE FROM products WHERE product_type_id = 1;
SELECT * FROM products;

-- Join, Union, Sub query, Function
-- 1. Gabungkan data transaksi dari user id 1 dan user id 2.
SELECT * FROM transactions WHERE user_id = 1 UNION
SELECT * FROM transactions WHERE user_id = 2;

-- 2. Tampilkan jumlah harga transaksi user id 1.
SELECT SUM(total_price) FROM transactions WHERE user_id = 1;

-- 3. Tampilkan total transaksi dengan product type 2.
SELECT SUM(total_qty) FROM transactions WHERE user_id IN (
    SELECT user_id
    FROM products
    WHERE product_type_id = 2
);

-- 4. Tampilkan semua field table product dan field name table product type yang saling berhubungan.
SELECT products.*, product_types.name AS product_type_name FROM products
JOIN product_types ON products.product_type_id = product_types.product_type_id;

-- 5. Tampilkan semua field table transaction, field name table product dan field name table user.
SELECT transactions.*, products.name AS product_name, users.name AS user_name FROM transactions
JOIN transaction_details ON transactions.transaction_id = transaction_details.transaction_id
JOIN products ON transaction_details.product_id = products.product_id
JOIN users ON transactions.user_id = users.user_id;

-- 6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.
DELIMITER //
CREATE FUNCTION DeleteTransactionWithDetails(transactionId INT) RETURNS INT BEGIN
    DELETE FROM transaction_details WHERE transaction_id = transactionId;
    DELETE FROM transactions WHERE transaction_id = transactionId;
    RETURN 1;
END;
//
DELIMITER ;

-- 7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.
DELIMITER //
CREATE FUNCTION UpdateTotalQty(transactionId INT) RETURNS INT BEGIN
    DECLARE totalQty INT;
    SELECT SUM(qty) INTO totalQty FROM transaction_details WHERE transaction_id = transactionId;
    UPDATE transactions SET total_qty = totalQty WHERE transaction_id = transactionId;
    RETURN 1;
END;
//
DELIMITER ;

-- 8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query.
SELECT * FROM products WHERE product_id NOT IN (SELECT DISTINCT product_id FROM transaction_details);