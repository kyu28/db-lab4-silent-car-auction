CREATE TABLE tb_auto (
  VIN CHAR(17) PRIMARY KEY,
  location CHAR(4) NOT NULL,
  type_of_auto VARCHAR(128) NOT NULL,
  mileage INT NOT NULL,
  year DATE NOT NULL
);

CREATE TABLE tb_customer (
  first_name VARCHAR(30) NOT NULL,
  last_name VARCHAR(30) NOT NULL,
  location CHAR(4) NOT NULL,
  email VARCHAR(64) PRIMARY KEY
);

CREATE TABLE tb_auto_index (
  auto_id SERIAL PRIMARY KEY,
  VIN CHAR(17) REFERENCES tb_auto(VIN) NOT NULL UNIQUE
);

CREATE TABLE tb_customer_index (
  customer_id SERIAL PRIMARY KEY,
  email VARCHAR(64) REFERENCES tb_customer(email) NOT NULL UNIQUE
);

CREATE TABLE tb_transaction (
  transaction_id SERIAL PRIMARY KEY,
  auto_id INT REFERENCES tb_auto_index(auto_id) NOT NULL,
  customer_id INT REFERENCES tb_customer_index(customer_id) NOT NULL,
  bid DECIMAL(16, 2) NOT NULL,
  bid_date DATE NOT NULL
);

CREATE TABLE tb_winner (
  auto_id INT REFERENCES tb_auto_index(auto_id) PRIMARY KEY,
  customer_id INT REFERENCES tb_customer_index(customer_id) NOT NULL
);