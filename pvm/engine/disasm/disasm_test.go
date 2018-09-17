package disasm_test

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"

	"paradigm-vm/pvm/engine/disasm"
	"paradigm-vm/pvm/engine/wasm"
)

func TestDisassemble(t *testing.T) {
	for _, dir := range testPaths {
		fnames, err := filepath.Glob(filepath.Join(dir, "*.wasm"))
		if err != nil {
			t.Fatal(err)
		}
		for _, fname := range fnames {
			name := fname
			t.Run(filepath.Base(name), func(t *testing.T) {
				raw, err := ioutil.ReadFile(name)
				if err != nil {
					t.Fatal(err)
				}

				r := bytes.NewReader(raw)
				m, err := wasm.ReadModule(r, nil)
				if err != nil {
					t.Fatalf("error reading module %v", err)
				}
				for _, f := range m.FunctionIndexSpace {
					_, err := disasm.NewDisassembly(f, m)
					if err != nil {
						t.Fatalf("disassemble failed: %v", err)
					}
				}
			})
		}
	}
}
