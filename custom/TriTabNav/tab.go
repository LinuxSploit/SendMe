package TriTabNav

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type TriTabNav struct {
	OneBtn   *widget.Button
	TwoBtn   *widget.Button
	ThirdBtn *widget.Button

	OneTab   *fyne.Container
	TwoTab   *fyne.Container
	ThirdTab *fyne.Container

	AppScreen *fyne.Container
}

func NewTriTabNav(oneBtnBlueIcon, oneBtnWhiteIcon, twoBtnBlueIcon, twoBtnWhiteIcon, thirdBtnBlueIcon, thirdBtnWhiteIcon fyne.Resource, oneTabContainer, twoTabContainer, thirdTabContainer *fyne.Container) *TriTabNav {

	oneBtn := widget.NewButtonWithIcon("", oneBtnWhiteIcon, func() {})
	twoBtn := widget.NewButtonWithIcon("", twoBtnBlueIcon, func() {})
	thirdBtn := widget.NewButtonWithIcon("", thirdBtnBlueIcon, func() {})

	oneBtn.Importance = widget.LowImportance
	twoBtn.Importance = widget.LowImportance
	thirdBtn.Importance = widget.LowImportance

	oneBtnFunc := func() {

		oneBtn.Importance = widget.HighImportance
		twoBtn.Importance = widget.LowImportance
		thirdBtn.Importance = widget.LowImportance

		oneBtn.SetIcon(oneBtnWhiteIcon)
		twoBtn.SetIcon(twoBtnBlueIcon)
		thirdBtn.SetIcon(thirdBtnBlueIcon)

		// make other tabs hide
		oneTabContainer.Show()
		twoTabContainer.Hide()
		thirdTabContainer.Hide()

		oneBtn.Refresh()
		twoBtn.Refresh()
		thirdBtn.Refresh()
	}
	twoBtnFunc := func() {
		oneBtn.Importance = widget.LowImportance
		twoBtn.Importance = widget.HighImportance
		thirdBtn.Importance = widget.LowImportance

		oneBtn.SetIcon(oneBtnBlueIcon)
		twoBtn.SetIcon(twoBtnWhiteIcon)
		thirdBtn.SetIcon(thirdBtnBlueIcon)

		// make other tabs hide
		oneTabContainer.Hide()
		twoTabContainer.Show()
		thirdTabContainer.Hide()

		oneBtn.Refresh()
		twoBtn.Refresh()
		thirdBtn.Refresh()
	}
	thirdBtnFunc := func() {
		oneBtn.Importance = widget.LowImportance
		twoBtn.Importance = widget.LowImportance
		thirdBtn.Importance = widget.HighImportance

		oneBtn.SetIcon(oneBtnBlueIcon)
		twoBtn.SetIcon(twoBtnBlueIcon)
		thirdBtn.SetIcon(thirdBtnWhiteIcon)
		// make other tabs hide
		oneTabContainer.Hide()
		twoTabContainer.Hide()
		thirdTabContainer.Show()

		oneBtn.Refresh()
		twoBtn.Refresh()
		thirdBtn.Refresh()
	}

	oneBtn.OnTapped = oneBtnFunc
	twoBtn.OnTapped = twoBtnFunc
	thirdBtn.OnTapped = thirdBtnFunc

	triTabNav := &TriTabNav{
		OneBtn:   oneBtn,
		TwoBtn:   twoBtn,
		ThirdBtn: thirdBtn,

		OneTab:   oneTabContainer,
		TwoTab:   twoTabContainer,
		ThirdTab: thirdTabContainer,
	}

	triTabNav.AppScreen = container.NewBorder(
		nil,
		container.NewAdaptiveGrid(
			3,
			triTabNav.OneBtn,
			triTabNav.TwoBtn,
			triTabNav.ThirdBtn,
		),
		nil,
		nil,
		container.NewStack(
			triTabNav.OneTab,
			triTabNav.TwoTab,
			triTabNav.ThirdTab,
		),
	)

	// make other tabs hide and focus oneTab
	oneBtn.Importance = widget.HighImportance
	twoTabContainer.Hide()
	thirdTabContainer.Hide()

	return triTabNav
}
