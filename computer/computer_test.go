package computer

import "testing"

func Test_GetPhysicalID(t *testing.T) {
	id, err := GetPhysicalID()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}

func Test_GetMachineGuid(t *testing.T) {
	guid, err := GetMachineGuid()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(guid)
}

func Test_GetMACAddress(t *testing.T) {
	mac, err := GetMACAddress()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(mac)
}

func Test_GetCPUInfo(t *testing.T) {
	cpu, err := GetCPUInfo()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cpu)
}
