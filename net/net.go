package net

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"runtime"
	"time"

	"github.com/charmbracelet/log"
	scriptish "github.com/ganbarodigital/go_scriptish"
	"github.com/jackpal/gateway"
	"github.com/rwxrob/to"
)

var (
	currentIP                  string
	currentNetworkHardwareName string
)

// wanSearch makes an HTTP GET request to ipify.,org and returns
// the current public IP address. Currently unable to remove the
// trailing % sign.
func wanSearch() {
	res, _ := http.Get("http://icanhazip.com")
	ip, _ := io.ReadAll(res.Body)
	curIP := to.String(ip)
	fmt.Printf("%s\n", curIP)
}

// lanSearch uses the net package to query the system for network
// interfaces. Returns the IPv4 addresses of all found
// interfaces.
func lanSearch() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Error(err)
	}

	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		// = GET LOCAL IP ADDRESS
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip := ipnet.IP.String()
				fmt.Printf("%s\n", ip)
			}
		}
	}
}

func routerSearch() {
	gateway, err := gateway.DiscoverGateway()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s\n", gateway.String())
	}
}

func dnsSearch() {

	// TODO: Find a cross-platform solution rather than switching on the GOOS

	os := runtime.GOOS
	switch os {
	case "windows":
		msg := "Unsupported operating system, exiting..."
		log.Fatal(msg)
	case "linux":
		result, err := scriptish.NewPipeline(
			scriptish.Exec("cat", "/etc/resolv.conf"),
			scriptish.Grep("(^nameserver)"),
			scriptish.CutFields("2"),
		).Exec().String()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", result)
	case "darwin":
		result, err := scriptish.NewPipeline(
			scriptish.Exec("grep", "-i", "nameserver", "/etc/resolv.conf"),
			scriptish.Exec("head", "-n1"),
			scriptish.CutFields("2"),
		).Exec().TrimmedString()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", result)
	}
}

// Uses nmcli on Linux/Mac systems to get the current wifi password.
// On Windows, uses netsh.
// Requires the SSID of the current network to be entered as a parameter
// on the command line.
// $ c get net wifipass [SSID]
// Can be aliased in the main Cmd for the monolith.
func wifiPasswd(network string) {
	osType := runtime.GOOS
	switch osType {
	// Windows case is untested, may not work as intended
	case "windows":
		if network == "" {
			msg := "Missing SSID. Exiting..."
			log.Fatal(msg)
		} else {
			result, err := scriptish.NewPipeline(
				scriptish.Exec("netsh", "wlan", "show", "profile name=", network, "key=clear"),
				scriptish.ToStdout(),
			).Exec().String()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", result)
		}
	case "linux", "darwin":
		result, err := scriptish.NewPipeline(
			scriptish.Exec("nmcli", "device", "wifi", "show-password"),
			scriptish.ToStdout(),
		).Exec().String()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", result)
	}
}

func macSearch() {
	fmt.Printf("%16.16X\n", macUint64())
}

func macUint64() uint64 {
	interfaces, err := net.Interfaces()
	if err != nil {
		return uint64(0)
	}

	for _, i := range interfaces {
		if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {

			// Skip locally administered addresses
			if i.HardwareAddr[0]&2 == 2 {
				continue
			}

			var mac uint64
			for j, b := range i.HardwareAddr {
				if j >= 8 {
					break
				}
				mac <<= 8
				mac += uint64(b)
			}

			return mac
		}
	}

	return uint64(0)
}

func checkNetwork() {
	hostName := "google.com"
	portNum := "80"
	seconds := 5
	timeOut := time.Duration(seconds) * time.Second

	conn, err := net.DialTimeout("tcp", hostName+":"+portNum, timeOut)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Checking %s via %s.......... Connection valid.\n", hostName, conn.LocalAddr().String())
}
