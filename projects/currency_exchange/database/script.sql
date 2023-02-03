
CREATE DATABASE   IF NOT EXISTS ExchangeRateDb;

Use ExchangeRateDb;

DROP TABLE IF EXISTS exchange_rates;

CREATE TABLE
exchange_rates (
  id INT AUTO_INCREMENT NOT NULL,
  time DATE NOT NULL,
  currency VARCHAR(128) NOT NULL,
  rate FLOAT NOT NULL,
  PRIMARY KEY (`id`)
);
