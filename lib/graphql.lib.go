package lib

import (
	"fmt"
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

func GetSelectedFields(params graphql.ResolveParams) (map[string]interface{}, error) {
	fieldASTs := params.Info.FieldASTs

	if len(fieldASTs) == 0 {
		return nil, fmt.Errorf("getSelectedFields: ResolveParams has no fields")
	}
	return selectedFieldsFromSelections(params, fieldASTs[0].SelectionSet.Selections)
}

func selectedFieldsFromSelections(params graphql.ResolveParams, selections []ast.Selection) (selected map[string]interface{}, err error) {
	selected = map[string]interface{}{}

	for _, s := range selections {
		switch s := s.(type) {
		case *ast.Field:
			if s.SelectionSet == nil {
				selected[s.Name.Value] = true
			} else {
				selected[s.Name.Value], err = selectedFieldsFromSelections(params, s.SelectionSet.Selections)
				if err != nil {
					return
				}
			}
		case *ast.FragmentSpread:
			n := s.Name.Value
			frag, ok := params.Info.Fragments[n]
			if !ok {
				err = fmt.Errorf("getSelectedFields: no fragment found with name %v", n)

				return
			}

			selected[s.Name.Value], err = selectedFieldsFromSelections(params, frag.GetSelectionSet().Selections)

			if err != nil {
				return
			}
		default:
			err = fmt.Errorf("getSelectedFields: found unexpected selection type %v", s)

			return
		}
	}
	return
}

func GetMappedArgs(p graphql.ResolveParams) (map[string]interface{}, error) {
	fieldASTs := p.Info.FieldASTs
	return getArgsFromQuery(p, fieldASTs[0].SelectionSet.Selections)
}

func IsFieldExist(param string, params graphql.ResolveParams) bool {
	fieldMap, _ := GetSelectedFields(params)

	//check if string contains dot (.)
	containsDot := strings.Contains(param, ".")

	if containsDot {
		sliceDot := strings.Split(param, ".")
		//spew.Dump(sliceDot)
		if len(sliceDot) == 2 {
			if _, ok := fieldMap[sliceDot[0]]; ok {
				slice1 := fieldMap[sliceDot[0]].(map[string]interface{})
				if _, ok2 := slice1[sliceDot[1]]; ok2 {
					return true
				}
			}
		}
	} else {
		if _, ok := fieldMap[param]; ok {
			return true
		}
	}

	return false
}

func getArgsFromQuery(params graphql.ResolveParams, selections []ast.Selection) (args map[string]interface{}, err error) {
	args = map[string]interface{}{}

	for _, l1 := range selections {
		switch l1 := l1.(type) {
		case *ast.Field:

			arg := map[string]interface{}{}
			for _, c1 := range l1.Arguments {

				arg[c1.Name.Value] = c1.Value.GetValue()
				for _, l2 := range l1.SelectionSet.Selections {
					switch l2 := l2.(type) {
					case *ast.Field:
						arg2 := map[string]interface{}{}
						for _, d := range l2.Arguments {
							arg2[d.Name.Value] = d.Value.GetValue()
						}

						if _, ok := arg[l2.Name.Value]; !ok {
							arg[l2.Name.Value] = arg2
						}
					}
				}

			}

			args[l1.Name.Value] = arg
		}
	}
	return args, err
}
