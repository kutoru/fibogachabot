create table users (
    id bigint not null,
    name varchar(255) not null,
    date_created datetime not null,
    illusions INT not null default 0,
    total_illusions int not null default 0,
    xcards INT not null default 0,
    total_xcards int not null default 0,
    gifts_bought int not null default 0,
    gifts_gifted int not null default 0,
    dailies_completed int not null default 0,
    notifications bool not null default 1,
    redeemed_codes json not null,
    primary key (id)
);

# dc_5 and dc_4 are dream counts that are responsible for guaranteed 5* and 4* drops respectively
# ge are responsible for guaranteed event drops
create table counters(
    user_id bigint not null,
    std_dc_5 INT not null default 0,
    std_dc_4 INT not null default 0,
    std_dc_total INT not null default 0,
    event_dc_5 INT not null default 0,
    event_dc_4 INT not null default 0,
    event_dc_total INT not null default 0,
    ge_5 bool not null default 0,
    ge_4 bool not null default 0,
    ge_3 bool not null default 0,
    foreign key (user_id) references users(id),
    primary key (user_id)
);

# S for std; E for event
create table dream_history(
    user_id bigint not null,
    banner_type enum('S', 'E') not null,
    history json not null,
    foreign key (user_id) references users(id),
    primary key (user_id, banner_type)
);

create table characters(
    id int not null,
    name varchar(255) NOT NULL,
    nickname varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    rarity int NOT NULL,
    primary key (id)
);

create table acquired_chars(
    user_id bigint NOT NULL,
    char_id int NOT NULL,
    friendship_exp INT not null default 0,
    friendship_lvl INT not null default 1,
    enigma INT not null default 0,
    completed_quests json not null,
    received_gifts json not null,
    date_acquired datetime not null,
    foreign key (user_id) references users(id),
    foreign key (char_id) references characters(id),
    PRIMARY KEY (user_id, char_id)
);

create table gifts(
    id int not null,
    name varchar(255) not null,
    type enum('Food', 'Tech', 'Music', 'Literature', 'Art', 'Toys') not null,
    rarity int not null,
    primary key (id)
);

create table acquired_gifts(
    user_id bigint not null,
    gift_id int not null,
    amount int not null,
    foreign key (user_id) references users(id),
    foreign key (gift_id) references gifts(id),
    primary key (user_id, gift_id)
);

create table achievements(
    id int auto_increment,
    title varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    primary key (id)
);

create table acquired_achievements(
    user_id bigint not null,
    achievement_id int not null,
    foreign key (user_id) references users(id),
    foreign key (achievement_id) references achievements(id),
    primary key (user_id, achievement_id)
);
