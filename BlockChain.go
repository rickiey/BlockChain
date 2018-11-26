package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//区块结构
type Block struct {
	Timestamp    int64  //时间戳
	Data         []byte //当前区块内容
	PreBlockHash []byte //上一个区块的hash
	Hash         []byte //当前hash
}

//设置区块hash的方法
func (this *Block) SetHash() {
	//当前区块的Hash，包括（timstamp+Date+PreBlockHash）
	//时间由int转 []byte
	timestamp := []byte(strconv.FormatInt(this.Timestamp, 10))
	//
	//三个[]byte进行拼接
	headers := bytes.Join([][]byte{this.PreBlockHash, this.Data, timestamp}, []byte{})
	//拼接后的[]byte进行hash运算
	//
	hash := sha256.Sum256(headers)
	//赋值给当前Hash
	this.Hash = hash[:]

}

/*
创建一个区块
*/
func NewBlock(data string, preBlockHash []byte) *Block {
	/** 生成一个区块变量 */
	block := Block{}
	/*取当前时间*/
	block.Timestamp = time.Now().Unix()
	//设计data和前一个区块的hash
	block.Data = []byte(data)
	block.PreBlockHash = preBlockHash
	//调用上面的方法生成当前区块的hash
	block.SetHash()
	return &block
}

/*
	定义区块链
	区块链=创世区块-->区块-->区块-->...........
*/
type BlockChain struct {
	//区块成链
	Blocks []*Block
}

/* 创建一个创世区块链（第一个区块） */
func NewGenesisBlock() *Block {
	genesisBlock := Block{}
	genesisBlock.Data = []byte("创世区块")
	genesisBlock.PreBlockHash = []byte{}
	genesisBlock.Timestamp = time.Now().Unix()
	genesisBlock.SetHash()
	return &genesisBlock
}

/* 新建一个区块链*/

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

//将一个区块添加到区块链中
func (this *BlockChain) AddBlock(data string) {
	//1 得到区块链中最后一个区块
	preblock := this.Blocks[len(this.Blocks)-1]
	//2 根据data新建一个区块
	newblock := NewBlock(data, preblock.Hash)
	//3 添加到区块链中
	this.Blocks = append(this.Blocks, newblock)

}
func main() {
	//
	block := NewBlock("测试时的第一个区块 ，用于验证区块", []byte{})
	fmt.Printf("测试区块的hash  %x \n", block.Hash)
	fmt.Println("测试区块的data ", string(block.Data))
	//时间布局不管什么格式数值都必须是 2006年1月2号15点4分5秒,下面是一些预定义板式,
	//在go语言标准库中有 https://studygolang.com/pkgdoc
	/*
			const (
		    ANSIC       = "Mon Jan _2 15:04:05 2006"
		    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
		    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
		    RFC822      = "02 Jan 06 15:04 MST"
		    RFC822Z     = "02 Jan 06 15:04 -0700" // 使用数字表示时区的RFC822
		    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
		    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
		    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // 使用数字表示时区的RFC1123
		    RFC3339     = "2006-01-02T15:04:05Z07:00"
		    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
		    Kitchen     = "3:04PM"
		    // 方便的时间戳
		    Stamp      = "Jan _2 15:04:05"
		    StampMilli = "Jan _2 15:04:05.000"
		    StampMicro = "Jan _2 15:04:05.000000"
		    StampNano  = "Jan _2 15:04:05.000000000"
			)
	*/
	fmt.Println("测试区块时间 ", time.Unix(block.Timestamp, 0).Format("2006-01-02 15:04:05"))
	//创建区块链
	bc := NewBlockChain()
	//用户输入
	var cmd string
	for {
		fmt.Println("按1添加区块，按2遍历区块 其他按键退出")
		fmt.Scanln(&cmd)
		switch cmd {
		case "1":
			//添加一个区块

			var input string
			fmt.Println("请输入添加的区块内容: ")
			fmt.Scanln(&input)
			bc.AddBlock(strings.TrimSpace(input))
		case "2":
			//遍历整个区块
			for i, block := range bc.Blocks {
				fmt.Println("********************************************************************************************")
				fmt.Println("* 第", i, "个区块的信息")
				fmt.Printf("* PreBlockHash : %x \n", block.PreBlockHash)
				fmt.Printf("* Data :  %s \n", block.Data) //读入数据的结尾绝逼有问题，读入了换行符，还他妈去不掉
				if i == 0 {
					fmt.Println()
				}
				fmt.Printf("* Hash :  %x \n", block.Hash)
				fmt.Printf("* Timestamp : %s \n", time.Unix(block.Timestamp, 0).Format("2006-01-02 15:04:05"))
				fmt.Println("--------------------------------------------------------------------------------------------")
			}
		default:
			fmt.Println("退出")
			return
		}
	}

}
