package main

import (
	"classfile"
	"classpath"
	"fmt"
	"os"
	"strings"
)

func main() {
	cmd := classpath.ParseCmd()
	if cmd.VersionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.HelpFlag {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

/**
	打印输出
 */
func printUsage() {
	fmt.Printf("Usage: %s [-options] Class [Args...]\n", os.Args[0])
}

func startJVM(cmd *classpath.Cmd) {
	//cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	//fmt.Printf("classpath:%v class:%v, args:%v \n",
	//	cp, cmd.Class, cmd.Args)
	//
	//className := strings.Replace(cmd.Class, ".", "/", -1)
	//classData, _, err := cp.ReadClass(className)
	//if err != nil {
	//	fmt.Printf("Could not find or load main class %s \n", cmd.Class)
	//	return
	//}
	//fmt.Printf("class data:%v \n", classData)

	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	cf := loadClass(className, cp)
	fmt.Println(cmd.Class)
	printClassInfo(cf)

}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf(" %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf(" %s\n", m.Name())
	}
}
