DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS loans;
DROP TABLE IF EXISTS openBoard;
DROP TABLE IF EXISTS closeBoard;


CREATE TABLE users (
    username       VARCHAR(256) NOT NULL,
    password       VARCHAR(256) NOT NULL,
    email          VARCHAR(256) NOT NULL,
    PRIMARY KEY (`username`),
    UNIQUE (username, email)
);

CREATE TABLE loans (
    id          INT AUTO_INCREMENT NOT NULL,
    title       VARCHAR(256) NOT NULL,
    description VARCHAR(256),
    fromUser    VARCHAR(256) NOT NULL,
    toUser      VARCHAR(256) NOT NULL,
    date        DATE NOT NULL,
    amount      INT NOT NULL,
    interest    INT,
    deadline    DATE,
    status      VARCHAR(256),
    PRIMARY KEY (`id`),
    UNIQUE (id),
    FOREIGN KEY (fromUser) REFERENCES users(username),
    FOREIGN KEY (toUser) REFERENCES users(username)
);

-- test populate database

INSERT INTO users
    (username, password, email)
VALUES
    ('fed', 'fedmaster64', 'fed@gov.com'),
    ('pomp', 'btcftw', 'pomp@motorboys.com'),
    ('michael', 'ALLIN81', 'saylor@microstrategy.com');


