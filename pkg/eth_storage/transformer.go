// VulcanizeDB
// Copyright © 2022 Vulcanize

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package eth_storage

import (
	"fmt"

	"github.com/vulcanize/migration-tools/pkg/interfaces"
)

// Transformer struct for transforming v2 DB eth.storage_cids models to v2 DB models
type Transformer struct {
}

// NewTransformer satisfies interfaces.TransformerConstructor for eth.storage_cids
func NewTransformer() interfaces.Transformer {
	return &Transformer{}
}

// Transform satisfies interfaces.Transformer for eth.storage_cids
func (t *Transformer) Transform(models interface{}, expectedRange [2]uint64) (interface{}, [][2]uint64, error) {
	v2Models, ok := models.(*[]StorageModelV2WithMeta)
	if !ok {
		return nil, [][2]uint64{expectedRange}, fmt.Errorf("expected models of type %T, got %T", new([]StorageModelV2WithMeta), models)
	}
	v3Models := make([]StorageModelV3, len(*v2Models))
	for i, model := range *v2Models {
		v3Models[i] = StorageModelV3{
			HeaderID:   model.BlockHash,
			StatePath:  model.StatePath,
			Path:       model.Path,
			StorageKey: model.StorageKey,
			NodeType:   model.NodeType,
			CID:        model.CID,
			MhKey:      model.MhKey,
			Diff:       model.Diff,
		}
	}
	return v3Models, nil, nil
}
