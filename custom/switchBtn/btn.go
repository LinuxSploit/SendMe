package switchBtn

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// SwitchButton a New switch btn widget
type SwitchButton struct {
	widget.Button

	onLabel  string
	offLabel string

	onIcon  fyne.Resource
	offIcon fyne.Resource

	onImportance  widget.Importance
	offImportance widget.Importance

	onFunc  func()
	offFunc func()

	Status bool
}

// Tapped make button to Operate as switch btn
func (btn *SwitchButton) Tapped(_ *fyne.PointEvent) {

	// Update label, icon, and importance based on the current state
	if btn.Status {
		// run OFF service func
		btn.offFunc()
	} else {
		// run ON service func
		btn.onFunc()
	}

}

// NewSwitchButton creates a new SwitchButton
func NewSwitchButton(onLabel, offLabel string, onIcon, offIcon fyne.Resource, onImportance widget.Importance, offImportance widget.Importance, onFunc func(), offFunc func()) *SwitchButton {
	tmp := &SwitchButton{
		onLabel:  onLabel,
		offLabel: offLabel,

		onIcon:  onIcon,
		offIcon: offIcon,

		onImportance:  onImportance,
		offImportance: offImportance,

		onFunc:  onFunc,
		offFunc: offFunc,

		Status: false,
	}

	tmp.SetText(tmp.offLabel)
	tmp.SetIcon(tmp.offIcon)

	tmp.Importance = tmp.offImportance

	tmp.ExtendBaseWidget(tmp)

	return tmp
}

func (btn *SwitchButton) ToggleSwitchState() {

	// Update label, icon, and importance based on the current state
	if btn.Status {
		btn.SetText(btn.offLabel)
		btn.SetIcon(btn.offIcon)
		btn.Importance = btn.offImportance
		btn.Refresh()
	} else {
		btn.SetText(btn.onLabel)
		btn.SetIcon(btn.onIcon)
		btn.Importance = btn.onImportance
		btn.Refresh()
	}

	// Toggle the switch state
	btn.Status = !btn.Status

}
