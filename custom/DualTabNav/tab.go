package DualTabNav

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type DualTabNav struct {
	OneBtn *widget.Button
	TwoBtn *widget.Button

	OneTab *fyne.Container
	TwoTab *fyne.Container

	AppScreen *fyne.Container
}

func NewDualTabNav(oneBtnLabel, twoBtnLabel string, oneBtnIcon, twoBtnIcon fyne.Resource, oneTabContainer, twoTabContainer *fyne.Container) *DualTabNav {

	oneBtn := widget.NewButtonWithIcon(oneBtnLabel, oneBtnIcon, func() {})
	twoBtn := widget.NewButtonWithIcon(twoBtnLabel, twoBtnIcon, func() {})

	oneBtnFunc := func() {
		oneBtn.Importance = widget.HighImportance
		twoBtn.Importance = widget.LowImportance

		// make other tabs hide
		oneTabContainer.Show()
		twoTabContainer.Hide()

		oneBtn.Refresh()
		twoBtn.Refresh()
	}
	twoBtnFunc := func() {
		oneBtn.Importance = widget.LowImportance
		twoBtn.Importance = widget.HighImportance

		// make other tabs hide
		oneTabContainer.Hide()
		twoTabContainer.Show()

		oneBtn.Refresh()
		twoBtn.Refresh()
	}

	oneBtn.OnTapped = oneBtnFunc
	twoBtn.OnTapped = twoBtnFunc

	DualTabNav := &DualTabNav{
		OneBtn: oneBtn,
		TwoBtn: twoBtn,

		OneTab: oneTabContainer,
		TwoTab: twoTabContainer,
	}

	DualTabNav.AppScreen = container.NewBorder(
		container.NewAdaptiveGrid(
			2,
			DualTabNav.OneBtn,
			DualTabNav.TwoBtn,
		),
		nil,
		nil,
		nil,
		container.NewStack(
			DualTabNav.OneTab,
			DualTabNav.TwoTab,
		),
	)

	// make other tabs hide and focus oneTab
	oneBtn.Importance = widget.HighImportance
	twoTabContainer.Hide()

	return DualTabNav
}
