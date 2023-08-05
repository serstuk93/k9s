// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package ui

import (
	"fmt"
	"strings"

	"github.com/derailed/k9s/internal/config"
	"github.com/derailed/tview"
)

// LogoSmall K9s small log.
var LogoSmall = []string{
	`    ╬╬╬╙╙║╬╬╙╙╣╬╬`,
	`    ╬╬   ║╬╡   ╠╬`,
	`         ║╬╡     `,
	`    ╬╬╬  ║╬╡  ╬╬╬`,
	`         ║╬╡     `,
	`       ╓╗╣╬╬╗╖   `,
}

var LogoSmall1 = []string{
	`     ]@╗╗╗╗╖╓.,,`,
	`     ]╩   ╬╬  ╙╬`,
	`     '    ╬╬   ╚`,
	`     ]╗╗  ╬╬  ╗m`,
	`          ╬╬    `,
	`        ╔@╬╝Ñª  `,
}

var LogoSmall2 = []string{
	`        ╗╖,     `,
	`        ╩ ╞▌╙▒  `,
	`          ▐▌   `,
	`        ╬⌐▐▌ ╬  `,
	`          ▐▌    `,
	`          ╣╬ε   `,
}

var LogoSmall3 = []string{
	`           ╖    `,
	`           ╝    `,
	`           ⌐    `,
	`           b░   `,
	`                `,
	`           ε    `,
}

var LogoSmall4 = []string{
	`          ╓     `,
	`          Γ⌐    `,
	`          .     `,
	`         Γ░     `,
	`                `,
	`          ]µ    `,
}
var LogoSmall5 = []string{
	`        ,╓╖.╖ `,
	`       ▐╙║▌ ╙▌  `,
	`         ▐▌  ⌐  `,
	`       ╘Γ▐▌ ╬⌐  `,
	`         ╞▌     `,
	`        :╣╬     `,
}
var LogoSmall6 = []string{
	`     ,,╓╖╔╗╗╗@ε `,
	`     ║╩  ╬╬   ╚Γ`,
	`         ╬╬     `,
	`     └╜⌐ ╬╬  ╝╝Γ`,
	`         ╬╬     `,
	`       ªå╝╬▒╗   `,
}

// LogoBig K9s big logo for splash page.
var LogoBig = []string{
	`╬╬╬╙╙║╬╬╙╙╣╬╬           `,
	`╬╬   ║╬╡   ╠╬           `,
	`     ║╬╡                `,
	`╬╬╬  ║╬╡  ╬╬╬  ╬╬╬  ╬╬╬ `,
	`     ║╬╡                `,
	`   ╓╗╣╬╬╗╖              `,
}

var currentNum = 0

func LogoChanger() []string {
	logos := [][]string{
		LogoSmall, LogoSmall1, LogoSmall2, LogoSmall3, LogoSmall4, LogoSmall5, LogoSmall6}
	currentNum++

	index := currentNum / 2

	if index > len(logos)-1 {
		currentNum = 0
		return logos[0]
	}

	return logos[index]
}

// Splash represents a splash screen.
type Splash struct {
	*tview.Flex
}

// NewSplash instantiates a new splash screen with product and company info.
func NewSplash(styles *config.Styles, version string) *Splash {
	s := Splash{Flex: tview.NewFlex()}
	s.SetBackgroundColor(styles.BgColor())

	logo := tview.NewTextView()
	logo.SetDynamicColors(true)
	logo.SetTextAlign(tview.AlignCenter)
	s.layoutLogo(logo, styles)

	vers := tview.NewTextView()
	vers.SetDynamicColors(true)
	vers.SetTextAlign(tview.AlignCenter)
	s.layoutRev(vers, version, styles)

	s.SetDirection(tview.FlexRow)
	s.AddItem(logo, 10, 1, false)
	s.AddItem(vers, 1, 1, false)

	return &s
}

func (s *Splash) layoutLogo(t *tview.TextView, styles *config.Styles) {
	logo := strings.Join(LogoBig, fmt.Sprintf("\n[%s::b]", styles.Body().LogoColor))
	fmt.Fprintf(t, "%s[%s::b]%s\n",
		strings.Repeat("\n", 2),
		styles.Body().LogoColor,
		logo)
}

func (s *Splash) layoutRev(t *tview.TextView, rev string, styles *config.Styles) {
	fmt.Fprintf(t, "[%s::b]Revision [red::b]%s", styles.Body().FgColor, rev)
}
