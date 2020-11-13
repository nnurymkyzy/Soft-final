INSERT INTO clients(id, login, password, full_name, birthday, status, passport) VALUES
(1, 'IvanNov', 'hashedText', 'Ivan Ivanov', '1982-03-14', 'ACTIVE', '1337 322'),
(2, 'Petrov', 'hashedText', 'Petr Petrov', '19820-11-13', 'ACTIVE', '9482 1322');

INSERT INTO icons(id, link) VALUES (1, 'https://i.pinimg.com/736x/37/71/b1/3771b1eb84ccf8eaa63365bc0649e9ab.jpg');

INSERT INTO cards(id, number, balance, issuer, holder, owner_id, status, created) VALUES
(1, '*** 001', 0, 'Visa', 'Ivan Ivanov', 1, 'ACTIVE', now()),
(2, '*** 002', 0, 'MIR', 'Petr Petrov', 2, 'ACTIVE', now()),
(3, '*** 003', 4700000, 'MasterCard', 'Petr Petrov', 2, 'ACTIVE', now());

INSERT INTO transactions(id, mcc, icon_id, amount, status, card) VALUES
(1, '5556', 1, 5000000, 'OK', 3),
(2, '5226', 1, -100000, 'OK', 3),
(3, '1234', 1, -100000, 'OK', 3),
(8, '1234', 1, -100000, 'OK', 3),
(5, '1234', 1, -100000, 'OK', 3),
(6, '1234', 1, -100000, 'OK', 3),
(7, '1234', 1, -100000, 'OK', 3),
(4, '5526', 1, -100000, 'OK', 3);