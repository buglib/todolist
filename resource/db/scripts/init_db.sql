-- Active: 1668432948475@@127.0.0.1@3306
CREATE DATABASE IF NOT EXISTS todolist
    DEFAULT CHARACTER SET = 'utf8mb4';

CREATE DATABASE IF NOT EXISTS todolist_test
    DEFAULT CHARACTER SET = 'utf8mb4';

drop user 'buglib'@'%';
flush privileges;
create user 'buglib'@'%' identified by '123456';

grant all privileges on todolist.* to 'buglib'@'%';

grant all privileges on todolist_test.* to 'buglib'@'%';