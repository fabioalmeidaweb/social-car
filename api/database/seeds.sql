DELETE FROM user;
DELETE FROM sqlite_sequence WHERE name='user';
INSERT INTO user (name, nick, email, password)
values
("User 1", "user_1", "user1@user.com", "$2a$10$ZYRZjL9r99alXUl6VWv1aOYK.uSC8xz6/E.HgsZqJ0UXeTCoMLldy"),
("User 2", "user_2", "user2@user.com", "$2a$10$ZYRZjL9r99alXUl6VWv1aOYK.uSC8xz6/E.HgsZqJ0UXeTCoMLldy"),
("User 3", "user_3", "user3@user.com", "$2a$10$ZYRZjL9r99alXUl6VWv1aOYK.uSC8xz6/E.HgsZqJ0UXeTCoMLldy");

DELETE FROM follower;
DELETE FROM sqlite_sequence WHERE name='follower';
INSERT INTO follower (user_id, follower_id)
values
(1, 2),
(1, 3),
(3, 1);
