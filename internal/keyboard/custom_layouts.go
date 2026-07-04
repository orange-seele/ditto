package keyboard

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/arvingarciabtw/ditto/internal/keyboard/base"
)

type customLayoutFile struct {
	Map   map[string]string `json:"map"`
	Shift map[string]string `json:"shift,omitempty"`
}

var customLayoutNames []string

func loadCustomLayouts() {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get config dir: %v\n", err)
		return
	}
	dir := filepath.Join(cfgDir, "ditto", "layouts")

	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		name := strings.TrimSuffix(entry.Name(), ".json")
		if name == "config" {
			continue
		}

		data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to read layout %s: %v\n", entry.Name(), err)
			continue
		}

		var clf customLayoutFile
		if err := json.Unmarshal(data, &clf); err != nil {
			fmt.Fprintf(os.Stderr, "failed to parse layout %s: %v\n", entry.Name(), err)
			continue
		}

		if clf.Map == nil {
			fmt.Fprintf(os.Stderr, "layout %s has no map, skipping\n", entry.Name())
			continue
		}

		customLayoutNames = append(customLayoutNames, name)
		layouts[name] = clf.Map
		if clf.Shift != nil {
			shiftMaps[name] = clf.Shift
		} else {
			shiftMaps[name] = base.USShift
		}
	}
}
