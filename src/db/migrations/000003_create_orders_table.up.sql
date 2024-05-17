CREATE TABLE IF NOT EXISTS orders (
    id INT AUTO_INCREMENT,
    user_id INT NOT NULL,
    quantity INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
