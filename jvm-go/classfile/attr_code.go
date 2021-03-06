package classfile

/**
变长属性，只存在于method_info结构中，存放字节码等方法相关信息
*/
type CodeAttribute struct {
	cp             ConstantPool           // 常量池
	maxStack       uint16                 // 操作数栈最大深度
	maxLocals      uint16                 // 局部变量表大小
	code           []byte                 //
	exceptionTable []*ExceptionTableEntry // 异常表
	attributes     []AttributeInfo        // 属性表
}

func (self *CodeAttribute) Code() []byte {
	return self.code
}

func (self *CodeAttribute) MaxStack() uint {
	return uint(self.maxStack)
}

func (self *CodeAttribute) MaxLocals() uint {
	return uint(self.maxLocals)
}



type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}
