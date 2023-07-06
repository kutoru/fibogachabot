package models

import (
	"fmt"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// a way to create enums in go
type MenuType int

const (
	MainMenu MenuType = iota
	PlayMenu
	DreamMenu
	ShopMenu
	ProfileMenu
	SettingsMenu

	ProfileCharactersMenu
	ProfileInventoryMenu
	ProfileAchievementsMenu
	ProfileQuestsMenu

	NotImplementedMenu
)

type ListMenu struct {
	CurrentPageIndex int
	MaxPageIndex     int
	EntriesPerPage   int
	Entries          []string
}

type Menu struct {
	UserID    int64
	MenuType  MenuType
	MessageID int
	ListData  *ListMenu
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

func (listMenu *ListMenu) GetMenuText(menuType MenuType) string {
	if listMenu.CurrentPageIndex < 0 || listMenu.CurrentPageIndex > listMenu.MaxPageIndex {
		return "How did you get here?"
	}

	var menuTitle string

	switch menuType {

	case ProfileCharactersMenu:
		menuTitle = fmt.Sprintf(
			"Your characters (%d/%d)\nType character number to view their detailed info",
			listMenu.CurrentPageIndex+1, listMenu.MaxPageIndex+1,
		)

	case ProfileInventoryMenu:
		menuTitle = fmt.Sprintf(
			"Gifts you currently have (%d/%d)",
			listMenu.CurrentPageIndex+1, listMenu.MaxPageIndex+1,
		)

	}

	startIndex := listMenu.CurrentPageIndex * listMenu.EntriesPerPage
	endIndex := ((listMenu.CurrentPageIndex + 1) * listMenu.EntriesPerPage)
	if endIndex > len(listMenu.Entries) {
		endIndex = len(listMenu.Entries)
	}

	text := "<pre>"
	text += menuTitle
	for _, entry := range listMenu.Entries[startIndex:endIndex] {
		text += "\n" + entry
	}
	text += "</pre>"

	return text
}

func (listMenu *ListMenu) GetKeyboard(menuType MenuType) tg.InlineKeyboardMarkup {
	var callbackData string
	var backCallbackData string

	switch menuType {

	case ProfileCharactersMenu:
		callbackData = "profile_characters"
		backCallbackData = "profile"

	case ProfileInventoryMenu:
		callbackData = "profile_inventory"
		backCallbackData = "profile"

	}

	prevButton := tg.NewInlineKeyboardButtonData("⬅️ Previous page", callbackData+"|0")
	nextButton := tg.NewInlineKeyboardButtonData("➡️ Next page", callbackData+"|1")
	backRow := tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData("🔙 Go Back", backCallbackData))

	if listMenu.CurrentPageIndex < 0 || listMenu.CurrentPageIndex > listMenu.MaxPageIndex ||
		listMenu.CurrentPageIndex == 0 && listMenu.MaxPageIndex == 0 {
		return tg.NewInlineKeyboardMarkup(
			backRow,
		)
	} else if listMenu.CurrentPageIndex == 0 {
		return tg.NewInlineKeyboardMarkup(
			tg.NewInlineKeyboardRow(nextButton),
			backRow,
		)
	} else if listMenu.CurrentPageIndex == listMenu.MaxPageIndex {
		return tg.NewInlineKeyboardMarkup(
			tg.NewInlineKeyboardRow(prevButton),
			backRow,
		)
	} else {
		return tg.NewInlineKeyboardMarkup(
			tg.NewInlineKeyboardRow(prevButton, nextButton),
			backRow,
		)
	}
}
