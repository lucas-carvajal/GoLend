DROP TABLE IF EXISTS users;
CREATE TABLE users (
    username       VARCHAR(256) NOT NULL,
    password       VARCHAR(256) NOT NULL,
    email          VARCHAR(256) NOT NULL,
    PRIMARY KEY (`username`),
    UNIQUE (username, email)
);

DROP TABLE IF EXISTS loans;
CREATE TABLE loans (
    id          INT AUTO_INCREMENT NOT NULL,
    fromUser    VARCHAR(256) NOT NULL,
    toUser      VARCHAR(256) NOT NULL,
    date        DATE NOT NULL,
    amount      INT NOT NULL,
    interest    INT,
    status      VARCHAR(256),
    PRIMARY KEY (`id`),
    UNIQUE (id),
    FOREIGN KEY (fromUser) REFERENCES users(username),
    FOREIGN KEY (toUser) REFERENCES users(username)
);

DROP TABLE IF EXISTS openBoard;
CREATE TABLE openBoard (
    openID      INT AUTO_INCREMENT NOT NULL,
    fromUser    VARCHAR(256) NOT NULL,
    toUser      VARCHAR(256) NOT NULL,
    loan        INT NOT NUlL,
    PRIMARY KEY (`openID`),
    FOREIGN KEY (fromUser) REFERENCES users(username),
    FOREIGN KEY (toUser) REFERENCES users(username),
    FOREIGN KEY (loan) REFERENCES loans(id)
);

DROP TABLE IF EXISTS closeBoard;
CREATE TABLE closeBoard (
    closeID     INT AUTO_INCREMENT NOT NULL,
    fromUser    VARCHAR(256) NOT NULL,
    toUser      VARCHAR(256) NOT NULL,
    loan        INT NOT NUlL,
    PRIMARY KEY (`closeID`),
    FOREIGN KEY (fromUser) REFERENCES users(username),
    FOREIGN KEY (toUser) REFERENCES users(username),
    FOREIGN KEY (loan) REFERENCES loans(id)
);



