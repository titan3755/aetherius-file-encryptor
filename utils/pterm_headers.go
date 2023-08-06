package utils

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func PtermHeaders(headerTitle string, colorPterm pterm.Color) {
	header := pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgRed))
	pterm.DefaultCenter.Println(header.Sprint("Aetherius File Encryptor"))
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle(headerTitle, &pterm.Style{colorPterm}),
	).Render()
}