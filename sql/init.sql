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
  FOREIGN KEY (cocktailID) REFERENCES cocktails(id),
  FOREIGN KEY (materialID) REFERENCES materials(id)
);

CREATE TABLE cocktail_images (
  id VARCHAR(100) NOT NULL,
  cocktailID VARCHAR(100) NOT NULL,
  userID VARCHAR(100) NOT NULL,
  image VARCHAR(100) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (cocktailID) REFERENCES cocktails(id),
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

INSERT INTO cocktails (id, name, description) VALUES ('cocktail00_id', 'モスコミュール', 'モスコミュールの説明');
INSERT INTO cocktails (id, name, description) VALUES ('cocktail01_id', 'スクリュードライバー', 'スクリュードライバーの説明');
INSERT INTO cocktails (id, name, description) VALUES ('cocktail02_id', 'テキーラサンライズ', 'テキーラサンライズの説明');
INSERT INTO cocktails (id, name, description) VALUES ('cocktail03_id', 'ロングアイランドアイスティー', 'ロングアイランドアイスティーの説明');

INSERT INTO materials (id, name, description) VALUES ('material00_id', 'ウォッカ', 'ウォッカの説明');
INSERT INTO materials (id, name, description) VALUES ('material01_id', 'ジンジャーエール', 'ジンジャーエールの説明');
INSERT INTO materials (id, name, description) VALUES ('material02_id', 'オレンジジュース', 'オレンジジュースの説明');
INSERT INTO materials (id, name, description) VALUES ('material03_id', 'テキーラ', 'テキーラの説明');
INSERT INTO materials (id, name, description) VALUES ('material04_id', 'グレナンデン・シロップ', 'グレナンデン・シロップの説明');
INSERT INTO materials (id, name, description) VALUES ('material05_id', 'ラム', 'ラムの説明');
INSERT INTO materials (id, name, description) VALUES ('material06_id', 'ジン', 'ジンの説明');
INSERT INTO materials (id, name, description) VALUES ('material07_id', 'コアントロー', 'コアントローの説明');
INSERT INTO materials (id, name, description) VALUES ('material08_id', 'コーラ', 'コーラの説明');
INSERT INTO materials (id, name, description) VALUES ('material09_id', 'レモン果汁', 'レモン果汁の説明');

INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes00_id', 'cocktail00_id', 'material00_id', 2);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes01_id', 'cocktail00_id', 'material01_id', 8);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes02_id', 'cocktail01_id', 'material00_id', 2);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes03_id', 'cocktail01_id', 'material02_id', 8);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes04_id', 'cocktail02_id', 'material03_id', 3);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes05_id', 'cocktail02_id', 'material02_id', 6);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes06_id', 'cocktail02_id', 'material04_id', 1);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes07_id', 'cocktail03_id', 'material03_id', 1);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes08_id', 'cocktail03_id', 'material05_id', 1);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes09_id', 'cocktail03_id', 'material06_id', 1);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes10_id', 'cocktail03_id', 'material07_id', 1);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes11_id', 'cocktail03_id', 'material08_id', 4);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes12_id', 'cocktail03_id', 'material09_id', 1);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes13_id', 'cocktail03_id', 'material00_id', 1);

INSERT INTO cocktail_images (id, cocktailID, userID, image) VALUES ('cocktail_image00_id', 'cocktail00_id', 'user00_id', '/images/user00_モスコミュール画像.png');
INSERT INTO cocktail_images (id, cocktailID, userID, image) VALUES ('cocktail_image01_id', 'cocktail01_id', 'user00_id', '/images/user00_スクリュードライバー画像.png');
INSERT INTO cocktail_images (id, cocktailID, userID, image) VALUES ('cocktail_image02_id', 'cocktail00_id', 'user01_id', '/images/user01_モスコミュール画像.png');

INSERT INTO stories (id, cocktailID, trivia, day) VALUES ('story00_id', 'cocktail00_id', 'trivia00', 1);
INSERT INTO stories (id, cocktailID, trivia, day) VALUES ('story01_id', 'cocktail01_id', 'trivia01', 2);
INSERT INTO stories (id, cocktailID, trivia, day) VALUES ('story02_id', 'cocktail02_id', 'trivia02', 3);
INSERT INTO stories (id, cocktailID, trivia, day) VALUES ('story03_id', 'cocktail03_id', 'trivia03', 4);
