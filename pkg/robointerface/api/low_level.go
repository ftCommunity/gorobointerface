package api

/*
#cgo LDFLAGS: -L/usr/local/lib/ -lroboint
#include "roboint.h"
*/
import "C"
import (
	"unsafe"

	"github.com/pkg/errors"
)

//GetNumUSBDevices : get number of usb devices
func GetNumUSBDevices() int {
	C.InitFtUsbDeviceList()
	return int(C.GetNumFtUsbDevice())
}

//Robointerface : struct for the abstraction of the robointerface
type Robointerface struct {
	handle       C.FT_HANDLE
	transferarea *C.FT_TRANSFER_AREA
}

func (r *Robointerface) InitUSB(config IFConfigUSB) error {
	C.InitFtUsbDeviceList()
	if config.iUSBSerial != 0 {
		r.handle = C.GetFtUsbDeviceHandleSerialNr(C.long(config.iUSBSerial), C.FT_AUTO_TYPE)
	} else {
		r.handle = C.GetFtUsbDeviceHandle(C.uchar(config.iDevice))
	}
	C.OpenFtUsbDevice(r.handle)
	return r.init(config.config)
}
func (r *Robointerface) InitSerial(config IFConfigSerial) error {
	r.handle = C.OpenFtCommDevice((*C.char)(unsafe.Pointer(&config.serialDevice)), C.long(config.SerialType), 10)
	return r.init(config.config)
}

func (r *Robointerface) init(config IFConfig) error {
	var d int
	if config.bEnableDist {
		d = 1
	} else {
		d = 0
	}
	C.SetFtDistanceSensorMode(r.handle, (C.long)(d), 20, 20, 100, 100, 3, 3)
	if config.bStartTransferArea {
		C.StartFtTransferArea(r.handle, nil)
		r.transferarea = C.GetFtTransferAreaAddress(r.handle)
		if r.transferarea == nil {
			return errors.New("Could not open transfer area!")
		}
	}
	return nil
}

func (r Robointerface) GetAV() int {
	return r.avToV(r.transferarea.AV)
}

func (r Robointerface) getDeviceType() C.long {
	return C.GetFtDeviceTyp(r.handle)
}

func (r Robointerface) avToV(av C.ushort) int {
	t := r.getDeviceType()
	if t == C.FT_ROBO_IO_EXTENSION || t == C.FT_ROBO_LT_CONTROLLER {
		return (int)(0.03 * (float32)(av))
	}
	return (int)(8.63*(float32)(av) - 1775)
}

func (r Robointerface) GetAX() int {
	return (int)(r.transferarea.AX)
}

func (r Robointerface) GetAY() int {
	return (int)(r.transferarea.AY)
}

func (r Robointerface) GetAZ() int {
	return (int)(r.transferarea.AZ)
}

func (r Robointerface) GetA1() int {
	return (int)(r.transferarea.A1)
}

func (r Robointerface) GetA2() int {
	return (int)(r.transferarea.A2)
}

func (r Robointerface) GetD1() int {
	return (int)(r.transferarea.D1)
}

func (r Robointerface) GetD2() int {
	return (int)(r.transferarea.D2)
}

func (r Robointerface) GetAVSlave1() int {
	return r.avToV(r.transferarea.AVS1)
}

func (r Robointerface) GetAVSlave2() int {
	return r.avToV(r.transferarea.AVS2)
}

func (r Robointerface) GetAVSlave3() int {
	return r.avToV(r.transferarea.AVS3)
}

func (r Robointerface) GetAXSlave1() int {
	return (int)(r.transferarea.AXS1)
}

func (r Robointerface) GetAXSlave2() int {
	return (int)(r.transferarea.AXS2)
}

func (r Robointerface) GetAXSlave3() int {
	return (int)(r.transferarea.AXS3)
}

func (r Robointerface) GetIR() int {
	return (int)(r.transferarea.IRKeys)
}

func (r Robointerface) GetNumDevices() int {
	return (int)(r.transferarea.BusModules)
}

func (r Robointerface) GetInput(i uint) bool {
	if i <= 8 {
		return (int)((r.transferarea.E_Main&1<<(i-1))>>(i-1)) == 1
	} else if i <= 16 {
		i -= 8
		return (int)((r.transferarea.E_Sub1&1<<(i-1))>>(i-1)) == 1
	} else if i <= 24 {
		i -= 16
		return (int)((r.transferarea.E_Sub2&1<<(i-1))>>(i-1)) == 1
	} else if i <= 32 {
		i -= 24
		return (int)((r.transferarea.E_Sub3&1<<(i-1))>>(i-1)) == 1
	}
	return false
}

func (r Robointerface) Digital(i uint) bool {
	return r.GetInput(i)
}
