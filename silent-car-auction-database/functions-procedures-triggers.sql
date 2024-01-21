-- Select all FORDS
CREATE OR REPLACE FUNCTION available_fords() RETURNS TABLE (
    auto_id INT,
    location CHAR(4),
    type_of_auto VARCHAR(128),
    mileage INT,
    year DATE
  ) AS $$ BEGIN RETURN QUERY
SELECT i.auto_id,
  a.location,
  a.type_of_auto,
  a.mileage,
  a.year
FROM tb_auto a LEFT JOIN tb_auto_index i ON a.VIN = i.VIN
WHERE a.type_of_auto LIKE '%FORD%';
END;
$$ LANGUAGE plpgsql;
-- Select all CHEVEROLETS
CREATE OR REPLACE FUNCTION available_cheverolets() RETURNS TABLE (
    auto_id INT,
    location CHAR(4),
    type_of_auto VARCHAR(128),
    mileage INT,
    year DATE
  ) AS $$ BEGIN RETURN QUERY
SELECT i.auto_id,
  a.location,
  a.type_of_auto,
  a.mileage,
  a.year
FROM tb_auto a LEFT JOIN tb_auto_index i ON a.VIN = i.VIN
WHERE a.type_of_auto LIKE '%CHEVEROLET%';
END;
$$ LANGUAGE plpgsql;
-- Select all cars that matches the brand
CREATE OR REPLACE FUNCTION select_by_brand(brand VARCHAR(128)) RETURNS TABLE (
    auto_id INT,
    location CHAR(4),
    type_of_auto VARCHAR(128),
    mileage INT,
    year DATE
  ) AS $$ BEGIN RETURN QUERY
SELECT i.auto_id,
  a.location,
  a.type_of_auto,
  a.mileage,
  a.year
FROM tb_auto a LEFT JOIN tb_auto_index i ON a.VIN = i.VIN
WHERE a.type_of_auto LIKE '%' || brand || '%';
END;
$$ LANGUAGE plpgsql;
-- Return the highest bids of each auto
CREATE OR REPLACE FUNCTION max_bid() RETURNS TABLE (
    auto_id INT,
    location CHAR(4),
    year DATE,
    bid DECIMAL(16, 2)
  ) AS $$ BEGIN RETURN QUERY
SELECT DISTINCT t.auto_id,
  a.location,
  a.year,
  MAX(t.bid) OVER (PARTITION BY t.auto_id) AS maximum_bid
FROM tb_transaction t
  INNER JOIN tb_auto_index i ON t.auto_id = i.auto_id
  LEFT JOIN tb_auto a ON i.VIN = a.VIN;
END;
$$ LANGUAGE plpgsql;
-- Return the winners and losers of a specific auto
CREATE OR REPLACE FUNCTION winners_and_losers(p_auto_id INT) RETURNS TABLE (
    auto_id INT,
    last_name VARCHAR(30),
    bid DECIMAL(16, 2),
    maximum_bid DECIMAL(16, 2),
    is_winner BOOLEAN
  ) AS $$ BEGIN RETURN QUERY
SELECT t.auto_id,
  c.last_name,
  t.bid,
  MAX(t.bid) OVER (PARTITION BY t.auto_id) AS maximum_bid,
  (t.bid = MAX(t.bid) OVER (PARTITION BY t.auto_id)) AS is_winner
FROM tb_transaction t
  INNER JOIN tb_customer_index i ON t.customer_id = i.customer_id
  LEFT JOIN tb_customer c ON i.email = c.email
WHERE t.auto_id = p_auto_id;
END;
$$ LANGUAGE plpgsql;
-- Add a new customer by first name, last name, location and email, aka. sign up
CREATE OR REPLACE PROCEDURE add_customer(
    first_name VARCHAR(30),
    last_name VARCHAR(30),
    location CHAR(4),
    email VARCHAR(64)
  ) LANGUAGE SQL AS $$
INSERT INTO tb_customer (first_name, last_name, location, email)
VALUES (first_name, last_name, location, email);
INSERT INTO tb_customer_index (email)
VALUES (email);
$$;
-- Add a new auto by VIN, location, type of auto, mileage and year
CREATE OR REPLACE PROCEDURE add_auto(
    VIN CHAR(17),
    location CHAR(4),
    type_of_auto VARCHAR(128),
    mileage INT,
    year DATE
  ) LANGUAGE SQL AS $$
INSERT INTO tb_auto (VIN, location, type_of_auto, mileage, year)
VALUES (VIN, location, type_of_auto, mileage, year);
INSERT INTO tb_auto_index (VIN)
VALUES (VIN);
$$;
-- Add a new transaction by auto id, customer id, bid and bid date, aka. place a bid
CREATE OR REPLACE PROCEDURE add_transaction(
    p_auto_id INT,
    p_customer_id INT,
    p_bid DECIMAL(16, 2),
    p_bid_date DATE
  ) AS $$ BEGIN IF (
    SELECT MAX(t.bid_date)
    FROM tb_transaction t
    WHERE t.auto_id = p_auto_id
  ) < (CURRENT_DATE - INTERVAL '1 month') THEN RAISE EXCEPTION 'The bid is overdue.';
END IF;
IF (
  SELECT MAX(t.bid)
  FROM tb_transaction t
  WHERE t.auto_id = p_auto_id
) >= p_bid THEN RAISE EXCEPTION 'The bid is lower than the current highest bid.';
END IF;
INSERT INTO tb_transaction (auto_id, customer_id, bid, bid_date)
VALUES (p_auto_id, p_customer_id, p_bid, p_bid_date);
END;
$$ LANGUAGE plpgsql;
-- If an auction's last bid is older than 1 month, it is over
-- Used by the trigger to set all the ended auctions' winners
CREATE OR REPLACE FUNCTION set_winners() RETURNS TRIGGER LANGUAGE plpgsql AS $$ BEGIN
INSERT INTO tb_winner (auto_id, customer_id)
SELECT t.auto_id,
  t.customer_id
FROM tb_transaction t
  RIGHT JOIN (
    SELECT t.auto_id,
      MAX(t.bid) max_bid
    FROM tb_transaction t
    GROUP BY t.auto_id
  ) m ON t.auto_id = m.auto_id
  LEFT JOIN tb_winner w ON w.auto_id = m.auto_id
WHERE t.bid_date < CURRENT_DATE - INTERVAL '1 month'
  AND w.customer_id IS NULL
  AND t.bid = m.max_bid;
RETURN NEW;
END;
$$;
-- Check before inserting a new transaction in case the auction is over
CREATE TRIGGER set_winners BEFORE
INSERT ON tb_transaction FOR EACH ROW EXECUTE PROCEDURE set_winners();