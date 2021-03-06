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
		jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecEvents{}).Type1()):                     LevelSpecEventsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecEventsBadEvents{}).Type1()):            LevelSpecEventsBadEventsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecEventsGoodEvents{}).Type1()):           LevelSpecEventsGoodEventsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecEventsValidEvents{}).Type1()):          LevelSpecEventsValidEventsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecObjective{}).Type1()):                  LevelSpecObjectiveCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecObjectiveTimeWindow{}).Type1()):        LevelSpecObjectiveTimeWindowCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecObjectiveTimeWindowRolling{}).Type1()): LevelSpecObjectiveTimeWindowRollingCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecEvents{}).Type1()):                     LevelSpecEventsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecEventsBadEvents{}).Type1()):            LevelSpecEventsBadEventsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecEventsGoodEvents{}).Type1()):           LevelSpecEventsGoodEventsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecEventsValidEvents{}).Type1()):          LevelSpecEventsValidEventsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecObjective{}).Type1()):                  LevelSpecObjectiveCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecObjectiveTimeWindow{}).Type1()):        LevelSpecObjectiveTimeWindowCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecObjectiveTimeWindowRolling{}).Type1()): LevelSpecObjectiveTimeWindowRollingCodec{},
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
type LevelSpecEventsCodec struct {
}

func (LevelSpecEventsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LevelSpecEvents)(ptr) == nil
}

