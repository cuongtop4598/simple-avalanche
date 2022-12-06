package core

func GetLastestBlock() *Block {
	if len(Blockchain) >= 1 {
		return &Blockchain[len(Blockchain)-1]
	} else {
		return nil
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
