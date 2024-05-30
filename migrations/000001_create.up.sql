CREATE TABLE IF NOT EXISTS roles
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO roles (name)
VALUES
    ('admin'),
    ('member');

CREATE TABLE IF NOT EXISTS users
(
    id         UUID PRIMARY KEY,
    first_name VARCHAR(50)  NOT NULL,
    last_name  VARCHAR(50)  NOT NULL,
    username   VARCHAR(50)  NOT NULL UNIQUE,
    email      VARCHAR(100) NOT NULL UNIQUE,
    phone      VARCHAR(20) NOT NULL UNIQUE,
    password   VARCHAR(100) NOT NULL,
    role_id    INT REFERENCES roles (id) DEFAULT 2,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (id, first_name, last_name, username, email, phone, password, role_id)
VALUES
    ('2489270f-9cb1-4e15-9a4c-bd0443382309', 'Irfan', 'Saf', 'irfansaf', 'irfansaf7@gmail.com', '085155024747', '$2a$10$6iWov3pW4vsCbgHdWE4cWeY.igg/6Xa/SEnVfWMCZDlT1BiMO5CXa', 1);
