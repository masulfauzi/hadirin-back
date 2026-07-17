CREATE TABLE menus (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    parent_id   UUID NULL,
    kode_menu   VARCHAR(50)  NOT NULL,
    nama_menu   VARCHAR(100) NOT NULL,
    icon        VARCHAR(100) NULL,
    route       VARCHAR(255) NULL,
    urutan      INT NOT NULL DEFAULT 0,
    is_active   BOOLEAN NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_menus_kode UNIQUE (kode_menu),
    CONSTRAINT fk_menus_parent FOREIGN KEY (parent_id) REFERENCES menus(id) ON DELETE CASCADE
);
