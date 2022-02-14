package configstore_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-tk/configstore"
)

func Example() {
	// 1. Create configuration files for testing.
	_ = os.Mkdir("./temp", 0755)

	ioutil.WriteFile("./temp/foo.yaml", []byte(`
user_id: 1000
nickname: roy
`), 0644)

	ioutil.WriteFile("./temp/bar.yaml", []byte(`
secrets:
  password: s0g00d
  luck_numbers:
    - 1
    - 3
    - 5
`), 0644)

	// 2. Override configuration values with environment variables.
	os.Setenv("CONFIGSTORE.foo.nickname", "lisa")             // env value should be valid YAML
	os.Setenv("CONFIGSTORE.bar.secrets.luck_numbers.1", "99") // env value should be valid YAML

	// 3. Read in configuration files.
	configstore.MustOpen("./temp")

	// 4. Dump the whole configuration in form of JSON for debugging.
	data, _ := json.MarshalIndent(configstore.Cache(), "", "  ")
	fmt.Println("===== Cache =====")
	fmt.Printf("%s\n", data)

	// 5. Load a configuration item by a path.
	var secrets struct {
		Password    string `json:"password"`     // should use json tag rather than yaml tag
		LuckNumbers []int  `json:"luck_numbers"` // should use json tag rather than yaml tag
	}
	configstore.MustLoadItem("bar.secrets", &secrets)
	fmt.Println("===== MustLoadItem  =====")
	fmt.Printf("%v\n", secrets)

	// output:
	// ===== Cache =====
	// {
	//   "bar": {
	//     "secrets": {
	//       "luck_numbers": [
	//         1,
	//         99,
	//         5
	//       ],
	//       "password": "s0g00d"
	//     }
	//   },
	//   "foo": {
	//     "nickname": "lisa",
	//     "user_id": 1000
	//   }
	// }
	// ===== MustLoadItem  =====
	// {s0g00d [1 99 5]}
}
