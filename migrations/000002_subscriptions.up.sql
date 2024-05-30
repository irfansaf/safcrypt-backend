CREATE TABLE IF NOT EXISTS plans
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    price       INT          NOT NULL,
    duration    INT          NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS subscriptions
(
    id         SERIAL PRIMARY KEY,
    user_id    UUID REFERENCES users (id),
    plan_id    INT REFERENCES plans (id),
    start_date TIMESTAMP   NOT NULL,
    end_date   TIMESTAMP   NOT NULL,
    status     VARCHAR(50) NOT NULL,
    order_id   VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO plans (name, description, price, duration)
VALUES ('Basic', 'Basic plan description', 10000, 30),
       ('Premium', 'Premium plan description', 30000, 30),
       ('Annual', 'Annual plan description', 100000, 365);
