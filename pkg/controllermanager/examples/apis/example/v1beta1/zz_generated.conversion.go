// +build !ignore_autogenerated

/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	example "github.com/gardener/controller-manager-library/pkg/controllermanager/examples/apis/example"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Example)(nil), (*example.Example)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Example_To_example_Example(a.(*Example), b.(*example.Example), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.Example)(nil), (*Example)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_Example_To_v1beta1_Example(a.(*example.Example), b.(*Example), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ExampleList)(nil), (*example.ExampleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ExampleList_To_example_ExampleList(a.(*ExampleList), b.(*example.ExampleList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.ExampleList)(nil), (*ExampleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_ExampleList_To_v1beta1_ExampleList(a.(*example.ExampleList), b.(*ExampleList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ExampleSpec)(nil), (*example.ExampleSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ExampleSpec_To_example_ExampleSpec(a.(*ExampleSpec), b.(*example.ExampleSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*example.ExampleSpec)(nil), (*ExampleSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_example_ExampleSpec_To_v1beta1_ExampleSpec(a.(*example.ExampleSpec), b.(*ExampleSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_Example_To_example_Example(in *Example, out *example.Example, s conversion.Scope) error {
	if err := Convert_v1beta1_ExampleSpec_To_example_ExampleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_Example_To_example_Example is an autogenerated conversion function.
func Convert_v1beta1_Example_To_example_Example(in *Example, out *example.Example, s conversion.Scope) error {
	return autoConvert_v1beta1_Example_To_example_Example(in, out, s)
}

func autoConvert_example_Example_To_v1beta1_Example(in *example.Example, out *Example, s conversion.Scope) error {
	if err := Convert_example_ExampleSpec_To_v1beta1_ExampleSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_example_Example_To_v1beta1_Example is an autogenerated conversion function.
func Convert_example_Example_To_v1beta1_Example(in *example.Example, out *Example, s conversion.Scope) error {
	return autoConvert_example_Example_To_v1beta1_Example(in, out, s)
}

func autoConvert_v1beta1_ExampleList_To_example_ExampleList(in *ExampleList, out *example.ExampleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]example.Example)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_ExampleList_To_example_ExampleList is an autogenerated conversion function.
func Convert_v1beta1_ExampleList_To_example_ExampleList(in *ExampleList, out *example.ExampleList, s conversion.Scope) error {
	return autoConvert_v1beta1_ExampleList_To_example_ExampleList(in, out, s)
}

func autoConvert_example_ExampleList_To_v1beta1_ExampleList(in *example.ExampleList, out *ExampleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Example)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_example_ExampleList_To_v1beta1_ExampleList is an autogenerated conversion function.
func Convert_example_ExampleList_To_v1beta1_ExampleList(in *example.ExampleList, out *ExampleList, s conversion.Scope) error {
	return autoConvert_example_ExampleList_To_v1beta1_ExampleList(in, out, s)
}

func autoConvert_v1beta1_ExampleSpec_To_example_ExampleSpec(in *ExampleSpec, out *example.ExampleSpec, s conversion.Scope) error {
	out.URL = in.URL
	return nil
}

// Convert_v1beta1_ExampleSpec_To_example_ExampleSpec is an autogenerated conversion function.
func Convert_v1beta1_ExampleSpec_To_example_ExampleSpec(in *ExampleSpec, out *example.ExampleSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_ExampleSpec_To_example_ExampleSpec(in, out, s)
}

func autoConvert_example_ExampleSpec_To_v1beta1_ExampleSpec(in *example.ExampleSpec, out *ExampleSpec, s conversion.Scope) error {
	out.URL = in.URL
	return nil
}

// Convert_example_ExampleSpec_To_v1beta1_ExampleSpec is an autogenerated conversion function.
func Convert_example_ExampleSpec_To_v1beta1_ExampleSpec(in *example.ExampleSpec, out *ExampleSpec, s conversion.Scope) error {
	return autoConvert_example_ExampleSpec_To_v1beta1_ExampleSpec(in, out, s)
}