func (LevelSpecEventsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LevelSpecEvents)(ptr)
	var objs []LevelSpecEvents
	if obj != nil {
		objs = []LevelSpecEvents{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecEvents{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LevelSpecEventsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LevelSpecEvents)(ptr) = LevelSpecEvents{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LevelSpecEvents

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecEvents{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LevelSpecEvents)(ptr) = objs[0]
			} else {
				*(*LevelSpecEvents)(ptr) = LevelSpecEvents{}
			}
		} else {
			*(*LevelSpecEvents)(ptr) = LevelSpecEvents{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LevelSpecEvents

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecEvents{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LevelSpecEvents)(ptr) = obj
		} else {
			*(*LevelSpecEvents)(ptr) = LevelSpecEvents{}
		}
	default:
		iter.ReportError("decode LevelSpecEvents", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type LevelSpecEventsBadEventsCodec struct {
}

func (LevelSpecEventsBadEventsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LevelSpecEventsBadEvents)(ptr) == nil
}

func (LevelSpecEventsBadEventsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LevelSpecEventsBadEvents)(ptr)
	var objs []LevelSpecEventsBadEvents
	if obj != nil {
		objs = []LevelSpecEventsBadEvents{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecEventsBadEvents{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LevelSpecEventsBadEventsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LevelSpecEventsBadEvents)(ptr) = LevelSpecEventsBadEvents{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LevelSpecEventsBadEvents

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecEventsBadEvents{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LevelSpecEventsBadEvents)(ptr) = objs[0]
			} else {
				*(*LevelSpecEventsBadEvents)(ptr) = LevelSpecEventsBadEvents{}
			}
		} else {
			*(*LevelSpecEventsBadEvents)(ptr) = LevelSpecEventsBadEvents{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LevelSpecEventsBadEvents

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecEventsBadEvents{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LevelSpecEventsBadEvents)(ptr) = obj
		} else {
			*(*LevelSpecEventsBadEvents)(ptr) = LevelSpecEventsBadEvents{}
		}
	default:
		iter.ReportError("decode LevelSpecEventsBadEvents", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type LevelSpecEventsGoodEventsCodec struct {
}

func (LevelSpecEventsGoodEventsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LevelSpecEventsGoodEvents)(ptr) == nil
}

func (LevelSpecEventsGoodEventsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LevelSpecEventsGoodEvents)(ptr)
	var objs []LevelSpecEventsGoodEvents
	if obj != nil {
		objs = []LevelSpecEventsGoodEvents{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecEventsGoodEvents{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LevelSpecEventsGoodEventsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LevelSpecEventsGoodEvents)(ptr) = LevelSpecEventsGoodEvents{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LevelSpecEventsGoodEvents

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecEventsGoodEvents{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LevelSpecEventsGoodEvents)(ptr) = objs[0]
			} else {
				*(*LevelSpecEventsGoodEvents)(ptr) = LevelSpecEventsGoodEvents{}
			}
		} else {
			*(*LevelSpecEventsGoodEvents)(ptr) = LevelSpecEventsGoodEvents{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LevelSpecEventsGoodEvents

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecEventsGoodEvents{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LevelSpecEventsGoodEvents)(ptr) = obj
		} else {
			*(*LevelSpecEventsGoodEvents)(ptr) = LevelSpecEventsGoodEvents{}
		}
	default:
		iter.ReportError("decode LevelSpecEventsGoodEvents", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type LevelSpecEventsValidEventsCodec struct {
}

func (LevelSpecEventsValidEventsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LevelSpecEventsValidEvents)(ptr) == nil
}

func (LevelSpecEventsValidEventsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LevelSpecEventsValidEvents)(ptr)
	var objs []LevelSpecEventsValidEvents
	if obj != nil {
		objs = []LevelSpecEventsValidEvents{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecEventsValidEvents{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LevelSpecEventsValidEventsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LevelSpecEventsValidEvents)(ptr) = LevelSpecEventsValidEvents{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LevelSpecEventsValidEvents

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecEventsValidEvents{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LevelSpecEventsValidEvents)(ptr) = objs[0]
			} else {
				*(*LevelSpecEventsValidEvents)(ptr) = LevelSpecEventsValidEvents{}
			}
		} else {
			*(*LevelSpecEventsValidEvents)(ptr) = LevelSpecEventsValidEvents{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LevelSpecEventsValidEvents

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecEventsValidEvents{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LevelSpecEventsValidEvents)(ptr) = obj
		} else {
			*(*LevelSpecEventsValidEvents)(ptr) = LevelSpecEventsValidEvents{}
		}
	default:
		iter.ReportError("decode LevelSpecEventsValidEvents", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type LevelSpecObjectiveCodec struct {
}

func (LevelSpecObjectiveCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LevelSpecObjective)(ptr) == nil
}

func (LevelSpecObjectiveCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LevelSpecObjective)(ptr)
	var objs []LevelSpecObjective
	if obj != nil {
		objs = []LevelSpecObjective{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecObjective{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LevelSpecObjectiveCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LevelSpecObjective)(ptr) = LevelSpecObjective{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LevelSpecObjective

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecObjective{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LevelSpecObjective)(ptr) = objs[0]
			} else {
				*(*LevelSpecObjective)(ptr) = LevelSpecObjective{}
			}
		} else {
			*(*LevelSpecObjective)(ptr) = LevelSpecObjective{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LevelSpecObjective

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecObjective{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LevelSpecObjective)(ptr) = obj
		} else {
			*(*LevelSpecObjective)(ptr) = LevelSpecObjective{}
		}
	default:
		iter.ReportError("decode LevelSpecObjective", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type LevelSpecObjectiveTimeWindowCodec struct {
}

func (LevelSpecObjectiveTimeWindowCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LevelSpecObjectiveTimeWindow)(ptr) == nil
}

func (LevelSpecObjectiveTimeWindowCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LevelSpecObjectiveTimeWindow)(ptr)
	var objs []LevelSpecObjectiveTimeWindow
	if obj != nil {
		objs = []LevelSpecObjectiveTimeWindow{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecObjectiveTimeWindow{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LevelSpecObjectiveTimeWindowCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LevelSpecObjectiveTimeWindow)(ptr) = LevelSpecObjectiveTimeWindow{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LevelSpecObjectiveTimeWindow

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecObjectiveTimeWindow{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LevelSpecObjectiveTimeWindow)(ptr) = objs[0]
			} else {
				*(*LevelSpecObjectiveTimeWindow)(ptr) = LevelSpecObjectiveTimeWindow{}
			}
		} else {
			*(*LevelSpecObjectiveTimeWindow)(ptr) = LevelSpecObjectiveTimeWindow{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LevelSpecObjectiveTimeWindow

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecObjectiveTimeWindow{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LevelSpecObjectiveTimeWindow)(ptr) = obj
		} else {
			*(*LevelSpecObjectiveTimeWindow)(ptr) = LevelSpecObjectiveTimeWindow{}
		}
	default:
		iter.ReportError("decode LevelSpecObjectiveTimeWindow", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type LevelSpecObjectiveTimeWindowRollingCodec struct {
}

func (LevelSpecObjectiveTimeWindowRollingCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LevelSpecObjectiveTimeWindowRolling)(ptr) == nil
}

func (LevelSpecObjectiveTimeWindowRollingCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LevelSpecObjectiveTimeWindowRolling)(ptr)
	var objs []LevelSpecObjectiveTimeWindowRolling
	if obj != nil {
		objs = []LevelSpecObjectiveTimeWindowRolling{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecObjectiveTimeWindowRolling{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LevelSpecObjectiveTimeWindowRollingCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LevelSpecObjectiveTimeWindowRolling)(ptr) = LevelSpecObjectiveTimeWindowRolling{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LevelSpecObjectiveTimeWindowRolling

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecObjectiveTimeWindowRolling{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LevelSpecObjectiveTimeWindowRolling)(ptr) = objs[0]
			} else {
				*(*LevelSpecObjectiveTimeWindowRolling)(ptr) = LevelSpecObjectiveTimeWindowRolling{}
			}
		} else {
			*(*LevelSpecObjectiveTimeWindowRolling)(ptr) = LevelSpecObjectiveTimeWindowRolling{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LevelSpecObjectiveTimeWindowRolling

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LevelSpecObjectiveTimeWindowRolling{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LevelSpecObjectiveTimeWindowRolling)(ptr) = obj
		} else {
			*(*LevelSpecObjectiveTimeWindowRolling)(ptr) = LevelSpecObjectiveTimeWindowRolling{}
		}
	default:
		iter.ReportError("decode LevelSpecObjectiveTimeWindowRolling", "unexpected JSON type")
	}
}
