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

INSERT INTO cocktails (id, name, description) VALUES ('cocktail00_id', 'cocktail00_name', 'cocktail00_description');
INSERT INTO cocktails (id, name, description) VALUES ('cocktail01_id', 'cocktail01_name', 'cocktail01_description');
INSERT INTO cocktails (id, name, description) VALUES ('cocktail02_id', 'cocktail02_name', 'cocktail02_description');
INSERT INTO cocktails (id, name, description) VALUES ('cocktail03_id', 'cocktail03_name', 'cocktail03_description');
INSERT INTO cocktails (id, name, description) VALUES ('cocktail04_id', 'cocktail04_name', 'cocktail04_description');
INSERT INTO cocktails (id, name, description) VALUES ('cocktail05_id', 'cocktail05_name', 'cocktail05_description');
INSERT INTO cocktails (id, name, description) VALUES ('cocktail06_id', 'cocktail06_name', 'cocktail06_description');
INSERT INTO cocktails (id, name, description) VALUES ('cocktail07_id', 'cocktail07_name', 'cocktail07_description');
INSERT INTO cocktails (id, name, description) VALUES ('cocktail08_id', 'cocktail08_name', 'cocktail08_description');
INSERT INTO cocktails (id, name, description) VALUES ('cocktail09_id', 'cocktail09_name', 'cocktail09_description');

INSERT INTO materials (id, name, description) VALUES ('material00_id', 'material00_name', 'material00_description');
INSERT INTO materials (id, name, description) VALUES ('material01_id', 'material01_name', 'material01_description');
INSERT INTO materials (id, name, description) VALUES ('material02_id', 'material02_name', 'material02_description');
INSERT INTO materials (id, name, description) VALUES ('material03_id', 'material03_name', 'material03_description');
INSERT INTO materials (id, name, description) VALUES ('material04_id', 'material04_name', 'material04_description');
INSERT INTO materials (id, name, description) VALUES ('material05_id', 'material05_name', 'material05_description');
INSERT INTO materials (id, name, description) VALUES ('material06_id', 'material06_name', 'material06_description');
INSERT INTO materials (id, name, description) VALUES ('material07_id', 'material07_name', 'material07_description');
INSERT INTO materials (id, name, description) VALUES ('material08_id', 'material08_name', 'material08_description');
INSERT INTO materials (id, name, description) VALUES ('material09_id', 'material09_name', 'material09_description');

INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes00_id', 'cocktail00_id', 'material00_id', 2);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes01_id', 'cocktail00_id', 'material01_id', 8);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes02_id', 'cocktail01_id', 'material02_id', 3);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes03_id', 'cocktail01_id', 'material01_id', 7);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes04_id', 'cocktail02_id', 'material03_id', 1);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes05_id', 'cocktail02_id', 'material04_id', 1);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes06_id', 'cocktail02_id', 'material00_id', 2);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes07_id', 'cocktail02_id', 'material05_id', 6);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes08_id', 'cocktail03_id', 'material00_id', 1);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes09_id', 'cocktail03_id', 'material02_id', 1);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes01_id', 'cocktail03_id', 'material06_id', 1);
INSERT INTO recipes (id, cocktailID, materialID, amount) VALUES ('recipes11_id', 'cocktail03_id', 'material07_id', 7);

INSERT INTO cocktail_images (id, cocktailID, userID, image) VALUES ('cocktail_image00_id', 'cocktail00_id', 'user00_id', "/images/image00_cocktail00_user00");
INSERT INTO cocktail_images (id, cocktailID, userID, image) VALUES ('cocktail_image01_id', 'cocktail01_id', 'user00_id', "/images/image01_cocktail01_user00");
INSERT INTO cocktail_images (id, cocktailID, userID, image) VALUES ('cocktail_image02_id', 'cocktail00_id', 'user01_id', "/images/image02_cocktail00_user01");
