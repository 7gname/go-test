package zk_test

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"fmt"
	"strings"
	"path"
	"testing"
)

func final() {
	if err := recover(); err != nil {
		fmt.Printf("Last err[%v]\n", err)
	} else {
		fmt.Println("Test over!")
	}
}

func createp(cli *zk.Conn, node string) (p string, err error) {
	b := path.Dir(node)
	exists, _, err := cli.Exists(b)
	if err != nil {
		return
	}
	if !exists {
		_, err = createp(cli, b)
		if err != nil {
			return
		}
	}
	p, err = cli.Create(node, []byte(""), 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		return
	}

	return p, nil
}

func TestZkOpt(t testing.T) {
	defer final()
	zkCli, _, err := zk.Connect(strings.Split(Servers, ","), DefaultTimeout)
	if err != nil {
		panic(err)
	}
	defer zkCli.Close()

	zkPath := "/aaa/bbb"
	exists, _, err := zkCli.Exists(zkPath)
	if err != nil {
		panic(err)
	}
	if !exists {
		fmt.Printf("Path[%s] not exists\n", zkPath)
		p, err := createp(zkCli, zkPath)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Create path success[%s]\n", p)
	}

	c, s, err := zkCli.Get(zkPath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get path content[%s]\n", string(c))

	s, err = zkCli.Set(zkPath, []byte(time.Now().String()), s.Version)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Set path content success\n")

	c, s, err = zkCli.Get(zkPath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get path content again[%s]\n", string(c))

	err = zkCli.Delete(zkPath, s.Version)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Delete path success\n")

	c, s, err = zkCli.Get(zkPath)
	if err != nil {
		panic(err)
	}

	return
}

func TestZkExistsWathc(t testing.T) {
	defer final()
	zkCli, _, err := zk.Connect(strings.Split(Servers, ","), DefaultTimeout, zk.WithEventCallback(func(e zk.Event) {
		fmt.Printf("Event[%s]\n", e.Type.String())
	}))
	if err != nil {
		panic(err)
	}
	defer zkCli.Close()

	zkPath := "/test/abc"
	_, _, _, err = zkCli.ExistsW(zkPath)
	if err != nil {
		panic(err)
	}

	p, err := zkCli.Create(zkPath, []byte("hello world!"), 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Create path success[%s]\n", p)

	c, s, err := zkCli.Get(zkPath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get path content[%s]\n", string(c))

	_, _, _, err = zkCli.ExistsW(zkPath)
	if err != nil {
		panic(err)
	}

	s, err = zkCli.Set(zkPath, []byte(time.Now().String()), s.Version)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Set path content success\n")

	c, s, err = zkCli.Get(zkPath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Get path content[%s]\n", string(c))

	_, _, _, err = zkCli.ExistsW(zkPath)
	if err != nil {
		panic(err)
	}

	err = zkCli.Delete(zkPath, s.Version)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Delete path success\n")

	return
}

func watchGet(conn *zk.Conn, path string) (chan []byte, chan error) {
	snapshots := make(chan []byte)
	errors := make(chan error)

	go func() {
		for {
			snapshot, _, events, err := conn.GetW(path)
			if err != nil {
				errors <- err
				return
			}
			snapshots <- snapshot
			evt := <-events
			if evt.Err != nil {
				errors <- evt.Err
				return
			}
			fmt.Printf("event type[%s]", evt.Type.String())
		}
	}()

	return snapshots, errors
}

func TestZkGetWatch(t testing.T) {
	zkCli, _, err := zk.Connect(strings.Split(Servers, ","), DefaultTimeout)
	if err != nil {
		println(err)
	}
	defer zkCli.Close()

	zkPath := "/test/abc"
	exists, _, err := zkCli.Exists(zkPath)
	if err != nil {
		println(err)
	}
	if !exists {
		fmt.Printf("Path[%s] not exists\n", zkPath)
		p, err := zkCli.Create(zkPath, []byte("hello world!"), 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			println(err)
		}
		fmt.Printf("Create path success[%s]\n", p)
	}

	chRes, chErr := watchGet(zkCli, zkPath)

	go func() {
		for {
			select {
			case rs := <-chRes:
				fmt.Println(string(rs))
			case err := <-chErr:
				fmt.Println(err.Error())
			}
		}
	}()
	time.Sleep(time.Second)

	c, s, err := zkCli.Get(zkPath)
	if err != nil {
		println(err)
	}
	fmt.Printf("Get path content[%s]\n", string(c))
	time.Sleep(time.Second)

	s, err = zkCli.Set(zkPath, []byte(time.Now().String()), s.Version)
	if err != nil {
		println(err)
	}
	fmt.Printf("Set path content success\n")
	time.Sleep(time.Second)

	c, s, err = zkCli.Get(zkPath)
	if err != nil {
		println(err)
	}
	fmt.Printf("Get path content again[%s]\n", string(c))
	time.Sleep(time.Second)

	c, s, err = zkCli.Get(zkPath)
	if err != nil {
		println(err)
	}
	fmt.Printf("Get path content again[%s]\n", string(c))
	time.Sleep(time.Second)

	s, err = zkCli.Set(zkPath, []byte(time.Now().String()), s.Version)
	if err != nil {
		println(err)
	}
	fmt.Printf("Set path content success\n")
	time.Sleep(time.Second)

	c, s, err = zkCli.Get(zkPath)
	if err != nil {
		println(err)
	}
	fmt.Printf("Get path content again[%s]\n", string(c))
	time.Sleep(time.Second)
}

func TestZkExistWatch(t testing.T) {
	zkCli, _, err := zk.Connect(strings.Split(Servers, ","), DefaultTimeout)
	if err != nil {
		println(err)
	}
	defer zkCli.Close()

	zkPath := "/test/abc"

	event := make(chan zk.Event)
	go func() {
		for {
			_, _, e, err := zkCli.ExistsW(zkPath)
			if err != nil {
				fmt.Errorf("Watch node[%s] err[%v]", zkPath, err)
			}
			t := <-e
			event <- t
		}
	}()

	go func() {
		for {
			select {
			case e := <-event:
				fmt.Printf("Watch node[%s] event[%s]\n", e.Path, e.Type.String())
				zkCli.ExistsW(zkPath)
			default:
			}
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 5)
	p, err := zkCli.Create(zkPath, []byte("hello world!"), 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		println(err)
	}
	fmt.Printf("Create path success[%s]\n", p)
	time.Sleep(time.Second * 5)
	err = zkCli.Delete(zkPath, 0)
	if err != nil {
		println(err)
	}
	fmt.Printf("Delete path success[%s]\n", p)

	time.Sleep(time.Second * 5)
	p, err = zkCli.Create(zkPath, []byte("hello world!"), 0, zk.WorldACL(zk.PermAll))
	if err != nil {
		println(err)
	}
	fmt.Printf("Create path success[%s]\n", p)
	time.Sleep(time.Second * 5)
	err = zkCli.Delete(zkPath, 0)
	if err != nil {
		println(err)
	}
	fmt.Printf("Delete path success[%s]\n", p)
}
