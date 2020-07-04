package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/filecoin-project/specs-actors/actors/abi"
	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/filecoin-project/specs-actors/actors/builtin/miner"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"
	"io"
)

func main() {
	var params string
	var method int
	flag.IntVar(&method, "method", 0, "method")
	flag.StringVar(&params, "params", "", "params")
	flag.Parse()
	fmt.Println("method:", method)
	fmt.Println("params:", params)
	//var p []byte
	p, _ := base64.StdEncoding.DecodeString(params)
	fmt.Println("params:", p)
	scratch := make([]byte, 8)
	br := bytes.NewReader(p)
	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println("maj:", maj)
	fmt.Println("extra", extra)
	Params := make([]byte, extra)
	if _, err := io.ReadFull(br, Params); err != nil {
		fmt.Println("err")
	}
	fmt.Println("Params", Params)
	decode(method, p)
}

func decode(method int, p []byte) (bool, error) {
	switch abi.MethodNum(method) {
	case builtin.MethodsMiner.PreCommitSector:
		fmt.Println("MethodsMiner.PreCommitSector")
		var params miner.SectorPreCommitInfo
		fmt.Println("MethodsMiner.PreCommitSector: ", p)
		if err := params.UnmarshalCBOR(bytes.NewReader(p)); err != nil {
			fmt.Println("MethodsMiner.PreCommitSector: ", err)
			return false, xerrors.Errorf("unmarshal pre commit: %w", err)
		}

		//for _, did := range params.DealIDs {
		//	//if did == abi.DealID(dealId) {
		//	//	sectorNumber = params.SectorNumber
		//	//	sectorFound = true
		//	//	return false, nil
		//	//}
		//	return false, nil
		//}

		return false, nil
	case builtin.MethodsMiner.ProveCommitSector:
		fmt.Println("MethodsMiner.ProveCommitSector")
		var params miner.ProveCommitSectorParams
		if err := params.UnmarshalCBOR(bytes.NewReader(p)); err != nil {
			fmt.Println("MethodsMiner.ProveCommitSector?")
			return false, xerrors.Errorf("failed to unmarshal prove commit sector params: %w", err)
		}

		//if !sectorFound {
		//	return false, nil
		//}

		//if params.SectorNumber != sectorNumber {
		//	return false, nil
		//}
		fmt.Println("MethodsMiner.ProveCommitSector: ", params.SectorNumber)

		return true, nil
	default:
		return false, nil
	}
}
