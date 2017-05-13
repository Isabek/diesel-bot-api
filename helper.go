package main

func InitMenu() *Menu {
	menu := Menu{Title:"Что вы ищете?", IsRoot:true, Slug: MAIN_MENU_TYPE}

	menuPhoneItem := &Menu{Title: "Телефон", IsSubMenu:true}
	menu.InsertItem(menuPhoneItem)

	menuFlatRoomQty := &Menu{
		Title:"Количесвто комнат в квартире?",
		Slug: "rooms_qty",
		Items: []Menu{
			{Title: "1"},
			{Title: "2"},
			{Title: "3"},
			{Title: "4"},
			{Title:"Не важно"},
			{Title: "Назад"},
		},
	}
	menuFlatPriceMenu := &Menu{
		Title:"Цена в сомах?",
		Slug:"price",
		Items:[]Menu{
			{Title:"5000-10000"},
			{Title:"10000-15000"},
			{Title:"15000-20000"},
			{Title:"Не важно"},
			{Title:"Назад"},
		},
	}
	menuFlatItem := &Menu{Title: "Квартира", IsSubMenu:true}
	menuFlatItem.Insert(menuFlatRoomQty)
	menuFlatRoomQty.Insert(menuFlatPriceMenu)
	menu.InsertItem(menuFlatItem)
	return &menu
}