package main

const (
	MAIN_MENU_TYPE = "type"
)

type Menu struct {
	Title              string
	parent, prev, next *Menu
	Items              []Menu
	IsRoot             bool
	IsSubMenu          bool
	Slug               string
}

func (menu *Menu) Next(title string) *Menu {
	if menu == nil {
		return menu
	}
	if menu.next == nil {
		for _, item := range menu.Items {
			if item.Title == title {
				return item.next
			}
		}
	}
	return menu.next
}

func (menu *Menu) Prev() *Menu {
	if menu.IsRoot {
		return menu
	}
	if menu.prev == nil {
		return menu.parent.prev
	}
	if menu.prev.IsSubMenu {
		return menu.prev.parent
	}
	return menu.prev
}

func (menu *Menu) Insert(e *Menu) *Menu {
	next := menu.next
	menu.next = e
	e.prev = menu
	e.next = next
	if next != nil {
		next.prev = e
	}
	return e
}

func (menu *Menu) InsertItem(item *Menu) {
	item.parent = menu
	menu.Items = append(menu.Items, *item)
}

func (menu *Menu) IsExistNextMenu(title string) bool {
	if menu.Next(title) == nil {
		return false
	}
	return true
}