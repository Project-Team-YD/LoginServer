CREATE TABLE account (
    uid VARCHAR(45) NOT NULL PRIMARY KEY,
    user_id VARCHAR(45) NOT NULL,
    user_name VARCHAR(20) DEFAULT NULL,
);

CREATE TABLE items (
    id int NOT NULL AUTO_INCREMENT,
    item_name VARCHAR(40),
    item_type int NOT NULL,
    img_name VARCHAR(45),
    is_stack BOOLEAN NOT NULL DEFAULT false,
    PRIMARY KEY(id),
);

CREATE TABLE item_weapon (
    id int NOT NULL,
    item_name VARCHAR(40),
    low_dmg int,
    high_dmg int,
    atk_spd int,
    atk_range int,
    PRIMARY KEY (id),
    FOREIGN KEY (id) REFERENCES items(id) ON DELETE CASCADE,
    FOREIGN KEY (item_name) REFERENCES items(item_name) ON UPDATE CASCADE,
);

CREATE TABLE item_effect (
    id int NOT NULL,
    max_hp
    regen_hp int,
    short_dmg int,
    long_dmg int,
    atk_spd int,
    atk_range int,
    def int,
)

CREATE TABLE shop (
    id int NOT NULL,
    money_type int,
    price int,
    PRIMARY KEY (id),
    FOREIGN KEY (id) REFERENCES items(id) ON DELETE CASCADE,
);

