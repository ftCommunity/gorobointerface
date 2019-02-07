package api

// IFConfig is the configuration for interface
type IFConfig struct {
	bEnableDist        bool
	bStartTransferArea bool
}

//IFConfigSerial : configuration for a interface via serial port
type IFConfigSerial struct {
	config       IFConfig
	serialDevice int
	SerialType   int32
}

//IFConfigUSB : configuration for a interface via USB
type IFConfigUSB struct {
	config     IFConfig
	iDevice    int
	iUSBSerial int32
}

func GetDefaultConfig() IFConfig {
	return IFConfig{true, true}
}

func GetDefaultUSBConfig() IFConfigUSB {
	return IFConfigUSB{GetDefaultConfig(), 0, 0}
}
