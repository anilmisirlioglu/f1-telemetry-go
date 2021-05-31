# F1 Game Telemetry Client in Go [![Made With Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg?color=007EC6)](http://golang.org)

Telemetry client for F1 Game, written in Go. Currently, supported only the UDP 2020 format.

- Event System
- Rich Env Constants
- UDP Stats (recv, err and packet per second rate),
- Vector3 support

## Install
```bash
go get -u github.com/anilmisirlioglu/f1-telemetry-go
```

## Quick Start
```go
func main() {
	client, err := telemetry.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	
    client.OnEventPacket(func(packet *packets.PacketEventData) {
        fmt.Printf("Code: %s\n", packet.EventCodeString())
    })
	
	client.Run()
}
```

## Docs

If you need more information on the F1 UDP specifications, see the [docs](/docs).