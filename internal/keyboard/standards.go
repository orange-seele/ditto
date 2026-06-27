package keyboard

type standardData struct {
	sizes    map[int][][]key
	shiftMap map[string]string
	altGrMap map[string]string
}

var standards = map[string]standardData{
	"ansi": {sizes: sizesANSI, shiftMap: usShift},
	"iso":  {sizes: sizesISO, shiftMap: ukShift},
	"abnt": {sizes: sizesABNT, shiftMap: abntShift, altGrMap: abntAltGr},
	"jis":  {sizes: sizesJIS, shiftMap: jisShift},
	"ks":   {sizes: sizesKS, shiftMap: ksShift},
}

var StandardListItems = []string{
	"ansi",
	"iso",
	"abnt",
	"jis",
	"ks",
}
