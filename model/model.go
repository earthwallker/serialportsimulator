package model

type DeviceProtocol struct {
	Cmd string
	Data string
	Tag string
}

var(

	Protocols []DeviceProtocol
	ABRN02Y_Cmd1,ABRN02Y_Cmd2,ABRN02Y_Data1,ABRN02Y_Data2,ABRN02Y_Sg_Cmd1,ABRN02Y_Sg_Data1 string //ABRN02Y多功能表协议印02Y
	ABRN01Y_Cmd1,ABRN01Y_Data1,ABRN01Y_Sg_Cmd1,ABRN01Y_Sg_Data1 string //ABRN01Y多功能表协议
	DLGD01_Cmd1,DLGD01_Cmd2,DLGD01_Data1,DLGD01_Data2,DLGD01_Sg_Cmd1,DLGD01_Sg_Data1 string //DLGD01多功能表协议

)

func init(){

	    initABRN02Y()
        initDLGD01()
		initABRN01Y()

		Protocols = []DeviceProtocol{
			{Cmd: ABRN02Y_Cmd1, Data: ABRN02Y_Data1, Tag: "ABRN02Y_Cmd1"},
			{Cmd: ABRN02Y_Cmd2, Data: ABRN02Y_Data2, Tag: "ABRN02Y_Cmd2"},
			{Cmd: ABRN02Y_Sg_Cmd1, Data: ABRN02Y_Sg_Data1, Tag: "ABRN02Y_Sg"},

			{Cmd: DLGD01_Cmd1, Data: DLGD01_Data1, Tag: "DLGD01_Cmd1"},
			{Cmd: DLGD01_Cmd2, Data: DLGD01_Data2, Tag: "DLGD01_Cmd2"},
			{Cmd: DLGD01_Sg_Cmd1, Data: DLGD01_Sg_Data1, Tag: "DLGD01_Sg"},

			{Cmd: ABRN01Y_Cmd1, Data: ABRN01Y_Data1, Tag: "ABRN01Y_Cmd1"},
			{Cmd: ABRN01Y_Sg_Cmd1, Data: ABRN01Y_Sg_Data1, Tag: "ABRN01Y_Sg"},
		}
}

//初始化ABRN01表协议slaveId 0B
func initABRN01Y() {
	//命令必须小写字母
	ABRN01Y_Cmd1 = "03 0025 0022"		//"03 003D 0031"
	ABRN01Y_Sg_Cmd1 = "02 0000 0008"   
	ABRN01Y_Data1 = "09 0B 09 0B 09 0B 0F AA 0F AA 0F AA 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 03 E8 03 E8 03 E8 03 E8 00 00 00 00 00 00 00 00 13 88 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00"
	
	ABRN01Y_Sg_Data1 = "00"  // 返回值"0A 03 02 00 00 1D 85"
}

//初始化DLGD01表协议
func initDLGD01() {
	//命令必须小写字母
	DLGD01_Cmd1 = "03 0017 0028"		//"03 003D 0031"
	DLGD01_Cmd2 = "03 003F 0022"		//"03 00E0 0006"
	DLGD01_Sg_Cmd1 = "03 0036 0001"     //待定
	
	DLGD01_Data1 = "43 67 4C 38 43 67 51 D0 43 67 47 A8 43 C8 51 B6 43 C8 4F BC 43 C8 4D 50 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00"
	DLGD01_Data2 = "00 00 00 00 3F 80 00 00 3F 80 00 00 3F 80 00 00 3F 80 00 00 42 48 2A 63 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00"
	DLGD01_Sg_Data1 = "00 00"  // 返回值"0A 03 02 00 00 1D 85"
}

//初始化ABRN02Y表协议 slaveId=10  表上印字02Y
func initABRN02Y() {
			//命令必须小写字母
			ABRN02Y_Cmd1 = "03003d0031"		//"03 003D 0031"
			ABRN02Y_Cmd2 = "0300e00006"		//"03 00E0 0006"
			ABRN02Y_Sg_Cmd1 = "0300360001"   
			
			//数据帧的:第三个字节=有效数据长度;
			//数据区总长度:地址域+功能码+有效数据+校验码    
			//CRC:校验除了自己外的总长度
			ABRN02Y_Data1 = "09 24 09 26 09 24 0F D7 0F D7 0F D5 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 13 8A 00 00 00 05 00 00 00 00 00 00 00 09 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 AA AA 00 01 0A 02 00 00 00 00 00 00 00 00 00 00 00 01 00 01 00 00"
			ABRN02Y_Data2 = "00 E6 00 E6 00 E9 00 00 00 00 00 00"
			ABRN02Y_Sg_Data1 = "00 00"  // 返回值"0A 03 02 00 00 1D 85"
}