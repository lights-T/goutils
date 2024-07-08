package sftp

import "testing"

func TestSftp(t *testing.T) {
	newSFTP, err := Connect(&Sftp{
		Ctx:      nil,
		User:     "Emerson",
		Password: "DeltaVE1",
		Host:     "192.168.111.3",
		Port:     21,
		Timeout:  0,
	})
	if err != nil {
		t.Fatalf("An error occurred connecting to the remote server. %s", err.Error())
	}
	t.Log(newSFTP)
}
