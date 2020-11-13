CREATE TABLE clients(
                        id BIGSERIAL PRIMARY  KEY,
                        login TEXT NOT NULL UNIQUE,
                        password TEXT NOT NULL,
                        full_name TEXT NOT NULL,
                        passport TEXT NOT NULL,
                        birthday DATE NOT NULL,
                        status TEXT NOT NULL DEFAULT  'INACTIVE' CHECK (status in ('INACTIVE', 'ACTIVE')),
                        created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE cards(
                      id BIGSERIAL PRIMARY KEY,
                      number TEXT NOT NULL,
                      type TEXT NOT NULL DEFAULT 'REAL' CHECK(type in ('REAL', 'VIRTUAL')),
                      balance BIGINT NOT NULL DEFAULT 0,
                      issuer TEXT NOT NULL CHECK (issuer in ('Visa', 'MasterCard', 'MIR')),
                      holder TEXT NOT NULL,
                      owner_id BIGINT NOT NULL REFERENCES clients (id),
                      status TEXT NOT NULL DEFAULT 'INACTIVE' check (status in ('INACTIVE', 'ACTIVE')),
                      created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE icons(
                      id BIGSERIAL PRIMARY KEY,
                      link TEXT NOT NULL
);
CREATE TABLE transactions(
                             id BIGSERIAL PRIMARY KEY,
                             mcc TEXT NOT NULL,
                             icon_id BIGINT REFERENCES icons (id),
                             status TEXT NOT NULL DEFAULT 'OK' check ( status in ('OK', 'FAIL') ),
                             amount BIGINT,
                             date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             card BIGINT REFERENCES cards (id)
);