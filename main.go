package main

import (
	"time"

	"github.com/ZongBen/GoFive/pkg/control"
	"github.com/ZongBen/GoFive/pkg/dialog"
	"github.com/ZongBen/GoFive/pkg/game"
	"github.com/ZongBen/GoFive/pkg/gui"
	"github.com/ZongBen/GoFive/pkg/menu"
	"github.com/ZongBen/GoFive/pkg/online"
)

var _homeMenu menu.HomeMenu
var _onlineMenu menu.OnlineMenu

func init() {
	homeMenu := menu.CreateHomeMenu()
	_homeMenu = &homeMenu

	onlineMenu := menu.CreateOnlineMenu()
	_onlineMenu = &onlineMenu
}

func main() {
	for !_homeMenu.IsQuit() {
		gui.Flush(34, 20, gui.RenderHome(_homeMenu), true)
		command := control.ExecuteCommand(_homeMenu, control.HomeMenuCommandHandler)
		switch command {
		case menu.LOCAL_PLAY:
			StartLocalGame()
		case menu.ONLINE_PLAY:
			OnlineGameMenu()
		case menu.EXIT:
			_homeMenu.Quit()
		}
	}
	gui.Clear()
	gui.Close()
}

func OnlineGameMenu() {
	gui.Clear()
	state := -1
	for state == -1 {
		gui.Flush(34, 20, gui.RenderOnline(_onlineMenu), true)
		state = control.ExecuteCommand(_onlineMenu, control.OnlineMenuCommandHandler)
		switch state {
		case menu.JOIN:
			JoinGame()
		case menu.HOST:
			WaitHostGame()
			break
		case menu.ONLINE_BACK:
			_onlineMenu.SetMenuSelect(menu.JOIN)
			break
		}
	}
}

func WaitHostGame() {
	ch := make(chan int)
	go online.StartHostServer(ch)
	go renderWaiting(ch)
	waitCommand(ch)
}

func renderWaiting(ch <-chan int) {
	dot := 0
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return
			}
		case <-time.After(1 * time.Second):
			dot = (dot + 1) % 4
			gui.Flush(34, 20, gui.RenderHostGame(dot), true)
		}
	}
}

func waitCommand(ch chan int) {
	state := -1
	for state == -1 {
		state = control.ExecuteCommand(ch, control.HostGameCommandHandler)
	}
}

func JoinGame() {
	var ip string
	state := -1
	for state == -1 {
		gui.Flush(34, 20, gui.RenderJoinGame(ip), true)
		state = control.ExecuteCommand(&ip, control.JoinGameCommandHandler)
	}
	if state == control.ESC {
		OnlineGameMenu()
	} else if state == control.ENTER_IP {
		online.ConnectToHost()
	}
}

func StartLocalGame() {
	var _gameBoard game.Board
	b := game.CreateBoard()
	_gameBoard = &b
	for !_gameBoard.IsFinish() {
		if _gameBoard.GetWinner() != 0 {
			result := showDialog(_gameBoard)
			if result == dialog.AGAIN {
				StartLocalGame()
				break
			} else if result == dialog.QUIT {
				_gameBoard.Quit()
				break
			}
		}
		gui.Flush(126, 60, gui.RenderBoard(_gameBoard), true)
		control.ExecuteCommand(_gameBoard, control.GameCommandHandler)
	}
	gui.Clear()
}

func showDialog(b game.Board) int {
	state := -1
	for state == -1 {
		gui.Flush(126, 60, gui.RenderBoard(b), true)
		state = control.ExecuteCommand(b.GetDialog(), control.DialogCommandHandler)
	}
	return state
}
