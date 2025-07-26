// Copyright © 2025 guitarinchen <guitarinchen@gmail.com>
// SPDX-License-Identifier: MIT
package unmarshall

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

func Unmarshall(input []byte) error {
	av, err := attributevalue.UnmarshalMapJSON(input)
	if err != nil {
		return err
	}

	var out any
	if err := attributevalue.UnmarshalMap(av, &out); err != nil {
		return err
	}

	res, err := json.Marshal(out)
	if err != nil {
		return err
	}

	fmt.Println(string(res))

	return nil
}
