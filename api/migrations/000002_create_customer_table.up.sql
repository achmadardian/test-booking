CREATE TABLE IF NOT EXISTS customer (
    cst_id SERIAL PRIMARY KEY NOT NULL,
    nationality_id INT REFERENCES nationality(nationality_id) NOT NULL,
    cst_name VARCHAR(50) NOT NULL,
    cst_dob DATE NOT NULL,
    cst_phone_num VARCHAR(20) NOT NULL,
    cst_email VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ
);

INSERT INTO customer (nationality_id, cst_name, cst_dob, cst_phone_num, cst_email)
VALUES
(1, 'Achmad Ardian', '2000-06-25', '08812632131', 'ardian@gmail.com'),
(2, 'John Doe', '2000-01-02', '0881712381', 'john@gmail.com'),
(1, 'Fulan', '2000-09-08', '08812731238', 'fulan@gmail.com')