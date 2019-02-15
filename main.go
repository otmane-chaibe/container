package main
import(
    "fmt"
    "os"
    "os/exec"

)

// go run main.go run cmd args
func main(){
    swich os.Args[i]{
        case "run"
            run()
        case "child"
            child()
        default:
            panic("what??")
    }

}
func run(){
    
    cmd:= exec.Command("/proc/self/exe", append([] string{"child"}os.Args[2]...),os.Args[3:]...)
    cmd.Stdin = os.stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Cloneflags: syscall.Clone_NEWUTS | syscall.Clone_NEWPID,
    }
    must(cmd.Run())
}

func child(){
    fmt.printf("running %v as PID %d/n", os.args[2:], os.Getpid())
    cmd:= exec.Command(os.Args[2],os.Args[3:]...)
    cmd.Stdin = os.stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
   
    must(syscall.Chroot("/home/jdjdj"))
    must(os.Chdir("/"))
    must(syscall.Mount("proc","proc","proc", 0,""))
    must(cmd.Run())
}
func must(err error){
    if err != nil{
        panic(err)
    }
}
