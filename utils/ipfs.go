package utils

import (
	"bytes"
	"context"
	"io/ioutil"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/ipfs/go-ipfs-api/options"
)

type Ipfs struct {
	Url string
	Sh  *shell.Shell
}

func NewIpfs(url string) *Ipfs {
	sh := shell.NewShell(url)
	return &Ipfs{
		Url: url,
		Sh:  sh,
	}
}

func (i *Ipfs) Add(str string, isPin bool) (cid string, err error) {
	cid, err = i.Sh.Add(bytes.NewBufferString(str), shell.Pin(isPin))
	if err != nil {
		return
	}
	return
}

func (i *Ipfs) Cat(cid string) (string, error) {
	read, err := i.Sh.Cat(cid)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(read)

	return string(body), nil
}

func (i *Ipfs) Get(cid, outdir string) (err error) {
	err = i.Sh.Get(cid, outdir)
	if err != nil {
		return
	}

	return nil
}

// pin start
func (i *Ipfs) Pin(cid string) (err error) {
	err = i.Sh.Pin(cid)
	if err != nil {
		return
	}
	return nil
}

func (i *Ipfs) Pins(pinType string) (infs map[string]shell.PinInfo, err error) {
	if pinType == "" {
		infs, err = i.Sh.Pins()
	} else {
		infs, err = i.Sh.PinsOfType(context.Background(), shell.PinType(pinType))
	}

	if err != nil {
		return
	}
	return infs, nil
}

func (i *Ipfs) UnPin(cid string) (err error) {
	err = i.Sh.Unpin(cid)
	if err != nil {
		return
	}

	err = i.Sh.Request("repo/gc", cid).
		Option("recursive", true).
		Exec(context.Background(), nil)
	if err != nil {
		return
	}

	return nil
}

// pin end

// block start
func (i *Ipfs) BlockGet(cid string) (content []byte, err error) {
	content, err = i.Sh.BlockGet(cid)
	if err != nil {
		return
	}
	return content, nil
}

func (i *Ipfs) BlockPut(block []byte, format, mhtype string, mhlen int) (cid string, err error) {
	cid, err = i.Sh.BlockPut(block, format, mhtype, mhlen)
	if err != nil {
		return
	}
	return cid, nil
}

func (i *Ipfs) BlockStat(cid string) (key string, size int, err error) {
	key, size, err = i.Sh.BlockStat(cid)
	if err != nil {
		return
	}
	return key, size, nil
}

// block end

// dag start
func (i *Ipfs) DagGet(path string) (out interface{}, err error) {
	err = i.Sh.DagGet(path, &out)
	if err != nil {
		return
	}
	return out, nil
}

func (i *Ipfs) DagPut(content interface{}, opts ...options.DagPutOption) (cid string, err error) {
	cid, err = i.Sh.DagPutWithOpts(content, opts...)
	if err != nil {
		return
	}
	return cid, nil
}

// dag end

// key start
func (i *Ipfs) KeyGen(name string) (key *shell.Key, err error) {
	key, err = i.Sh.KeyGen(context.Background(), name)
	if err != nil {
		return
	}
	return key, nil
}

// key end

// name start
func (i *Ipfs) NamePublish(key, contentHash string, lifetime int, resolve bool) (resp *shell.PublishResponse, err error) {
	resp, err = i.Sh.PublishWithDetails(contentHash, key, time.Duration(lifetime)*time.Hour, 0, resolve)
	if err != nil {
		return
	}
	return resp, nil
}

// name end
