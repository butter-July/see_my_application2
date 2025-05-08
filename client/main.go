package main

import (
	"golang.org/x/sys/windows" //提供windows系统调用的绑定,包括访问DLL和系统A
	"log"
	"net/http"
	"os"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

type information struct {
	Username string
	usingApp string //需要定义一个结构体来记录用户的名字以及用户正在使用的软件,最后返回这两个
}

// GetForegroundWindow function
var (
	//那么mod就可以连接DLL了
	user32           = windows.NewLazyDLL("user32.dll") //使用NewLazyDLL来访问user32.dll因为GetforegroundWindows函数要求DLL-User32.dll(NewLazyDLL他的作用是实现对单个DLL的访问)
	getWindowText    = user32.NewProc("GetWindowTextW") //获取名字(最上面的页眉)
	foregroundWindow = user32.NewProc("GetForegroundWindow")
	//使addr时可以进行查找操作(查询的是:过程的地址:就像是窗口的地址)//也是为获取窗口标题做准备的一步
	GetWindowsTextLength = user32.NewProc("GetWindowTextLengthW")
)

type (
	HANDLE uintptr //
	HWND   HANDLE  //返回值getforegroundWindows的
)

func GetWindowTextLengthW(hwnd uintptr) int {
	length, _, _ := GetWindowsTextLength.Call(hwnd)
	return int(length)

}

func GetForegroundWindow() (hwnd HWND) {
	r0, _, _ := syscall.SyscallN(foregroundWindow.Addr(), 0, 0, 0, 0) //需要五个参数传进去,返回三个但是只需要第一个所以r0,_,_
	hwnd = HWND(r0)                                                   //把ro的类型转化为hwnd类型
	return hwnd                                                       //最后返回hwnd

}

func GetWindowTextW(hwnd uintptr, nMaxCount int) string { //在获取了hwnd之后可以使用这个函数来获取窗口标题,参数有三个
	//Length, _, _ := GetWindowsTextLength.Call(uintptr(hwnd)) //获取一个不大于15的标题长度,之后可以用来找标题是什么
	textlength := GetWindowTextLengthW(uintptr(hwnd)) + 1
	LpString := make([]uint16, textlength)
	length, _, _ := getWindowText.Call(
		hwnd,
		uintptr(unsafe.Pointer(&LpString[0])),
		uintptr(nMaxCount),
	)

	return syscall.UTF16ToString(LpString[:length])
}

func main() {
	for {
		time.Sleep(1 * time.Second)
		foregroundWindow := GetForegroundWindow()
		foregroundwindowtextlength := GetWindowTextLengthW(uintptr(foregroundWindow))
		foregroundwindowtext := GetWindowTextW(uintptr(foregroundWindow), foregroundwindowtextlength)
		log.Println(foregroundwindowtext)
		http.Post(os.Getenv("UNTITLED_SERVER_URL"), "text/plain", strings.NewReader(foregroundwindowtext))
	}
}
