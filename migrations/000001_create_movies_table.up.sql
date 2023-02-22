CREATE TABLE IF NOT EXISTS books (
         id bigserial PRIMARY KEY,
         bookName text NOT NULL,
         author text NOT NULL,
         releasedYear int,
         onlineBookName text,
         onlineBookFile bytea
);
