CREATE TABLE divisions (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    kode_divisi  VARCHAR(50)  NOT NULL,
    nama_divisi  VARCHAR(100) NOT NULL,
    deskripsi    VARCHAR(255) NULL,
    is_active    BOOLEAN NOT NULL DEFAULT TRUE,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_divisions_kode UNIQUE (kode_divisi)
);
