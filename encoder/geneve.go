package encoder

import (
	"github.com/dreadl0ck/netcap/types"
	"github.com/golang/protobuf/proto"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var geneveEncoder = CreateLayerEncoder(types.Type_NC_Geneve, layers.LayerTypeGeneve, func(layer gopacket.Layer, timestamp string) proto.Message {
	if geneve, ok := layer.(*layers.Geneve); ok {
		var opts = []*types.GeneveOption{}
		if len(geneve.Options) > 0 {
			for _, o := range geneve.Options {
				opts = append(opts, &types.GeneveOption{
					Class:  int32(o.Class),
					Type:   int32(o.Type),
					Flags:  int32(o.Flags),
					Length: int32(o.Length),
					Data:   o.Data,
				})
			}
		}
		return &types.Geneve{
			Timestamp:      timestamp,
			Version:        int32(geneve.Version),
			OptionsLength:  int32(geneve.OptionsLength),
			OAMPacket:      bool(geneve.OAMPacket),
			CriticalOption: bool(geneve.CriticalOption),
			Protocol:       int32(geneve.Protocol),
			VNI:            uint32(geneve.VNI),
			Options:        opts,
		}
	}
	return nil
})
