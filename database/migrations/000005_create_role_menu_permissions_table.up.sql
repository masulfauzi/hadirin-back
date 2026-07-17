CREATE TABLE role_menu_permissions (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    role_id     UUID NOT NULL,
    menu_id     UUID NOT NULL,
    can_show    BOOLEAN NOT NULL DEFAULT FALSE,
    can_read    BOOLEAN NOT NULL DEFAULT FALSE,
    can_insert  BOOLEAN NOT NULL DEFAULT FALSE,
    can_update  BOOLEAN NOT NULL DEFAULT FALSE,
    can_delete  BOOLEAN NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_role_menu UNIQUE (role_id, menu_id),
    CONSTRAINT fk_rmp_role FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    CONSTRAINT fk_rmp_menu FOREIGN KEY (menu_id) REFERENCES menus(id) ON DELETE CASCADE
);
