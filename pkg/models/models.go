package models

import "fmt"

type User struct {
	ID          int
	Name        string
	DateCreated string
}

type Menu struct {
	UserID    int64
	MenuType  string
	ChatID    int64
	MessageID int
	Args      []string
}

type Menus struct {
	menus []*Menu
}

// Check if there is a menu struct with the user's id
func (menus *Menus) UserHasMenu(userId int64) bool {
	for _, menu := range menus.menus {
		if menu.UserID == userId {
			return true
		}
	}

	return false
}

// Append created menu struct
func (menus *Menus) Append(menu *Menu) {
	menus.menus = append(menus.menus, menu)
}

// Print all menu structs
func (menus *Menus) PrintAll() {
	for index, menu := range menus.menus {
		fmt.Printf("%d) %v\n", index, menu)
	}
}

// Remove by user's id
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

// Returns user's menu struct if found
func (menus *Menus) GetMenu(userId int64) *Menu {
	for _, menu := range menus.menus {
		if menu.UserID == userId {
			return menu
		}
	}

	return nil
}
