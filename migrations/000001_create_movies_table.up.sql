CREATE TABLE IF NOT EXISTS books (
         id bigserial PRIMARY KEY,
         bookname text NOT NULL,
         author text NOT NULL,
         releasedyear text not null
);
