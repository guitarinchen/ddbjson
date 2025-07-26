// Copyright © 2025 guitarinchen <guitarinchen@gmail.com>
// SPDX-License-Identifier: MIT
package marshall

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func avToMap(av map[string]types.AttributeValue) map[string]map[string]any {
	result := make(map[string]map[string]any)

	for k, v := range av {
		switch v := v.(type) {
		case *types.AttributeValueMemberS:
			result[k] = map[string]any{"S": v.Value}
		case *types.AttributeValueMemberN:
			result[k] = map[string]any{"N": v.Value}
		case *types.AttributeValueMemberBOOL:
			result[k] = map[string]any{"BOOL": v.Value}
		case *types.AttributeValueMemberSS:
			result[k] = map[string]any{"SS": v.Value}
		case *types.AttributeValueMemberNS:
			result[k] = map[string]any{"NS": v.Value}
		case *types.AttributeValueMemberL:
			list := make([]map[string]any, 0)
			for _, item := range v.Value {
				sub := avToMap(map[string]types.AttributeValue{"__": item})
				list = append(list, sub["__"])
			}
			result[k] = map[string]any{"L": list}
		case *types.AttributeValueMemberM:
			result[k] = map[string]any{"M": avToMap(v.Value)}
		case *types.AttributeValueMemberNULL:
			result[k] = map[string]any{"NULL": v.Value}
		}
	}

	return result
}

func Marshall(input []byte) error {
	var data map[string]any
	if err := json.Unmarshal(input, &data); err != nil {
		return err
	}

	av, err := attributevalue.MarshalMap(data)
	if err != nil {
		return err
	}

	res, err := json.Marshal(avToMap(av))
	if err != nil {
		return err
	}

	fmt.Println(string(res))

	return nil
}
