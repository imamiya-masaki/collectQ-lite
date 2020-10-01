CREATE TABLE posts (
    id int auto_increment not null primary key, 
    title VARCHAR(32), 
    content VARCHAR(256)
);

INSERT INTO collectq_db.posts (title, content) VALUES 
("gitむずかちい", "gitってなんであんなに難しいんですか?\n簡単にしてほしい"), ("docker is god", "dockerって神ですね^^\n簡単に初期化できるサイト知ってますか?"), ("好きなアニメを教えてください", "好きなアニメを教えてください。\n私は宇宙パトロールです");