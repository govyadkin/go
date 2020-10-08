package unique

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

func FindUnique(input []string, options Options) (output []string, err error) {
	duplicateFrequency := make(map[string]*Data)
	var origin string
	for _, item := range input {
		origin = item
		if options.Fields {
			itemFields := strings.Fields(item)
			if options.NumFields > len(itemFields) {
				return output, errors.New("incorrect parameters")
			}
			item = strings.Join(itemFields[options.NumFields:], " ")
		}
		if options.Chars {
			if options.NumFields > len(item) {
				return output, errors.New("incorrect parameters")
			}
			item = item[options.NumChars:]
		}
		if options.Register {
			item = strings.ToLower(item)
		}
		_, exist := duplicateFrequency[item]
		if exist {
			duplicateFrequency[item].Count += 1
		} else {
			data := Data{
				Origin: origin,
				Count:  1,
			}
			duplicateFrequency[item] = &data
		}
	}

	for _, v := range duplicateFrequency {
		if !options.Duplicate && !options.Unique || options.Unique && options.Duplicate {
			if options.Count {
				output = append(output, strconv.Itoa(v.Count)+" "+v.Origin)
			} else {
				output = append(output, v.Origin)
			}
			continue
		}
		if options.Unique {
			if v.Count > 1 {
				continue
			} else {
				if options.Count {
					output = append(output, strconv.Itoa(v.Count)+" "+v.Origin)
				} else {
					output = append(output, v.Origin)
				}
			}
		}
		if options.Duplicate {
			if v.Count <= 1 {
				continue
			} else {
				if options.Count {
					output = append(output, strconv.Itoa(v.Count)+" "+v.Origin)
				} else {
					output = append(output, v.Origin)
				}
			}
		}
	}

	sort.Strings(output)
	return output, nil
}
