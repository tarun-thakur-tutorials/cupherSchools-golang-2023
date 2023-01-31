DROP TABLE IF EXISTS user;

CREATE TABLE
  user (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(128) NOT NULL,
    unique_phone_number VARCHAR(128) NOT NULL,
    credit_limit FLOAT (10, 2) NOT NULL,
    dues FLOAT (10, 2),
    PRIMARY KEY (`id`)
  );

DROP TABLE IF EXISTS userCreds;

CREATE TABLE
  userCreds (
    id INT NOT NULL,
    pass VARCHAR(128) NOT NULL,
    PRIMARY KEY (`id`)
  );

DROP TABLE IF EXISTS merchant;

CREATE TABLE
  merchant (
    id INT AUTO_INCREMENT NOT NULL,
    unique_name VARCHAR(128) NOT NULL,
    discount_offered DECIMAL(5, 2) NOT NULL,
    total_discount DECIMAL(10, 2),
    PRIMARY KEY (`id`)
  );

DROP TABLE IF EXISTS merchantCreds;

CREATE TABLE
  merchantCreds (
    id INT NOT NULL,
    pass VARCHAR(128) NOT NULL,
    PRIMARY KEY (`id`)
  );

DROP TABLE IF EXISTS transactions;

CREATE TABLE
  transactions (
    id INT NOT NULL,
    user_id INT NOT NULL,
    merchant_id INT NOT NULL,
    PRIMARY KEY (`id`)
  );
