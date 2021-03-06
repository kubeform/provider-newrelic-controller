/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func GetEncoder() map[string]jsoniter.ValEncoder {
	return map[string]jsoniter.ValEncoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ChannelSpecConfig{}).Type1()):       ChannelSpecConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MutingRuleSpecCondition{}).Type1()): MutingRuleSpecConditionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MutingRuleSpecSchedule{}).Type1()):  MutingRuleSpecScheduleCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ChannelSpecConfig{}).Type1()):       ChannelSpecConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MutingRuleSpecCondition{}).Type1()): MutingRuleSpecConditionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MutingRuleSpecSchedule{}).Type1()):  MutingRuleSpecScheduleCodec{},
	}
}

func getEncodersWithout(typ string) map[string]jsoniter.ValEncoder {
	origMap := GetEncoder()
	delete(origMap, typ)
	return origMap
}

func getDecodersWithout(typ string) map[string]jsoniter.ValDecoder {
	origMap := GetDecoder()
	delete(origMap, typ)
	return origMap
}

// +k8s:deepcopy-gen=false
type ChannelSpecConfigCodec struct {
}

func (ChannelSpecConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ChannelSpecConfig)(ptr) == nil
}

func (ChannelSpecConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ChannelSpecConfig)(ptr)
	var objs []ChannelSpecConfig
	if obj != nil {
		objs = []ChannelSpecConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ChannelSpecConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ChannelSpecConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ChannelSpecConfig)(ptr) = ChannelSpecConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ChannelSpecConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ChannelSpecConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ChannelSpecConfig)(ptr) = objs[0]
			} else {
				*(*ChannelSpecConfig)(ptr) = ChannelSpecConfig{}
			}
		} else {
			*(*ChannelSpecConfig)(ptr) = ChannelSpecConfig{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ChannelSpecConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ChannelSpecConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ChannelSpecConfig)(ptr) = obj
		} else {
			*(*ChannelSpecConfig)(ptr) = ChannelSpecConfig{}
		}
	default:
		iter.ReportError("decode ChannelSpecConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type MutingRuleSpecConditionCodec struct {
}

func (MutingRuleSpecConditionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*MutingRuleSpecCondition)(ptr) == nil
}

func (MutingRuleSpecConditionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*MutingRuleSpecCondition)(ptr)
	var objs []MutingRuleSpecCondition
	if obj != nil {
		objs = []MutingRuleSpecCondition{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MutingRuleSpecCondition{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (MutingRuleSpecConditionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*MutingRuleSpecCondition)(ptr) = MutingRuleSpecCondition{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []MutingRuleSpecCondition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MutingRuleSpecCondition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*MutingRuleSpecCondition)(ptr) = objs[0]
			} else {
				*(*MutingRuleSpecCondition)(ptr) = MutingRuleSpecCondition{}
			}
		} else {
			*(*MutingRuleSpecCondition)(ptr) = MutingRuleSpecCondition{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj MutingRuleSpecCondition

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MutingRuleSpecCondition{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*MutingRuleSpecCondition)(ptr) = obj
		} else {
			*(*MutingRuleSpecCondition)(ptr) = MutingRuleSpecCondition{}
		}
	default:
		iter.ReportError("decode MutingRuleSpecCondition", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type MutingRuleSpecScheduleCodec struct {
}

func (MutingRuleSpecScheduleCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*MutingRuleSpecSchedule)(ptr) == nil
}

func (MutingRuleSpecScheduleCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*MutingRuleSpecSchedule)(ptr)
	var objs []MutingRuleSpecSchedule
	if obj != nil {
		objs = []MutingRuleSpecSchedule{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MutingRuleSpecSchedule{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (MutingRuleSpecScheduleCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*MutingRuleSpecSchedule)(ptr) = MutingRuleSpecSchedule{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []MutingRuleSpecSchedule

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MutingRuleSpecSchedule{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*MutingRuleSpecSchedule)(ptr) = objs[0]
			} else {
				*(*MutingRuleSpecSchedule)(ptr) = MutingRuleSpecSchedule{}
			}
		} else {
			*(*MutingRuleSpecSchedule)(ptr) = MutingRuleSpecSchedule{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj MutingRuleSpecSchedule

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MutingRuleSpecSchedule{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*MutingRuleSpecSchedule)(ptr) = obj
		} else {
			*(*MutingRuleSpecSchedule)(ptr) = MutingRuleSpecSchedule{}
		}
	default:
		iter.ReportError("decode MutingRuleSpecSchedule", "unexpected JSON type")
	}
}
