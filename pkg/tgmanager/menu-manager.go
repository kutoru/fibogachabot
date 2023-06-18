package tgmanager

import (
	"fmt"
	"math"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kutoru/fibogachabot/pkg/dbmanager"
	"github.com/kutoru/fibogachabot/pkg/glb"
	"github.com/kutoru/fibogachabot/pkg/models"
)

func openMenu(msg tg.MessageConfig, menuType models.MenuType, listData *models.ListMenu, menuArgs ...string) {
	// fmt.Println("Before:")
	// glb.OpenedMenus.PrintAll()

	userId := msg.ChatID

	openedMenu := glb.OpenedMenus.GetMenu(userId)
	if openedMenu != nil {
		deleteMessage(openedMenu.UserID, openedMenu.MessageID)
	}

	sentMessage, err := glb.Bot.Send(msg)
	glb.CE(err)

	glb.OpenedMenus.Append(&models.Menu{
		UserID:    userId,
		MenuType:  menuType,
		MessageID: sentMessage.MessageID,
		ListData:  listData,
		Args:      menuArgs,
	})

	// fmt.Println("After:")
	// glb.OpenedMenus.PrintAll()
}

// Opens a special List menu. If toNext is false, it will open previous page, otherwise opens next page. isNew ignores toNext and opens the first page.
func openListMenu(msg tg.MessageConfig, menuType models.MenuType, isNew bool, toNext bool) {
	userId := msg.ChatID
	oldMenu := glb.OpenedMenus.GetMenu(userId)
	var listMenu *models.ListMenu

	if isNew || oldMenu == nil {
		listMenu = constructNewList(userId, menuType)
	} else {
		listMenu = oldMenu.ListData
		oldMenu.ListData = nil

		if toNext {
			listMenu.CurrentPageIndex += 1
		} else {
			listMenu.CurrentPageIndex -= 1
		}
	}

	msg.ParseMode = "HTML"
	msg.Text = listMenu.GetMenuText(menuType)
	msg.ReplyMarkup = listMenu.GetKeyboard(menuType)

	openMenu(msg, menuType, listMenu)
}

func constructNewList(userId int64, menuType models.MenuType) *models.ListMenu {
	var currentPageIndex int
	var maxPageIndex int
	var entriesPerPage int
	var entries []string

	switch menuType {

	case models.ProfileCharactersMenu:
		for index, acqChar := range dbmanager.GetAllAcqCharacters(userId, true) {
			entry := fmt.Sprintf("%d) %s", index+1, acqChar.CharacterInfo.Name)
			entries = append(entries, entry)
		}

		currentPageIndex = 0
		entriesPerPage = 10

	case models.ProfileInventoryMenu:
		for index, acqGift := range dbmanager.GetAllAcqGifts(userId, true) {
			var value string
			var spaces string = "   "

			switch acqGift.GiftInfo.Rarity {
			case 1:
				value = "Ok"
			case 2:
				value = "Good"
			case 3:
				value = "Great"
			}

			if index+1 >= 10 {
				spaces += " "
			}

			entry := fmt.Sprintf(
				"%d) Name: %s;\n%sType: %s;\n%sValue: %s;\n%sAmount: %d;",
				index+1, acqGift.GiftInfo.Name, spaces, acqGift.GiftInfo.Type, spaces, value, spaces, acqGift.Amount,
			)
			entries = append(entries, entry)
		}

		currentPageIndex = 0
		entriesPerPage = 5
	}

	if len(entries) <= entriesPerPage {
		maxPageIndex = 0
	} else {
		maxPageIndex = int(math.Ceil(float64(len(entries))/float64(entriesPerPage))) - 1
	}

	return &models.ListMenu{
		CurrentPageIndex: currentPageIndex,
		MaxPageIndex:     maxPageIndex,
		EntriesPerPage:   entriesPerPage,
		Entries:          entries,
	}
}
