insert into users (name, nick, email, senha)
values
("José", "zezis", "juserobertu@gmail.com", "$2a$10$UWUCQJpZlUocL5Cn8BMlS.WEM6.SGX68JXsvxFTn6oknYOolMaNmG"),
("Thiago", "thi", "thiago@gmail.com", "$2a$10$UWUCQJpZlUocL5Cn8BMlS.WEM6.SGX68JXsvxFTn6oknYOolMaNmG"),
("Samuel", "Samu", "samuel@gmail.com", "$2a$10$UWUCQJpZlUocL5Cn8BMlS.WEM6.SGX68JXsvxFTn6oknYOolMaNmG");

insert into followers(user_id, follower_id)
values
(1,2),
(2,1),
(1,3);