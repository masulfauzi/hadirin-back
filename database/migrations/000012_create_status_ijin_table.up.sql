CREATE TABLE status_ijin (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    kode_status   VARCHAR(20) NOT NULL,
    nama_status   VARCHAR(50) NOT NULL,
    warna_badge   VARCHAR(20) NULL,
    urutan        INT NOT NULL DEFAULT 0,
    CONSTRAINT uq_status_ijin_kode UNIQUE (kode_status)
);
