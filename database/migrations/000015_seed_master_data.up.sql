INSERT INTO status_ijin (kode_status, nama_status, warna_badge, urutan) VALUES
('MENUNGGU', 'Menunggu Persetujuan', 'yellow', 1),
('DISETUJUI', 'Disetujui', 'green', 2),
('DITOLAK', 'Ditolak', 'red', 3),
('DIBATALKAN', 'Dibatalkan', 'gray', 4);

INSERT INTO jenis_ijin (kode_jenis, nama_jenis, max_hari, perlu_lampiran) VALUES
('SAKIT', 'Sakit', NULL, TRUE),
('CUTI', 'Cuti Tahunan', 12, FALSE),
('IZIN', 'Izin Pribadi', NULL, FALSE),
('DUKA', 'Duka Cita', 3, FALSE);
