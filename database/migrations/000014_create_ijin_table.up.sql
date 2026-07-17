CREATE TABLE ijin (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    karyawan_id         UUID NOT NULL,
    jenis_ijin_id       UUID NOT NULL,
    status_ijin_id      UUID NOT NULL,
    tanggal_mulai       DATE NOT NULL,
    tanggal_selesai     DATE NOT NULL,
    jumlah_hari         INT NOT NULL DEFAULT 1,
    alasan              TEXT NULL,
    file_lampiran       VARCHAR(255) NULL,
    disetujui_oleh      UUID NULL,
    tanggal_approval    TIMESTAMPTZ NULL,
    catatan_approval    VARCHAR(255) NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_ijin_karyawan FOREIGN KEY (karyawan_id) REFERENCES karyawan(id) ON DELETE CASCADE,
    CONSTRAINT fk_ijin_jenis FOREIGN KEY (jenis_ijin_id) REFERENCES jenis_ijin(id),
    CONSTRAINT fk_ijin_status FOREIGN KEY (status_ijin_id) REFERENCES status_ijin(id),
    CONSTRAINT fk_ijin_approver FOREIGN KEY (disetujui_oleh) REFERENCES users(id)
);

CREATE INDEX idx_ijin_karyawan_tanggal ON ijin(karyawan_id, tanggal_mulai);
