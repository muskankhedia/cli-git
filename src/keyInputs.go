package main

import "github.com/chzyer/readline"

// const (
// 	CharLineStart = 1
// 	CharBackward  = 2
// 	CharInterrupt = 3
// 	CharDelete    = 4
// 	CharLineEnd   = 5
// 	CharForward   = 6
// 	CharBell      = 7
// 	CharCtrlH     = 8
// 	CharTab       = 9
// 	CharCtrlJ     = 10
// 	CharKill      = 11
// 	CharCtrlL     = 12
// 	CharEnter     = 13
// 	CharNext      = 14
// 	CharPrev      = 16
// 	CharBckSearch = 18
// 	CharFwdSearch = 19
// 	CharTranspose = 20
// 	CharCtrlU     = 21
// 	CharCtrlW     = 23
// 	CharCtrlY     = 25
// 	CharCtrlZ     = 26
// 	CharEsc       = 27
// 	CharEscapeEx  = 91
// 	CharBackspace = 127
// )

const (
	// KeyBackward is the default key to page up during selection.
	KeyBackward        rune = readline.CharCtrlW
	KeyBackwardDisplay      = "ctrl + w"

	// KeyForward is the default key to page down during selection.
	KeyForward        rune = readline.CharCtrlL
	KeyForwardDisplay      = "ctrl + l"
)
