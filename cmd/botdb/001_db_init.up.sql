/*
    Schema version: 001

    Database schema init
      - Create user table
      - Create role table
      - Create user to role association table
*/
CREATE TABLE IF NOT EXISTS bot_user(
  id SERIAL PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  ignored BOOLEAN NOT NULL DEFAULT FALSE,
  first_seen TIMESTAMP NOT NULL DEFAULT NOW(),
  last_seen TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS bot_user_role(
  id SERIAL PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  description VARCHAR(128),
  created_on TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_on TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO bot_user_role (name, description) VALUES ('admin', 'Admin user who can do anything');
INSERT INTO bot_user_role (name, description) VALUES ('nobody', 'Regular user who cannot do anything');

CREATE TABLE IF NOT EXISTS bot_user_to_role_association(
  bot_user_id INT REFERENCES bot_user(id) ON UPDATE CASCADE,
  bot_user_role_id INT REFERENCES bot_user_role(id) ON UPDATE CASCADE,
  created_on TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_on TIMESTAMP NOT NULL DEFAULT NOW()
);
