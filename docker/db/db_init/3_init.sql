CREATE TABLE quetions (
    id int auto_increment not null primary key, 
    title VARCHAR(32), 
    content VARCHAR(256),
    created_at TIMESTAMP,
    modifyed_at INT,
    group_id VARCHAR (256),
    privated BOOLEAN,
    create_user_id int
);

CREATE TABLE quetion_tags (
    id int auto_increment not null primary key,
    target_id int,
    tag_id int
);

CREATE TABLE tags (
    id int auto_increment not null primary key,
    tag_name VARCHAR (256)
);

CREATE TABLE user_ids (
    id int auto_increment not null primary key,
    group_id int,
    name_info VARCHAR (256)
);

CREATE TABLE group_ids (
    id int auto_increment not null primary key,
    name_info  VARCHAR (256)
);