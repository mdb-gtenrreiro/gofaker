package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleJSON_object() {
	Seed(11)

	value, err := JSON(&JSONOptions{
		Type: "object",
		Fields: []Field{
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "address", Function: "address"},
			{Name: "password", Function: "password", Params: map[string][]string{"special": {"false"}}},
		},
		Indent: true,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: {
	//     "first_name": "Markus",
	//     "last_name": "Moen",
	//     "address": {
	//         "address": "4599 Dale ton, Lake Carroll, Mississippi 90635",
	//         "street": "4599 Dale ton",
	//         "city": "Lake Carroll",
	//         "state": "Mississippi",
	//         "zip": "90635",
	//         "country": "Saint Pierre and Miquelon",
	//         "latitude": 22.008873,
	//         "longitude": 158.531956
	//     },
	//     "password": "YjJbXclnVN0H"
	// }
}

func ExampleJSON_array() {
	Seed(11)

	value, err := JSON(&JSONOptions{
		Type: "array",
		Fields: []Field{
			{Name: "id", Function: "autoincrement"},
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "password", Function: "password", Params: map[string][]string{"special": {"false"}}},
		},
		RowCount: 3,
		Indent:   true,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: [
	//     {
	//         "id": 1,
	//         "first_name": "Markus",
	//         "last_name": "Moen",
	//         "password": "Dc0VYXjkWABx"
	//     },
	//     {
	//         "id": 2,
	//         "first_name": "Osborne",
	//         "last_name": "Hilll",
	//         "password": "XPJ9OVNbs5lm"
	//     },
	//     {
	//         "id": 3,
	//         "first_name": "Mertie",
	//         "last_name": "Halvorson",
	//         "password": "eyl3bhwfV8wA"
	//     }
	// ]
}

func TestJSONLookup(t *testing.T) {
	info := GetFuncLookup("json")

	m := map[string][]string{
		"type":     {"array"},
		"rowcount": {"10"},
		"fields": {
			`{"name":"id","function":"autoincrement"}`,
			`{"name":"first_name","function":"firstname"}`,
			`{"name":"password","function":"password","params":{"special":["false"],"length":["20"]}}`,
		},
	}
	_, err := info.Call(&m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	// t.Fatal(fmt.Sprintf("%s", value.([]byte)))
}

func BenchmarkJSONLookup100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("json")
		m := map[string][]string{
			"type":     {"array"},
			"rowcount": {"100"},
			"fields": {
				`{"name":"id","function":"autoincrement"}`,
				`{"name":"first_name","function":"firstname"}`,
				`{"name":"last_name","function":"lastname"}`,
				`{"name":"password","function":"password"}`,
				`{"name":"description","function":"paragraph"}`,
				`{"name":"created_at","function":"date"}`,
			},
		}
		_, err := info.Call(&m, info)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}

func BenchmarkJSONLookup1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("json")
		m := map[string][]string{
			"type":     {"array"},
			"rowcount": {"1000"},
			"fields": {
				`{"name":"id","function":"autoincrement"}`,
				`{"name":"first_name","function":"firstname"}`,
				`{"name":"last_name","function":"lastname"}`,
				`{"name":"password","function":"password"}`,
				`{"name":"description","function":"paragraph"}`,
				`{"name":"created_at","function":"date"}`,
			},
		}
		_, err := info.Call(&m, info)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}

func BenchmarkJSONLookup10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("json")
		m := map[string][]string{
			"type":     {"array"},
			"rowcount": {"10000"},
			"fields": {
				`{"name":"id","function":"autoincrement"}`,
				`{"name":"first_name","function":"firstname"}`,
				`{"name":"last_name","function":"lastname"}`,
				`{"name":"password","function":"password"}`,
				`{"name":"description","function":"paragraph"}`,
				`{"name":"created_at","function":"date"}`,
			},
		}
		_, err := info.Call(&m, info)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}

func BenchmarkJSONLookup100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("json")
		m := map[string][]string{
			"type":     {"array"},
			"rowcount": {"100000"},
			"fields": {
				`{"name":"id","function":"autoincrement"}`,
				`{"name":"first_name","function":"firstname"}`,
				`{"name":"last_name","function":"lastname"}`,
				`{"name":"password","function":"password"}`,
				`{"name":"description","function":"paragraph"}`,
				`{"name":"created_at","function":"date"}`,
			},
		}
		_, err := info.Call(&m, info)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}
