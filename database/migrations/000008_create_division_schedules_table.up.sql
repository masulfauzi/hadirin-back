CREATE TABLE division_schedules (
    id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    division_id      UUID NOT NULL,
    hari             VARCHAR(10) NOT NULL CHECK (hari IN ('senin','selasa','rabu','kamis','jumat','sabtu','minggu')),
    is_hari_kerja    BOOLEAN NOT NULL DEFAULT TRUE,
    jam_masuk        TIME NULL,
    jam_keluar       TIME NULL,
    toleransi_menit  INT NOT NULL DEFAULT 0,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_division_hari UNIQUE (division_id, hari),
    CONSTRAINT fk_ds_division FOREIGN KEY (division_id) REFERENCES divisions(id) ON DELETE CASCADE
);
