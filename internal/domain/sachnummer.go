// internal/domain/sachnummer.go
package domain

import (
	"fmt"
	"strings"
)

type SachnummerKey struct {
	TypSymbol             int
	HerstellungsartSymbol int
	VerschleissteilSymbol int
	FunktionSymbol        int
	MaterialSymbol        int
	OberflaecheSymbol     int
	FarbeSymbol           int
	ReserveSymbol         int
}

func (k SachnummerKey) GroupKey() string {
	return fmt.Sprintf(
		"T%d-H%d-V%d-F%d-M%d-O%d-C%d-R%d",
		k.TypSymbol,
		k.HerstellungsartSymbol,
		k.VerschleissteilSymbol,
		k.FunktionSymbol,
		k.MaterialSymbol,
		k.OberflaecheSymbol,
		k.FarbeSymbol,
		k.ReserveSymbol,
	)
}

func GenerateHexSuffix(idx int64) string {
	hex := strings.ToUpper(fmt.Sprintf("%X", idx))
	if len(hex) < 4 {
		hex = strings.Repeat("0", 4-len(hex)) + hex
	}
	return hex
}

func BuildSachnummer(key SachnummerKey, suffix string) string {
	return fmt.Sprintf(
		"%d-%d-%d-%d-%d-%d-%d-%d-%s",
		key.TypSymbol,
		key.HerstellungsartSymbol,
		key.VerschleissteilSymbol,
		key.FunktionSymbol,
		key.MaterialSymbol,
		key.OberflaecheSymbol,
		key.FarbeSymbol,
		key.ReserveSymbol,
		suffix,
	)
}
