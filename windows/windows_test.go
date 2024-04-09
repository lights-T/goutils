package windows

import "testing"

func Test_Notify(t *testing.T) {
	Notify("Example App", "notification", "Some message about how important something is...")
}
