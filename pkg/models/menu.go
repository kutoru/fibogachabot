package models

import "fmt"

// a way to create enums in go
type MenuType int

const (
	MainMenu MenuType = iota
	PlayMenu
	DreamMenu
	ShopMenu
	ProfileMenu
	SettingsMenu
	NotImplementedMenu
)

type Menu struct {
	UserID    int64
	MenuType  MenuType
	MessageID int
	Args      []string
}

type Menus struct {
	menus []*Menu
}

// Checks if there is a menu struct with the user's id
func (menus *Menus) UserHasMenu(userId int64) bool {
	for _, menu := range menus.menus {
		if menu.UserID == userId {
			return true
		}
	}

	return false
}

// Appends the menu struct
func (menus *Menus) Append(menu *Menu) {
	menus.menus = append(menus.menus, menu)
}

// Prints all menu structs
func (menus *Menus) PrintAll() {
	for index, menu := range menus.menus {
		fmt.Printf("%d) %v\n", index, menu)
	}
}

// Removes menu by user's id, does nothing if the id does not exist
func (menus *Menus) Remove(userId int64) {
	for index, menu := range menus.menus {
		if menu.UserID == userId {
			lastIndex := len(menus.menus) - 1
			menus.menus[index] = menus.menus[lastIndex]
			menus.menus = menus.menus[:lastIndex]
			return
		}
	}
}

// Returns user's menu struct if found, otherwise returns nil
func (menus *Menus) GetMenu(userId int64) *Menu {
	for _, menu := range menus.menus {
		if menu.UserID == userId {
			return menu
		}
	}

	return nil
}
