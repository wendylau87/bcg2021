CREATE DATABASE bcg;
USE bcg;


CREATE TABLE `item` (
                        `id` int NOT NULL AUTO_INCREMENT,
                        `name` varchar(255) NOT NULL,
                        `sku` varchar(255) NOT NULL,
                        `price` decimal(18,2) NOT NULL,
                        `quantity` int NOT NULL,
                        PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE `check_out` (
                             `id` int NOT NULL AUTO_INCREMENT,
                             `customer_name` varchar(255) NOT NULL,
                             `total_price`decimal(18,2) NOT NULL,
                             `created_at` datetime(6) NULL DEFAULT CURRENT_TIMESTAMP(6),
                             PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `check_out_detail` (
                                    `id` int NOT NULL AUTO_INCREMENT,
                                    `check_out_id` int NOT NULL,
                                    `item_id` int NOT NULL,
                                    `promotion_item_id` int NULL DEFAULT 0,
                                    `total_quantity` int NOT NULL,
                                    `sub_total_price` decimal(18,2) NOT NULL,
                                    PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `promotion` (
                             `id` int NOT NULL AUTO_INCREMENT,
                             `name` varchar(255) NOT NULL,
                             `minimum_item`int NULL DEFAULT 0,
                             `minimum_item_discount` decimal(18,2) NULL DEFAULT 0,
                             `count_origin` int NULL DEFAULT 0,
                             `count_destination` int NULL DEFAULT 0,
                             `free_item` bool NULL DEFAULT 0,
                             PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE `promotion_item` (
                                  `id` int NOT NULL AUTO_INCREMENT,
                                  `promotion_id` int NOT NULL,
                                  `item_id` int NOT NULL,
                                  `free_item_id` int NULL DEFAULT 0,
                                  `is_active` bool NULL DEFAULT 0,
                                  `priority` int NULL DEFAULT 1,
                                  PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO item(name, sku, price, quantity) values("Google Home", "120P90","49.99","10");
INSERT INTO item(name, sku, price, quantity) values("Macbook Pro", "43N23P","5399.99","5");
INSERT INTO item(name, sku, price, quantity) values("Alexa Speaker", "A304SD","109.50","10");
INSERT INTO item(name, sku, price, quantity) values("Raspberry Pi B", "234234","30.00","2");

INSERT INTO promotion(name, minimum_item, minimum_item_discount, count_origin, count_destination, free_item) values("FREE Item",0,0,0,0,1);
INSERT INTO promotion(name, minimum_item, minimum_item_discount, count_origin, count_destination, free_item) values("Buy 3 Pay 2",0,0,3,2,0);
INSERT INTO promotion(name, minimum_item, minimum_item_discount, count_origin, count_destination, free_item) values("Buy 3 Discount 10%",3,10,0,0,0);

INSERT INTO promotion_item(promotion_id, item_id, free_item_id, is_active) values(1,2,4,1);
INSERT INTO promotion_item(promotion_id, item_id, free_item_id, is_active) values(2,1,0,1);
INSERT INTO promotion_item(promotion_id, item_id, free_item_id, is_active) values(3,3,0,1);