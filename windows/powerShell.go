package windows

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/lights-T/goutils"
	"github.com/lights-T/goutils/files"
)

var WIN_CREATE_STARTUP_SHORTCUT = `$WshShell = New-Object -comObject WScript.Shell
$Shortcut = $WshShell.CreateShortcut("STARTUP")
$Shortcut.TargetPath = "PLACEHOLDER"
$Shortcut.Save()`

var WIN_CREATE_DESKTOP_SHORTCUT = `$WshShell = New-Object -comObject WScript.Shell
$Shortcut = $WshShell.CreateShortcut("DESKTOP")
$Shortcut.TargetPath = "PLACEHOLDER"
$Shortcut.Save()`

type PowerShell struct {
	powerShell string
	Err        error
}

// New create new session
func New() (*PowerShell, error) {
	ps, err := exec.LookPath("powershell.exe")
	return &PowerShell{
		powerShell: ps,
	}, err
}

func (p *PowerShell) execute(args ...string) (stdOut string, stdErr string, err error) {
	args = append([]string{"-NoProfile", "-NonInteractive"}, args...)
	cmd := exec.Command(p.powerShell, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	stdOut, stdErr = stdout.String(), stderr.String()
	return
}

// EnableAutostartWin creates a shortcut to MyAPP in the shell:startup folder
func (p *PowerShell) EnableAutostartWin(execPath string) error {
	if r := files.CheckFileIsExist(execPath); !r {
		return goutils.Errorf("The file could not be found.")
	}
	name := strings.TrimSuffix(filepath.Base(execPath), ".exe")
	shortcutPath := fmt.Sprintf("$HOME\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\%s.lnk", name)
	WIN_CREATE_STARTUP_SHORTCUT = strings.Replace(WIN_CREATE_STARTUP_SHORTCUT, "STARTUP", shortcutPath, 1)
	WIN_CREATE_STARTUP_SHORTCUT = strings.Replace(WIN_CREATE_STARTUP_SHORTCUT, "PLACEHOLDER", execPath, 1)
	stdOut, stdErr, err := p.execute(WIN_CREATE_STARTUP_SHORTCUT)
	if err != nil {
		return goutils.Errorf("CreateShortcut: StdOut : '%s'; StdErr: '%s'; Err: %s",
			strings.TrimSpace(stdOut), stdErr, err.Error())
	}
	return nil
}

// EnableDesktopWin creates a shortcut to MyAPP in desktop
func (p *PowerShell) EnableDesktopWin(execPath string) error {
	if r := files.CheckFileIsExist(execPath); !r {
		return goutils.Errorf("The file could not be found.")
	}
	name := strings.TrimSuffix(filepath.Base(execPath), ".exe")
	shortcutPath := fmt.Sprintf("$HOME\\Desktop\\%s.lnk", name)
	WIN_CREATE_DESKTOP_SHORTCUT = strings.Replace(WIN_CREATE_DESKTOP_SHORTCUT, "DESKTOP", shortcutPath, 1)
	WIN_CREATE_DESKTOP_SHORTCUT = strings.Replace(WIN_CREATE_DESKTOP_SHORTCUT, "PLACEHOLDER", execPath, 1)
	stdOut, stdErr, err := p.execute(WIN_CREATE_DESKTOP_SHORTCUT)
	if err != nil {
		return goutils.Errorf("CreateShortcut: StdOut : '%s'; StdErr: '%s'; Err: %s",
			strings.TrimSpace(stdOut), stdErr, err.Error())
	}
	return nil
}
