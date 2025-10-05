CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE subs (
    id              uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id         uuid NOT NULL,

    service_name    varchar(100) NOT NULL,
    price           int8 CHECK (price BETWEEN 1 AND 100000),
    start_date      date NOT NULL,
    end_date        date
);

CREATE INDEX idx_id_pagination ON subs (user_id, id);
CREATE INDEX idx_svc_name_filter ON subs (user_id, service_name);
