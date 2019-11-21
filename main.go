package main

import (
	"encoding/json"
	"fmt"
)

type T struct{}
type Person struct {
	Name      string   `json:"name"`       // `json:"-"`でnameを-で隠す
	Age       int      `json:"age,string"` // unmarshal時はintにして、marshal時はstringにする指定。,をつける。ageが0の時はkey?になり、見せないようにする `json:"age,omitempty"``
	Nicknames []string `json:"nicknames"`
	T         *T       `json:"T,omitempty"` // 実態を指定する*T
}

// omitempty・・・ ポインタの時はnil, sliceは[]string, intのときは0, stringは空のときに何も入っていない時表示させないのがomitempty

// json.Marshalが呼ばれた時独自のカスタマイズをする方法
// func (p Person) MarshalJSON() ([]byte, error) { //MarshalJSONの名前でなければダメ
// 	// a := &struct{ Name string }{Name: "test"}
// 	v, err := json.Marshal(&struct { // 上記のPersonのようなstructを関数内だけでつ買う場合、短絡的に書く方法
// 		Name string
// 	}{
// 		Name: "Mr ." + p.Name,
// 	})
// 	return v, err
// }

// Unmarshalしたさいに呼ばれるカスタムUnmarshal
func (p *Person) UnmarshalJSON(b []byte) error {
	type Person2 struct {
		Name string
	}
	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	p.Name = p2.Name + "!"
	return err
}

func main() {
	b := []byte(`{ "name": "mike", "age": "20", "nicknames": ["a", "b", "c"]}`)
	// b := []byte(`{ "Name": "mike", "ge": 20, "nicknames": ["a", "b", "c"]}`) // パブリックなのでロワーケースでもキャピタルでも入れてくれるが geはないのでこの場合 デフォルト値がはいる
	var p Person
	if err := json.Unmarshal(b, &p); err != nil { // ネットワークから得たデータをPerson Structのkeyをみて入れてくれるUnmarshal。参照先のアドレスと、データを渡す
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	// 以下 Personに入っているデータをJSONに戻してネットワーク越しに投げる
	v, _ := json.Marshal(p) // vはこの時点ではbyte
	fmt.Println(string(v))  // string　コンバージョンしてあげる // {"Name":"mike","Age":20,"Nicknames":["a","b","c"]}
	// このままだとNameなどのkeyがstructと同じキャピタルのままになってしまっているので、Marshalする時はリテラルで指定してあげる -> `json:"name"`
	// {"name":"mike","age":20,"nicknames":["a","b","c"]}
}
