package api

import (
	"pvftools/backend/common/log"
	"pvftools/backend/internal/box"

	"github.com/dof-dev/godof/parser"
)

func (a *App) GenRandomBox(params *box.BoxParams) string {
	p := box.GenerateRandomBox(params)
	return "#PVF_File\n" + p.Render()
}

func (a *App) GenBox(params *box.BoxParams) string {
	var p *parser.Parser
	switch params.BoxType {
	case box.BoxTypeRandom:
		p = box.GenerateRandomBox(params)
	case box.BoxTypeSelection:
		p = box.GenerateSelectionBox(params)
	default:
		log.LogError("Invalid box type")
		return ""
	}
	return "#PVF_File\n" + p.Render()
}
