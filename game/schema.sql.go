package game

const schemasql string = `create table card (
	_id     int auto_increment primary key,
	title   text,
	deck    text,
	image   text,
	cat     text,
	subcat  text,
	cost    text,
	tcost   text,
	attack  text,
	defense text,
	attr    text,
	rarity  text,
	descr   text
);

create table profile (
	_id  int auto_increment primary key,
	name text
);

create table profile_card (
	profile_id int,
	card_id    int
);
create unique index idx_profile_card on profile_card (profile_id, card_id);`
