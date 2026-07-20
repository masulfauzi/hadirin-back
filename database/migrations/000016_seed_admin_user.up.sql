INSERT INTO users (kode_identitas, username, password_hash, is_active) VALUES
('ADMIN002', 'admin2', '$2a$10$UNSu2WuAIoxvmuH03f.tCudQRp1q5ztdaUTfaL.YV50yaLhHlHpbS', TRUE);

INSERT INTO roles (kode_role, nama_role, deskripsi, is_active) VALUES
('ADMIN', 'Administrator', 'Akses penuh ke seluruh fitur', TRUE)
ON CONFLICT (kode_role) DO NOTHING;

INSERT INTO user_roles (user_id, role_id)
SELECT u.id, r.id FROM users u, roles r
WHERE u.username = 'admin2' AND r.kode_role = 'ADMIN'
ON CONFLICT (user_id, role_id) DO NOTHING;

INSERT INTO menus (kode_menu, nama_menu, route, urutan, is_active) VALUES
('MENU_MANAGER', 'Menu Manager', '/menu-manager', 1, TRUE)
ON CONFLICT (kode_menu) DO NOTHING;

INSERT INTO role_menu_permissions (role_id, menu_id, can_show, can_read, can_insert, can_update, can_delete)
SELECT r.id, m.id, TRUE, TRUE, TRUE, TRUE, TRUE
FROM roles r, menus m
WHERE r.kode_role = 'ADMIN' AND m.kode_menu = 'MENU_MANAGER'
ON CONFLICT (role_id, menu_id) DO NOTHING;
