# pscan.go
A Golang script to check if common ports are open on a server or device.

## Dependencies
pscan.go is written in Golang so it is **REQUIRED**.

A working 1.10.2 [Go](https://golang.org/dl/) environment.

## Usage

Upon executing the program input a server address or IP.

Example:
  ```
  www.hackthissite.org
  ```

* The program will then begin to scan from the common ports slice.

* The slice can be altered yourself if you require more ports to be scanned.
* The slice was just used for examples.
    ```
    common_ports := [num_of_slice_items]int{ports}
    ```


## Download Options -- Installing
Currently you can only clone or download the project ZIP file. (Recommended clone if you're going to be contributing.)

Extract and navigate to the zipfile directory and run pscan.go by executing the main entry point file (pscan.go):
  ```
  go run pscan.go
  ```

## Authors -- Contributors

* **Daniel Brennand** - *Author* - [Dextroz](https://github.com/Dextroz)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) for details.
