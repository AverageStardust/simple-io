package input

var onSelected func()
var onEntered func()

func callOnSelected() {
	if onSelected != nil {
		onSelected()
	}
}

func callOnEntered() {
	if onEntered != nil {
		onEntered()
	}
}

func SetOnSelected(listener func()) {
	onSelected = listener
}

func SetOnEntered(listener func()) {
	onEntered = listener
}

func UnsetOnSelected() {
	onSelected = nil
}

func UnsetOnEntered() {
	onEntered = nil
}
