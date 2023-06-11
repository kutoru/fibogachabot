drop table if exists users;
drop table if exists characters;
drop table if exists acquired_chars;
drop table if exists achievements;

create table users (
		id int not null,
		name varchar(255) not null,
		date_created datetime not null,
		illusions INT default 100,
		xcards INT default 0,
		notifications INT default 1,
		redeemed_codes json not null,
		primary key (id)
	);
    
create table characters(
        id int auto_increment,
        name TEXT NOT NULL,
        nickname TEXT NOT NULL,
        description TEXT NOT NULL,
        rarity INTEGER NOT NULL,
        original_skin_path TEXT NOT NULL,
		primary key (id)
    );
    
create table acquired_chars(
        user_id INT NOT NULL,
        char_name varchar(255) NOT NULL,
        friendship_exp INT default 0,
        friendship_lvl INT default 1,
        enigma INT default 0,
        completed_quests json not null,
        acquired_gifts json not null,
        PRIMARY KEY (user_id, char_name)
    );
    
create table achievements(
        id int,
        title TEXT NOT NULL,
        description TEXT NOT NULL,
		primary key (id)
    );