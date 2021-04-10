package playground

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"golang.org/x/sync/errgroup"
)

func M() {
	pctx := context.Background()
	ctx, cancel := context.WithCancel(pctx)
	go w(ctx)
	cancel()
	for {
	}
}
func OP() {
	ctx := context.Background()
	g, c := errgroup.WithContext(ctx)
	errs := make(chan error)

	g.Go(func() (err error) {
		time.Sleep(time.Second * 3)
		return fmt.Errorf("wrong")
	})

	g.Go(func() (err error) {
		for {
		}
	})
	go func() {
		err := g.Wait()
		if err != nil {
			errs <- err
		}
	}()

	for {
		select {
		case <-errs:
			fmt.Println("goroutine returned with error")
			return
		case <-c.Done():
			fmt.Println("hmmmmmm")
		default:
			fmt.Println("still at it")
		}
	}
}

func w(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return nil
		default:
			time.Sleep(time.Second * 5)
			fmt.Println("working")
		}
	}
}

func UI() {

	slice_2 := []string{"FGHJ", "FGH", "gfg", "KJHGF"}

	var f1, f2, f3 string
	f1 = "GEEKs"
	f2 = "C"
	f3 = "gfg"

	sort.Strings(slice_2)

	fmt.Println("Slice 2: ", slice_2)

	res1 := sort.SearchStrings(slice_2, f1)
	res2 := sort.SearchStrings(slice_2, f2)
	res3 := sort.SearchStrings(slice_2, f3)

	fmt.Println(res1, res2, res3)

}

func IUGTFV() {
	sss := time.Now().UnixNano() / 1000000
	fmt.Println(sss)
	tm := time.Unix(sss, 0)
	fmt.Println(tm)
	tm2 := time.Unix(sss+1, 0)
	fmt.Println(sss + 1)
	fmt.Println(tm2)
}

var (
	iface    = "wlo1"
	snaplen  = int32(1600)
	promisc  = false
	timeout  = pcap.BlockForever
	filter   = "tcp and port 80"
	devFound = false
)

func SOKOQAJqi() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panicln(err)
	}

	for _, device := range devices {

		if device.Name == iface {
			devFound = true
		}
	}

	if !devFound {
		log.Panicf("Device named '%s' does not exist\n", iface)
	}

	handle, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
	if err != nil {
		log.Panicln(err)
	}
	defer handle.Close()

	if err := handle.SetBPFFilter(filter); err != nil {
		log.Panicln(err)
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range source.Packets() {
		app := packet.ApplicationLayer()
		if app != nil {
			payload := app.Payload()
			// dst := packet.NetworkLayer().NetworkFlow().Dst()
			if bytes.Contains(payload, []byte("USER")) {
				fmt.Print(string(payload))
			} else if bytes.Contains(payload, []byte("PASS")) {
				fmt.Print(string(payload))
			}
		}
	}

}

type compute struct {
	name      string
	age       int
	isAwesome bool
}

func set1(c *compute) {
	c.age = 22
}
func set2(c *compute) {
	c.name = "amogh"
}
func set3(c *compute) {
	c.isAwesome = true
}

func ComputeThings() {
	s := &compute{}

	go set1(s)
	go set2(s)
	go set3(s)

	time.Sleep(time.Second * 3)

	fmt.Println(s.name, s.age, s.isAwesome)
}
