CREATE TYPE relation AS ENUM ('father', 'mother', 'brother', 'sister');

CREATE TABLE IF NOT EXISTS family_list (
    fl_id SERIAL PRIMARY KEY NOT NULL,
    cst_id INT REFERENCES customer(cst_id) NOT NULL,
    fl_relation relation NOT NULL,
    fl_name VARCHAR(50) NOT NULL,
    fl_dob DATE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ
);

INSERT INTO family_list (cst_id, fl_relation, fl_name, fl_dob)
VALUES
('1', 'brother', 'Fauzi', '2008-01-02'),
('1', 'brother', 'Angga', '2003-03-04'),
('2', 'father', 'John', '1980-02-04'),
('2', 'mother', 'Rebecca', '1980-04-29'),
('3', 'father', 'Fulan', '1970-09-23'),
('3', 'sister', 'Fulanah', '2002-01-02');