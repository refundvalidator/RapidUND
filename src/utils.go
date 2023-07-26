package main

import (
    "fmt"
    "os"
    "os/exec"
    "net/http"
    "io"
    "log"
    "archive/tar"
    "compress/gzip"
)

func getBinary(name, path, url string) (err error) {
    // Create the file
    out, err := os.Create(fmt.Sprintf("%v%v.tar.gz",path,name))
    if err != nil  {
        return err
    }
    defer out.Close()

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Check server response
    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("bad status: %s", resp.Status)
    }

    // Writer the body to file
    _, err = io.Copy(out, resp.Body)
    if err != nil  {
        return err
    }
    extractTarGz(fmt.Sprintf("%v%v.tar.gz", path,name))
    os.Chmod(fmt.Sprintf("%v%v",path,name), 0755)

    cmd := exec.Command(fmt.Sprintf("%v%v", path, name),"version")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

    return nil
}

func extractTarGz(path string) {
    r, err := os.Open(path)
    if err != nil {
        fmt.Println("error")
    }
    uncompressedStream, err := gzip.NewReader(r)
    if err != nil {
        log.Fatal("ExtractTarGz: NewReader failed")
    }

    tarReader := tar.NewReader(uncompressedStream)

    for true {
        header, err := tarReader.Next()

        if err == io.EOF {
            break
        }

        if err != nil {
            log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
        }

        switch header.Typeflag {
        case tar.TypeDir:
            if err := os.Mkdir(header.Name, 0755); err != nil {
                log.Fatalf("ExtractTarGz: Mkdir() failed: %s", err.Error())
            }
        case tar.TypeReg:
            outFile, err := os.Create(header.Name)
            if err != nil {
                log.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
            }
            if _, err := io.Copy(outFile, tarReader); err != nil {
                log.Fatalf("ExtractTarGz: Copy() failed: %s", err.Error())
            }
            outFile.Close()

        default:
            log.Fatalf(
                "ExtractTarGz: uknown type: %s in %s",
                header.Typeflag,
                header.Name)
        }

    }
}

