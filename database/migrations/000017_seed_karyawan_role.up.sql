INSERT INTO roles (kode_role, nama_role, deskripsi, is_active) VALUES
('KARYAWAN', 'Karyawan', 'Akses standar untuk karyawan', TRUE)
ON CONFLICT (kode_role) DO NOTHING;
