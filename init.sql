CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(50) NOT NULL
);

CREATE TABLE cars (
    id SERIAL PRIMARY KEY,
    model VARCHAR(50) NOT NULL,
    details TEXT NOT NULL
);

CREATE TABLE reservations (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    car_id INTEGER NOT NULL REFERENCES cars(id),
    details TEXT NOT NULL
);

INSERT INTO cars (model, details) VALUES
('Sedan', 'A nice sedan car'),
('SUV', 'A powerful SUV'),
('Coupe', 'A stylish coupe'),
('Convertible', 'A cool convertible'),
('Hatchback', 'A practical hatchback');
