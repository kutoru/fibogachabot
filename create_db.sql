drop table if exists acquired_achievements;
drop table if exists achievements;
drop table if exists acquired_chars;
drop table if exists characters;
drop table if exists dream_history;
drop table if exists counters;
drop table if exists users;

create table users (
    id int not null,
    name varchar(255) not null,
    date_created datetime not null,
    illusions INT default 0,
    xcards INT default 0,
    notifications bool default 1,
    redeemed_codes json not null,
    primary key (id)
);

# dc are dream counts that are responsible for guaranteed rarity drops
# ge are responsible for guaranteed event drops
create table counters(
    user_id int not null,
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
    user_id int not null,
    banner_type enum('S', 'E') not null,
    history json not null,
    foreign key (user_id) references users(id),
    primary key (user_id, banner_type)
);

create table characters(
    id int auto_increment,
    name varchar(255) NOT NULL,
    nickname varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    rarity int NOT NULL,
    original_skin_path varchar(255) NOT NULL,
    primary key (id)
);

create table acquired_chars(
    user_id INT NOT NULL,
    char_id int NOT NULL,
    friendship_exp INT not null default 0,
    friendship_lvl INT not null default 1,
    enigma INT not null default 0,
    completed_quests json not null,
    acquired_gifts json not null,
    foreign key (user_id) references users(id),
    foreign key (char_id) references characters(id),
    PRIMARY KEY (user_id, char_id)
);

create table achievements(
    id int auto_increment,
    title varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    primary key (id)
);

create table acquired_achievements(
    user_id int not null,
    achievement_id int not null,
    foreign key (user_id) references users(id),
    foreign key (achievement_id) references achievements(id),
    primary key (user_id, achievement_id)
)
