package menu

import (
	"errors"

	"github.com/google/uuid"

	"hadirin-back/modules/role"
)

type Service struct {
	repo        *Repository
	roleService *role.Service
}

func NewService(repo *Repository, roleService *role.Service) *Service {
	return &Service{repo: repo, roleService: roleService}
}

func (s *Service) GetAll() ([]Menu, error) {
	return s.repo.FindAll()
}

func (s *Service) GetByID(id uuid.UUID) (*Menu, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Create(m *Menu) error {
	if m.KodeMenu == "" || m.NamaMenu == "" {
		return errors.New("kode_menu dan nama_menu wajib diisi")
	}
	return s.repo.Create(m)
}

func (s *Service) Update(m *Menu) error {
	if m.KodeMenu == "" || m.NamaMenu == "" {
		return errors.New("kode_menu dan nama_menu wajib diisi")
	}
	return s.repo.Update(m)
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *Service) GetPermissions(menuID uuid.UUID) ([]RoleMenuPermission, error) {
	if _, err := s.repo.FindByID(menuID); err != nil {
		return nil, err
	}
	return s.repo.FindPermissionsByMenuID(menuID)
}

func (s *Service) SetPermission(menuID uuid.UUID, p *RoleMenuPermission) error {
	if _, err := s.repo.FindByID(menuID); err != nil {
		return errors.New("menu tidak ditemukan")
	}
	if _, err := s.roleService.GetByID(p.RoleID); err != nil {
		return errors.New("role_id tidak ditemukan")
	}
	p.MenuID = menuID
	return s.repo.UpsertPermission(p)
}

// MenuNode adalah bentuk menu untuk sidebar frontend: sudah disaring
// berdasarkan permission role user yang login (can_show) dan disusun
// bertingkat lewat ParentID/Children.
type MenuNode struct {
	ID        uuid.UUID   `json:"id"`
	ParentID  *uuid.UUID  `json:"parent_id"`
	KodeMenu  string      `json:"kode_menu"`
	NamaMenu  string      `json:"nama_menu"`
	Icon      *string     `json:"icon"`
	Route     *string     `json:"route"`
	Urutan    int         `json:"urutan"`
	CanRead   bool        `json:"can_read"`
	CanInsert bool        `json:"can_insert"`
	CanUpdate bool        `json:"can_update"`
	CanDelete bool        `json:"can_delete"`
	Children  []*MenuNode `json:"children"`
}

func (s *Service) GetMenuForUser(userID uuid.UUID) ([]*MenuNode, error) {
	roleIDs, err := s.roleService.GetRoleIDsByUserID(userID)
	if err != nil {
		return nil, err
	}

	rows, err := s.repo.FindVisibleForRoles(roleIDs)
	if err != nil {
		return nil, err
	}

	nodes := make(map[uuid.UUID]*MenuNode, len(rows))
	for _, row := range rows {
		nodes[row.ID] = &MenuNode{
			ID:        row.ID,
			ParentID:  row.ParentID,
			KodeMenu:  row.KodeMenu,
			NamaMenu:  row.NamaMenu,
			Icon:      row.Icon,
			Route:     row.Route,
			Urutan:    row.Urutan,
			CanRead:   row.CanRead,
			CanInsert: row.CanInsert,
			CanUpdate: row.CanUpdate,
			CanDelete: row.CanDelete,
			Children:  []*MenuNode{},
		}
	}

	tree := make([]*MenuNode, 0)
	for _, row := range rows {
		node := nodes[row.ID]
		// Kalau parent-nya tidak ikut tampil (mis. permission parent belum
		// diset), naikkan menu ini jadi level atas supaya tetap terlihat.
		if row.ParentID != nil {
			if parent, ok := nodes[*row.ParentID]; ok {
				parent.Children = append(parent.Children, node)
				continue
			}
		}
		tree = append(tree, node)
	}
	return tree, nil
}
