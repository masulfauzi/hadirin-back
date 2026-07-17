CREATE TABLE hari_libur (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    division_id  UUID NULL,
    tanggal      DATE NOT NULL,
    keterangan   VARCHAR(150) NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_hl_division FOREIGN KEY (division_id) REFERENCES divisions(id) ON DELETE CASCADE
);

CREATE INDEX idx_hari_libur_tanggal ON hari_libur(tanggal);
