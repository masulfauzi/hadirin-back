CREATE TABLE presensi (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    karyawan_id         UUID NOT NULL,
    tanggal             DATE NOT NULL,

    jam_hadir           TIMESTAMPTZ NULL,
    foto_hadir          VARCHAR(255) NULL,
    lat_hadir           DECIMAL(10,8) NULL,
    long_hadir          DECIMAL(11,8) NULL,
    jarak_hadir_meter   DECIMAL(10,2) NULL,
    status_hadir        VARCHAR(20) NULL CHECK (status_hadir IS NULL OR status_hadir IN ('tepat_waktu','terlambat')),

    jam_pulang          TIMESTAMPTZ NULL,
    foto_pulang         VARCHAR(255) NULL,
    lat_pulang          DECIMAL(10,8) NULL,
    long_pulang         DECIMAL(11,8) NULL,
    jarak_pulang_meter  DECIMAL(10,2) NULL,
    status_pulang       VARCHAR(20) NULL CHECK (status_pulang IS NULL OR status_pulang IN ('normal','pulang_cepat')),

    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_presensi_karyawan_tanggal UNIQUE (karyawan_id, tanggal),
    CONSTRAINT fk_presensi_karyawan FOREIGN KEY (karyawan_id) REFERENCES karyawan(id) ON DELETE CASCADE
);
