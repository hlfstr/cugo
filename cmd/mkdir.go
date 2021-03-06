// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/hlfstr/cugo/src/mkdir"
	"github.com/hlfstr/flagger"
)

type mkdirCmd struct {
	Mode    uint
	Parents bool
	Verbose bool
}

func (m *mkdirCmd) Prepare(flags *flagger.Flags) {
	flags.BoolVar(&m.Parents, "Create missing parent directories", "-p", "--parents")
	flags.BoolVar(&m.Verbose, "Display each directory after it was created", "-v", "--verbose")
	flags.UintVar(&m.Mode, 0755, "Set permissions to MODE value", "-m", "--mode")
}

func (m *mkdirCmd) Action(s []string, flags *flagger.Flags) error {
	if data, err := flags.Parse(s); err != nil {
		return err
	} else {
		mkdir.Mkdir(m.Mode, m.Parents, m.Verbose, data)
	}
	return nil
}

func init() {
	Command.Add("mkdir", &mkdirCmd{})
}
