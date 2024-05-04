CREATE TABLE IF NOT EXISTS transacoes (
    id VARCHAR(32) primary key,
    valor INTEGER NOT NULL,
    tipo CHAR(1) NOT NULL CHECK (tipo IN ('c', 'd')),
    descricao TEXT NOT NULL,
    realizada_em TEXT NOT NULL
);
