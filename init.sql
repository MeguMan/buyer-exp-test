CREATE TABLE users (
    user_id bigserial PRIMARY KEY,
    email varchar NOT NULL
);

CREATE TABLE ads (
    ad_id bigserial PRIMARY KEY,
    link varchar NOT NULL,
    price int NOT NULL
);

CREATE TABLE users_ads (
    user_id    int REFERENCES users (user_id) ON UPDATE CASCADE ON DELETE CASCADE,
    ad_id int REFERENCES ads (ad_id) ON UPDATE CASCADE ON DELETE CASCADE
);