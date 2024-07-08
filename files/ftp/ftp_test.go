package ftp

import "testing"

func TestFtp(t *testing.T) {
	newFTP, err := Connect(&Ftp{
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
	t.Run("Upload", func(t *testing.T) {
		if err = newFTP.Upload("./11.txt", "./uwu/QW2/as/qa"); err != nil {
			t.Fatal(err.Error())
		}
	})
	t.Run("Download", func(t *testing.T) {
		if err = newFTP.Download("./uwu/QW2/23.txt", "./newDir.txt"); err != nil {
			t.Fatal(err.Error())
		}
	})

}
