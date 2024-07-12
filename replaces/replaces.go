package replaces

import (

	"serialportsimulator/utils"
	"fmt"
)

func ReplaceData(datas []byte,functionCode byte,registerAddress uint16,registerCount uint16) {
	//"03 0050 001C"
	if functionCode==3 && registerAddress==80 && registerCount==28 {
       SimAEM96(datas)
	}
	
}

//模拟安科瑞AEM96多功能表
// 03 0050 001C",
// "A相电压^V^/1 B相电压^V^/1 C相电压^V^/1 =5
//AB相电压^V^/1 BC相电压^V^/1 CA相电压^V^/1 =11
// A相电流^A^/3 B相电流^A^/3 C相电流^A^/3 零序电流^A^/3",19
// "A有功功率^KW^/4 B有功功率^KW^/4 C有功功率^KW^/4 总有功功率^KW^/4 27 A无功功率^KVar^/4 B无功功率^KVar^/4 C无功功率^KVar^/4 总无功功率^KVar^/4 35 
//A视在功率^KVA^/4 B视在功率^KVA^/4",39
// "C视在功率^KVA^/4 总视在功率^KVA^/4 A功率因数^—^/3 B功率因数^—^/3 C功率因数^—^/3 总功率因数^—^/3 50
//功率方向^- 频率^Hz^/2",57
// "ret": "i2=[0,27] i4=[28,32] i2=[33,38] i2=[39,43]",
// "ctv": "v=[0,5] c=[6,9] t=[10,21] n=[22,27] t=[28,32] n=[33,43]"
func SimAEM96(datas []byte) {
    fmt.Println("安科瑞  AEM96  多功能电流表")

	b1,b2 := utils.SplitUint16ToBytes(2200)
	datas[0]=b1
	datas[1]=utils.RandomInRange(b2)
	b1,b2 = utils.SplitUint16ToBytes(2400)
	datas[2]=b1
	datas[3]=utils.RandomInRange(b2)
	b1,b2 = utils.SplitUint16ToBytes(2300)
	datas[4]=b1
	datas[5]=utils.RandomInRange(b2)

    
	b1,b2 = utils.SplitUint16ToBytes(3800)
	datas[6]=b1
	datas[7]=utils.RandomInRange(b2)
	b1,b2 = utils.SplitUint16ToBytes(3800)
	datas[8]=b1
	datas[9]=utils.RandomInRange(b2)
	b1,b2 = utils.SplitUint16ToBytes(3800)
	datas[10]=b1
	datas[11]=utils.RandomInRange(b2)

	
	//4电流
	b1,b2 = utils.SplitUint16ToBytes(1200)
	datas[12]=b1
	datas[13]=utils.RandomInRange(b2)
	b1,b2 = utils.SplitUint16ToBytes(900)
	datas[14]=b1
	datas[15]=utils.RandomInRange(b2)
	b1,b2 = utils.SplitUint16ToBytes(1100)
	datas[16]=b1
	datas[17]=utils.RandomInRange(b2)

	b1,b2 = utils.SplitUint16ToBytes(10)
	datas[18]=b1
	datas[19]=utils.RandomInRange(b2)

	//4有功27

	//4无功2829 3031 3233 3435

	//4视在3637 3839 4041 4243
	//4功率因数
	// fillByte(datas,1000,44,51)
	v1,v2 := utils.SplitUint16ToBytes(1000)
	datas[44]=v1
	datas[45]=utils.RandomInRange(v2)
	datas[46]=v1
	datas[47]=utils.RandomInRange(v2)
	datas[48]=v1
	datas[49]=utils.RandomInRange(v2)
	datas[50]=v1
	datas[51]=utils.RandomInRange(v2)
	//功率方向、
	//频率
	b1,b2 = utils.SplitUint16ToBytes(5000)
	datas[54]=b1
	datas[55]=b2

}
