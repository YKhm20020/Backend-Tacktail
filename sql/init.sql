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

INSERT INTO users (id, name, password) VALUES ('akojfdak', 'a1_name', 'a1_password');
