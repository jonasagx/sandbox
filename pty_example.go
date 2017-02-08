package main

import (
    "github.com/kr/pty"
    "io"
    "os"
    "os/exec"
)

func main() {
    // c := exec.Command("grep", "--color=auto", "amor")
    c := exec.Command("ls", "-lha")
    f, err := pty.Start(c)
    if err != nil {
        panic(err)
    }

    // go func() {
    //     f.Write([]byte("amor\n"))
    //     // f.Write([]byte("bar\n"))
    //     // f.Write([]byte("baz\n"))
    //     f.Write([]byte{4}) // EOT
    // }()
    io.Copy(os.Stdout, f)
}