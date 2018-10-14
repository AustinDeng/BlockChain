# BlockChain

基于 Go 构建一个区块链 Demo

## 安装

    // 获取项目源码
    $ cd $GOPATH/src
    $ git clone git@github.com:AustinDeng/BlockChain.git
    

## 运行

    // 进入目录
    $ cd $GOPATH/src/BlockChain
    
    // 执行命令
    $ go run ./cmd/main.go
    
    Index: 0
    PrevBlockHash:
    CurrHash: 0d8845eb2da42f75aef4ee920f644975d73347e0331d17b37209c4f32ef4867f
    Data: Geneesis Block
    Timestamp: 1539490069
    
    Index: 1
    PrevBlockHash: 0d8845eb2da42f75aef4ee920f644975d73347e0331d17b37209c4f32ef4867f
    CurrHash: b11305449703848e79f02f0ba7f7db6bdd085a4a5ea50382ea4cca77644c376b
    Data: This is the first block.
    Timestamp: 1539490069
    
    Index: 2
    PrevBlockHash: b11305449703848e79f02f0ba7f7db6bdd085a4a5ea50382ea4cca77644c376b
    CurrHash: 751c3793ee3492f5e050c6b662f4d832bc125dde0aae813147e5459abc23f29a
    Data: This is the second block.
    Timestamp: 1539490069

## 通过 RPC 接口访问数据

### 启动服务

    // 进入目录
    $ cd $GOPATH/src/BlockChain
    
    // 开启http服务监听
    $ go run rpc/Server.go
    

### 查看数据

打开浏览器输入 URL 地址：`http://localhost:8000/blockchain/get`

    {
        "Block": [
            {
                "Index": 0,
                "TimeStamp": 1539490490,
                "PrevBlockHash": "",
                "Hash": "0d8845eb2da42f75aef4ee920f644975d73347e0331d17b37209c4f32ef4867f",
                "Data": "Genesis Block"
            }
        ]
    }

### 写入区块链数据

打开浏览器输入 URL 地址：`http://localhost:8000/blockchain/write?data=This is the first block.`

    {
        "Block": [
            {
                "Index": 0,
                "TimeStamp": 1539490490,
                "PrevBlockHash": "",
                "Hash": "0d8845eb2da42f75aef4ee920f644975d73347e0331d17b37209c4f32ef4867f",
                "Data": "Genesis Block"
            },
            {
                "Index": 1,
                "TimeStamp": 1539490635,
                "PrevBlockHash": "0d8845eb2da42f75aef4ee920f644975d73347e0331d17b37209c4f32ef4867f",
                "Hash": "b11305449703848e79f02f0ba7f7db6bdd085a4a5ea50382ea4cca77644c376b",
                "Data": "This is the first block."
            }
        ]
    }

## 项目部分源码

### 区块结构体定义
    
    type Block struct {
    	Index int64
    	TimeStamp int64
    	PrevBlockHash string
    	Hash string
    	Data string
    }
    
### 区块链结构体定义
    
    type BlockChain struct {
	    Block []*Block
    }

### 计算区块的 Hash 值

    func calculateHash(block Block) string {
    	data := string(block.Index) + string(block.TimeStamp) + block.PrevBlockHash + block.Data
    	blockInBytes := sha256.Sum256([]byte(data))
    	return hex.EncodeToString(blockInBytes[:])
    }
