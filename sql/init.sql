CREATE TABLE users (
  id VARCHAR(100) NOT NULL,
  name VARCHAR(20) NOT NULL,
  password VARCHAR(100) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE cocktails (
  id VARCHAR(100) NOT NULL,
  name VARCHAR(50) NOT NULL,
  description VARCHAR(100),
  PRIMARY KEY (id)
);

CREATE TABLE materials (
  id VARCHAR(100) NOT NULL,
  name VARCHAR(50) NOT NULL,
  description VARCHAR(100),
  PRIMARY KEY (id)
);

CREATE TABLE recipes (
  id VARCHAR(100) NOT NULL,
  cocktailID VARCHAR(100) NOT NULL,
  materialID VARCHAR(100) NOT NULL,
  amount INTEGER NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (cocktailID) REFERENCES cocktails(id)
  FOREIGN KEY (materialID) REFERENCES materials(id)
);

CREATE TABLE cocktail_images (
  id VARCHAR(100) NOT NULL,
  cocktailID VARCHAR(100) NOT NULL,
  userID VARCHAR(100) NOT NULL,
  description VARCHAR(100) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (cocktailID) REFERENCES cocktails(id)
  FOREIGN KEY (userID) REFERENCES users(id)
);

CREATE TABLE stories (
  id VARCHAR(100) NOT NULL,
  cocktailID VARCHAR(100) NOT NULL,
  trivia VARCHAR(100) NOT NULL,
  day INTEGER NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (cocktailID) REFERENCES cocktails(id)
);

INSERT INTO users (id, name, password) VALUES ('user00_id', 'user00_name', 'user00_password');
INSERT INTO users (id, name, password) VALUES ('user01_id', 'user01_name', 'user01_password');
INSERT INTO users (id, name, password) VALUES ('user02_id', 'user02_name', 'user02_password');
INSERT INTO users (id, name, password) VALUES ('user03_id', 'user03_name', 'user03_password');
INSERT INTO users (id, name, password) VALUES ('user04_id', 'user04_name', 'user04_password');
INSERT INTO users (id, name, password) VALUES ('user05_id', 'user05_name', 'user05_password');
INSERT INTO users (id, name, password) VALUES ('user06_id', 'user06_name', 'user06_password');
INSERT INTO users (id, name, password) VALUES ('user07_id', 'user07_name', 'user07_password');
INSERT INTO users (id, name, password) VALUES ('user08_id', 'user08_name', 'user08_password');
INSERT INTO users (id, name, password) VALUES ('user09_id', 'user09_name', 'user09_password');
INSERT INTO users (id, name, password) VALUES ('user10_id', 'user10_name', 'user10_password');
