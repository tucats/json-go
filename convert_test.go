package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		base    string
		want    string
		wantErr bool
	}{
		{
			name: "simple array",
			arg:  `[1, 2]`,
			want: "[]int",
		},
		{
			name: "mixed array",
			arg:  `[1, true]`,
			want: "[]interface{}",
		},
		{
			name: "simple struct",
			arg:  `{ "age": 63, "married":true}`,
			base: "item",
			want: "type item struct {\n  Age     int  `json:\"age,omitempty\"`\n  Married bool `json:\"married,omitempty\"`\n}\n",
		},
		{
			name: "simple bool",
			arg:  `false`,
			want: "bool",
		},
		{
			name: "simple int",
			arg:  `134`,
			want: "int",
		},
		{
			name: "simple string",
			arg:  `"Tom"`,
			want: "string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Convert([]byte(tt.arg), tt.base, false)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriver(t *testing.T) {
	type t1 struct {
		Age  int
		Name string
	}

	type t2 struct {
		Team []t1
		Read bool
	}

	v1 := t2{
		Team: []t1{
			{
				Age:  63,
				Name: "Tom",
			},
			{
				Age:  22,
				Name: "Debbie",
			},
		},
		Read: false,
	}

	j, err := json.Marshal(v1)
	if err != nil {
		t.Error(err)
	}

	g, _ := Convert(j, "item", false)

	fmt.Println(g)
}
