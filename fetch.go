package main

import (
    "fmt"
    "time"
    "os"
    "os/user"
    
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/host"
    
    "github.com/ukrfetch/ukrfetch/flagmtx"
)

const MSG = "We express unity for UKRAINE and their peace."
var HASH = []string{"#StandWithUkraine", "#WeStandWithUkraine"}

const ESC_RESET          = "\033[0m"
const ESC_FONT_GREEN     = "\033[32m"
const ESC_FONT_YELLOW    = "\033[33m"
const ESC_FONT_BLUE      = "\033[34m"
const ESC_FONT_CYAN      = "\033[36m"
const ESC_BG_BLUE        = "\033[44m"

const DRAWAREA_H = 18       // height of drawing area
const SYSTEM_INFO_X = 52    // position(left-top) of system info table

func greeting(t time.Time) string{
    
    hour := t.Hour()
    
    if hour >= 4 && hour <= 10 {
        return "Good morning"    
    }else if hour >= 11 && hour <= 16{
        return "Hello"
    }else if hour >= 17 && hour <= 21{
        return "Good evening"
    }else{
        return "Good night"
    }
    
}

func makeSystemInfoTable() []string{
    
    infoList := make([]string, 0, 16)
    
	// get system info
    user, err := user.Current()
    checkError(err)
	hostInfo, err := host.Info()
	checkError(err)
	cpuInfo, err := cpu.Info()
	checkError(err)
	v_mem, err := mem.VirtualMemory()
	checkError(err)
	
	bootTime := time.Unix(int64(hostInfo.BootTime), 0)
	OSInfo := fmt.Sprintf("%s %s %s", hostInfo.Platform, hostInfo.PlatformVersion, hostInfo.KernelArch)
	memUsage := fmt.Sprintf("%d MiB / %d MiB", (v_mem.Used / 1000 / 1000), (v_mem.Total / 1000 / 1000))
	
	infoList = append(infoList, ESC_FONT_BLUE + user.Username + "@" + hostInfo.Hostname + ESC_RESET)
	infoList = append(infoList, "-------------------------")
	infoList = append(infoList, ESC_FONT_YELLOW + "OS"          + ESC_RESET + ": " + OSInfo)
	infoList = append(infoList, ESC_FONT_YELLOW + "Kernel"      + ESC_RESET + ": " + hostInfo.KernelVersion)
	infoList = append(infoList, ESC_FONT_YELLOW + "BootTime"    + ESC_RESET + ": " + bootTime.Format(time.UnixDate))
	infoList = append(infoList, ESC_FONT_YELLOW + "CPU"         + ESC_RESET + ": " + cpuInfo[0].ModelName)
	infoList = append(infoList, ESC_FONT_YELLOW + "Memory"      + ESC_RESET + ": " + memUsage)
	
	// margin
	infoList = append(infoList, "")
	
	nowTime := time.Now()
	
	// get UKR times
    EET := time.FixedZone("EET", 2*60*60)
    EEST := time.FixedZone("EEST", 3*60*60)
    
	infoList = append(infoList, ESC_FONT_CYAN+"[current time(UKR)]"+ESC_RESET)
    infoList = append(infoList, nowTime.In(EET).Format(time.UnixDate))
    infoList = append(infoList, nowTime.In(EEST).Format(time.UnixDate) + "(summer time)")
    infoList = append(infoList, ESC_FONT_GREEN + "-> " + greeting(nowTime.In(EET)) + ", Kyiv!" + ESC_RESET)
    
	// margin
	infoList = append(infoList, "")
	
	// get local Time
	infoList = append(infoList, ESC_FONT_CYAN+"[current time(Local)]"+ESC_RESET)
    infoList = append(infoList, nowTime.Format(time.UnixDate))
    greetingStr := fmt.Sprintf("-> %s, %s!", greeting(nowTime), user.Username)
    infoList = append(infoList, ESC_FONT_GREEN + greetingStr + ESC_RESET)
	
    return infoList
}

func drawFlag_row(row *[]string) string{
    
    rowStr := ""
    
    for _, px := range (* row) {
        //fmt.Printf("index: %d, name: %s\n", i, s)
        if len(px) == 2{
            color := byte(px[0])
            char := byte(px[1])
            if color == 'b' {
                //putBlue(char)
                rowStr += ESC_FONT_BLUE + string(char) + ESC_RESET
            }else if color == 'y' {
                //putYellow(char)
                rowStr += ESC_FONT_YELLOW + string(char) + ESC_RESET
            }
        }
    }
    
    return rowStr
    
}

func draw(flgMtx *[][]string, infoList *[]string) {
    
    // drawing
    for i := 0; i < DRAWAREA_H; i++ {
        
        idx := 0
        
        // magin
        fmt.Printf(" ")
        idx += 1
        
        // draw flag(each row)
        if(i < len((*flgMtx))){
            fmt.Printf(drawFlag_row(&(*flgMtx)[i]))
            idx += len((*flgMtx)[i])
        }
        
        // draw time
        
        // offset
        offsetStr := ""
        for i := 0; i < SYSTEM_INFO_X - idx; i++{
            offsetStr += " "
        }
        fmt.Printf(offsetStr)
        
        // draw system info
        if(i < len((* infoList))){
            fmt.Printf((* infoList)[i])
        }
        
        
        fmt.Println()
    }
    
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}

func main() {
    
    // margin
    fmt.Println()
    
    // get your system info
    infoList := makeSystemInfoTable()
    
    // draw contents
    draw(&flagmtx.FLAG_MTX, &infoList)
    
    // draw msg
    fmt.Printf(" ")
    fmt.Println(MSG)
    fmt.Printf(" ")
    for _, hash := range(HASH){fmt.Printf("%s%s%s ",ESC_BG_BLUE,hash,ESC_RESET)}

    // margin
    fmt.Println()
    fmt.Println()
    
}