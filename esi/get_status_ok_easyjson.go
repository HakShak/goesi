// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC17a1f4DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetStatusOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetStatusOkList, 0, 1)
			} else {
				*out = GetStatusOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetStatusOk
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC17a1f4EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetStatusOkList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GetStatusOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC17a1f4EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetStatusOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC17a1f4EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetStatusOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC17a1f4DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetStatusOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC17a1f4DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonC17a1f4DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetStatusOk) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "players":
			out.Players = int32(in.Int32())
		case "server_version":
			out.ServerVersion = string(in.String())
		case "start_time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.StartTime).UnmarshalJSON(data))
			}
		case "vip":
			out.Vip = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC17a1f4EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetStatusOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Players != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"players\":")
		out.Int32(int32(in.Players))
	}
	if in.ServerVersion != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"server_version\":")
		out.String(string(in.ServerVersion))
	}
	if true {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"start_time\":")
		out.Raw((in.StartTime).MarshalJSON())
	}
	if in.Vip {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"vip\":")
		out.Bool(bool(in.Vip))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetStatusOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC17a1f4EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetStatusOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC17a1f4EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetStatusOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC17a1f4DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetStatusOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC17a1f4DecodeGithubComAntihaxGoesiEsi1(l, v)
}