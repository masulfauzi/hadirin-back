CREATE TABLE jenis_ijin (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    kode_jenis      VARCHAR(20)  NOT NULL,
    nama_jenis      VARCHAR(100) NOT NULL,
    max_hari        INT NULL,
    perlu_lampiran  BOOLEAN NOT NULL DEFAULT FALSE,
    CONSTRAINT uq_jenis_ijin_kode UNIQUE (kode_jenis)
);
