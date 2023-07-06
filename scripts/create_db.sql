CREATE TABLE users (
    id BIGINT NOT NULL,
    name VARCHAR(255) NOT NULL,
    date_created DATETIME NOT NULL,
    illusions INT NOT NULL DEFAULT 0,
    total_illusions INT NOT NULL DEFAULT 0,
    xcards INT NOT NULL DEFAULT 0,
    total_xcards INT NOT NULL DEFAULT 0,
    gifts_bought INT NOT NULL DEFAULT 0,
    gifts_gifted INT NOT NULL DEFAULT 0,
    dailies_completed INT NOT NULL DEFAULT 0,
    notifications BOOL NOT NULL DEFAULT 1,
    redeemed_codes JSON NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE dailies (
    user_id BIGINT NOT NULL,
    daily_index INT NOT NULL,
    type VARCHAR(255) NOT NULL,
    completed BOOL NOT NULL DEFAULT 0,
    CONSTRAINT max_daily_index CHECK (daily_index >= 0 AND daily_index <= 3),
    FOREIGN KEY (user_id) REFERENCES users(id),
    PRIMARY KEY (user_id, daily_index)
);

# dc_5 and dc_4 are dream counts that are responsible for guaranteed 5* and 4* drops respectively
# ge are responsible for guaranteed event drops
CREATE TABLE dream_counters (
    user_id BIGINT NOT NULL,
    std_dc_5 INT NOT NULL DEFAULT 0,
    std_dc_4 INT NOT NULL DEFAULT 0,
    std_dc_total INT NOT NULL DEFAULT 0,
    event_dc_5 INT NOT NULL DEFAULT 0,
    event_dc_4 INT NOT NULL DEFAULT 0,
    event_dc_total INT NOT NULL DEFAULT 0,
    ge_5 BOOL NOT NULL DEFAULT 0,
    ge_4 BOOL NOT NULL DEFAULT 0,
    ge_3 BOOL NOT NULL DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES users(id),
    PRIMARY KEY (user_id)
);

# S for std; E for event
CREATE TABLE dream_history (
    user_id bigint not null,
    banner_type enum('S', 'E') not null,
    history json not null,
    foreign key (user_id) references users(id),
    primary key (user_id, banner_type)
);

CREATE TABLE characters (
    id int not null,
    name varchar(255) NOT NULL,
    nickname varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    rarity int NOT NULL,
    CONSTRAINT max_rarity CHECK (rarity >= 3 AND rarity <= 5),
    primary key (id)
);

CREATE TABLE acquired_chars (
    user_id bigint NOT NULL,
    char_id int NOT NULL,
    friendship_exp INT not null default 0,
    friendship_lvl INT not null default 1,
    enigma INT not null default 0,
    received_gifts json not null,
    date_acquired datetime not null,
    foreign key (user_id) references users(id),
    foreign key (char_id) references characters(id),
    PRIMARY KEY (user_id, char_id)
);

CREATE TABLE gifts (
    id int not null,
    name varchar(255) not null,
    type enum('Food', 'Tech', 'Music', 'Literature', 'Art', 'Toys') not null,
    rarity int not null,
    price INT NOT NULL,
    CONSTRAINT max_rarity CHECK (rarity >= 1 AND rarity <= 3),
    primary key (id)
);

CREATE TABLE acquired_gifts (
    user_id bigint not null,
    gift_id int not null,
    amount int not null,
    foreign key (user_id) references users(id),
    foreign key (gift_id) references gifts(id),
    primary key (user_id, gift_id)
);

CREATE TABLE achievements (
    id int auto_increment,
    title varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    primary key (id)
);

CREATE TABLE acquired_achievements (
    user_id bigint not null,
    achievement_id int not null,
    date_acquired DATETIME NOT NULL,
    foreign key (user_id) references users(id),
    foreign key (achievement_id) references achievements(id),
    primary key (user_id, achievement_id)
);

# In theory, the constraint makes sure that if the quest type is Character, the char_id must be specified
CREATE TABLE quests (
    id INT NOT NULL,
    char_id INT DEFAULT NULL,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    type ENUM('Main', 'Character', 'Side') NOT NULL,
    requirements JSON NOT NULL,
    rewards JSON NOT NULL,
    CONSTRAINT null_char_id CHECK ((NOT type = 'Character') AND char_id IS NULL),
    CONSTRAINT not_null_char_id CHECK (type = 'Character' AND char_id IS NOT NULL),
    FOREIGN KEY (char_id) REFERENCES characters(id),
    PRIMARY KEY (id)
);

CREATE TABLE acquired_quests (
    user_id BIGINT NOT NULL,
    quest_id INT NOT NULL,
    completed BOOL NOT NULL DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (quest_id) REFERENCES quests(id),
    PRIMARY KEY (user_id, quest_id)
);
