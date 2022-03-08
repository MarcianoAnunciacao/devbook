insert into users(name, nick_name, email, password)
values
("Name", "NickName", "Email", "Password"),
("Name", "NickName", "Email", "Password"),
("Name", "NickName", "Email", "Password");

insert into followers(user_id, follower_id)
values
(1,2),
(3,2),
(1,3);

insert into publications(title, content, author_id)
values
("title", "content", 0),
("title", "content", 0),
("title", "content", 0)