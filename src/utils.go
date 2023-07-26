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

func fetchBinary(info BinaryInfo) error {
    tar_path := fmt.Sprintf("%v%v.tar.gz", info.path, info.binary_name)
    binary_path := fmt.Sprintf("%v%v", info.path, info.binary_name)

    // Create the file
    os.MkdirAll(info.path, 0755)
    tar_file, err := os.Create(tar_path)
    if err != nil  {
        return err
    }
    defer tar_file.Close()

    // Get the data
    resp, err := http.Get(info.url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Check server response
    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("bad status: %s", resp.Status)
    }

    // Writer the body to file
    _, err = io.Copy(tar_file, resp.Body)
    if err != nil  {
        return err
    }
    // Extract tar
    err = unTarGZ(tar_path, info.path)
    if err != nil{
        log.Fatalf("unTarGz failed with %s\n",err)
    }

    // Allow binary to be executed
    os.Chmod(binary_path, 0755)

    // Debuggings
    cmd := exec.Command(fmt.Sprintf("%v",binary_path),"version")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

    //Return nil on success
    return nil
}

func fetchGenesis(info JSONInfo) error {
    return nil
}

func unTarGZ(tar_path, path string) error {
    gzipStream , err := os.Open(tar_path)
    uncompressedStream, err := gzip.NewReader(gzipStream)
    if err != nil {
        log.Fatal("ExtractTarGz: NewReader failed")
    }

    tarReader := tar.NewReader(uncompressedStream)

    // For loop is golang's while loop too, for with no conditionals is like while True:
    for {
        header, err := tarReader.Next()
        // Stops at the end of the tar
        if err == io.EOF {
            break
        }
        // Breaks on all other errors
        if err != nil {
            log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
        }
        // Switch case for each file type in tar (directory or regular file)
        switch header.Typeflag {
            case tar.TypeDir:
                if err := os.Mkdir(header.Name, 0755); err != nil {
                    log.Fatalf("ExtractTarGz: Mkdir() failed: %s", err.Error())
                }
            case tar.TypeReg:
                outFile, err := os.Create(fmt.Sprintf("%v%v",path,header.Name))
                if err != nil {
                    log.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
                }
                if _, err := io.Copy(outFile, tarReader); err != nil {
                    log.Fatalf("ExtractTarGz: Copy() failed: %s", err.Error())
                }
                outFile.Close()

            default:
                log.Fatalf(
                    "ExtractTarGz: uknown type in %s",
                    // header.Typeflag,
                    header.Name)
            }

    }

    return nil
}
