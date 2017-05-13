package main

import (
	"testing"
)

func TestMenu_Root(t *testing.T) {
	menu := Menu{Title: "Main Menu", IsRoot:true}
	if menu.IsRoot != true {
		t.Error("Expected true, got ", false)
	}
}

func TestMenu_Insert(t *testing.T) {
	menu := Menu{Title: "Main Menu", IsRoot:true}
	secondMenu := &Menu{Title: "Second Menu"}
	menu.Insert(secondMenu)
	if menu.Title != "Main Menu" || menu.next.Title != "Second Menu" {
		t.Error("Expected Main Menu and Second Menu, got ", menu.Title, menu.next.Title)
	}
}

func TestMenu_InsertItem(t *testing.T) {
	menu := Menu{Title: "Main Menu", IsRoot:true}
	menuItem := &Menu{Title:"Menu Item"}
	menu.InsertItem(menuItem)
	if menu.Title != "Main Menu" || menu.Items[0].Title != "Menu Item" {
		t.Error("Expected Main Menu has one Item, got zero Item")
	}
}

func TestMenu_Next(t *testing.T) {
	menu := Menu{Title: "Main Menu", IsRoot:true}
	secondMenu := &Menu{Title: "Second Menu"}
	menu.Insert(secondMenu)
	if menu.Title != "Main Menu" || menu.Next("Second Menu").Title != "Second Menu" {
		t.Error("Expected Next Menu Second Menu, got ", menu.next.Title)
	}
}

func TestMenu_Prev(t *testing.T) {
	menu := Menu{Title: "Main Menu", IsRoot:true}
	secondMenu := &Menu{Title: "Second Menu"}
	secondMenu = menu.Insert(secondMenu)
	if secondMenu.Prev().Title != "Main Menu" {
		t.Error("Expected Main Menu, got ", secondMenu.Prev().Title)
	}
}

func TestMenu_InsertItem2(t *testing.T) {
	menu := &Menu{Title:"Root Level", IsRoot:true}
	menuItem1 := &Menu{Title:"Root Level Item 1"}
	menuItem2 := &Menu{Title:"Root Level Item 2"}
	menuItem3 := &Menu{Title:"Root Level Item 3"}

	menuItem1NextMenu := &Menu{Title: "First Level"}
	menuItem1NextMenuItem1 := &Menu{Title: "First Level Item 1"}
	menuItem1NextMenuItem2 := &Menu{Title: "First Level Item 2"}

	menuItem1NextMenu.InsertItem(menuItem1NextMenuItem1)
	menuItem1NextMenu.InsertItem(menuItem1NextMenuItem2)

	menuItem1.Insert(menuItem1NextMenu)

	menu.InsertItem(menuItem1)
	menu.InsertItem(menuItem2)
	menu.InsertItem(menuItem3)

	if menuItem1NextMenuItem1.Prev().Title != "Root Level Item 1" {
		t.Error("Expected Root Level Item 1, got ", menuItem1NextMenuItem1.Prev().Title)
	}
	if menu.Next("Root Level Item 1").Title != "First Level" {
		t.Error("Expected First Level, got ", menu.Next("Root Level Item 1").Title)
	}
}