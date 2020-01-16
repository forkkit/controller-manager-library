/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// taken from controller runtime project

package resources

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/json"
)

// Decoder knows how to decode the contents of an admission
// request into a concrete object.
type Decoder struct {
	codecs serializer.CodecFactory
}

// NewDecoder creates a Decoder given the runtime.Scheme
func NewDecoder(scheme *runtime.Scheme) *Decoder {
	return &Decoder{codecs: serializer.NewCodecFactory(scheme)}
}

// Decode decodes on object given as byte stream into a runtimeObject or
// similar Object
func (d *Decoder) Decode(data []byte, into interface{}) error {
	switch target := into.(type) {
	case *unstructured.Unstructured:
		// unmarshal into unstructured's underlying object to avoid calling the decoder
		if err := json.Unmarshal(data, &target.Object); err != nil {
			return err
		}
		return nil
	case runtime.Object:
		deserializer := d.codecs.UniversalDeserializer()
		return runtime.DecodeInto(deserializer, data, target)
	default:
		if err := json.Unmarshal(data, &target); err != nil {
			return err
		}
		return nil
	}
}

// DecodeFromUnstructued decodes an unstruvtured object into a structured one
func (d *Decoder) DecodeFromUnstructued(data *unstructured.Unstructured, into runtime.Object) error {
	return d.DecodeFromMap(data.Object, into)
}

// DecodeFromMap decodes from a map into a runtime Object.
// data is a JSON compatible map with string, float, int, bool, []interface{}, or
// map[string]interface{}
// children.
func (d *Decoder) DecodeFromMap(data map[string]interface{}, into runtime.Object) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return d.Decode(bytes, into)
}

// DecodeRaw decodes a RawExtension object into the passed-in runtime.Object.
func (d *Decoder) DecodeRaw(rawObj runtime.RawExtension, into interface{}) error {
	return d.Decode(rawObj.Raw, into)
}
