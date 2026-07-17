CREATE TABLE karyawan (
    id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    kode_identitas    VARCHAR(50)  NOT NULL,
    nik               VARCHAR(20)  NULL,
    nama_lengkap      VARCHAR(150) NOT NULL,
    division_id       UUID NOT NULL,
    jabatan           VARCHAR(100) NULL,
    no_hp             VARCHAR(20)  NULL,
    email             VARCHAR(100) NULL,
    alamat            TEXT NULL,
    tanggal_masuk     DATE NULL,
    status_karyawan   VARCHAR(20) NOT NULL DEFAULT 'aktif' CHECK (status_karyawan IN ('aktif','nonaktif','resign')),
    foto_profile      VARCHAR(255) NULL,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_karyawan_kode_identitas UNIQUE (kode_identitas),
    CONSTRAINT uq_karyawan_nik UNIQUE (nik),
    CONSTRAINT fk_karyawan_division FOREIGN KEY (division_id) REFERENCES divisions(id)
);
