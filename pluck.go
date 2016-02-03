/*
The MIT License (MIT)

Copyright (c) 2016 Ravi Kant

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
        "errors"
        "fmt"
        "reflect"

        "github.com/fatih/structs"
)

// Function pluck is used to retrieve an array of subset of fields(branch) present in original structure(plant).
// Input : 'plant' is the source from which a branch needs to be plucked. An array of structure is expected.
//         'branch' is the output structure(subset). This should not be an array as this will be used just to
//              form the output structure.
// Output : []map[string]interface{} - An array of map is returned. See example for more details.
func pluck(plant interface{}, branch interface{}) ([]map[string]interface{}, error) {

        flag := 0
        // Read the value from interface{}
        srcExtract := reflect.ValueOf(plant)

        // For branch, only format is needed(and the input is not an array)
        // The value is extracted and converted to map.
        destExtract := reflect.ValueOf(branch).Interface()
        destValMap := structs.Map(destExtract)

        // The result map[string]interface{} to be returned
        branchExtract := make([]map[string]interface{}, srcExtract.Len())

        // Retrieve the source elements one by one and copy to dest
        for i := 0; i < srcExtract.Len(); i++ {
                indexVal := srcExtract.Index(i).Interface()
                indexValMap := structs.Map(indexVal)
                // Create a temp variable to hold the trimmed values
                destTempMap := make(map[string]interface{})
                for key := range destValMap {
                        if value, present := indexValMap[key]; present {
                                destTempMap[key] = value
                                flag = 1
                        }
                }
                // append the temp var into output
                branchExtract[i] = destTempMap
        }
                // This is to make sure that at least one value got extracted from plant to branch
        if flag == 0 {
                err := errors.New("Source Destination Type Mismatch")
                return nil, err
        }
        return branchExtract, nil
}

// Function pluckElement is used to retrieve an array of just one field(destKeyName) present in original structure(plant).
// Input : 'plant' is the source from which an element needs to be plucked. An array of structure is expected.
//         'destKeyName' is the output element key name. This should not be an array as this will be used just to
//              form the output structure.
// Output : []interface{} - An array is returned. Type assertion can be used to derive an array of required type.
// See example for more details.
func pluckElement(plant interface{}, destKeyName string) ([]interface{}, error) {

        flag := 0
        // Read the value from interface{}
        srcExtract := reflect.ValueOf(plant)

        // The result map[string]interface{} to be returned
        var elementExtract []interface{}

        // Retrieve the source elements one by one and copy to dest
        for i := 0; i < srcExtract.Len(); i++ {
                indexVal := srcExtract.Index(i).Interface()
                indexValMap := structs.Map(indexVal)
                if value, present := indexValMap[destKeyName]; present {
                        elementExtract = append(elementExtract, value)
                        flag = 1
                }
        }

        // This is to make sure that at least one value got extracted from plant to branch
        if flag == 0 {
                err := errors.New("Source Destination Type Mismatch")
                return nil, err
        }
        return elementExtract, nil
}
func main() {

        first := []struct {
                Name     string
                City     string
                Position string
        }{
                {
                        Name:     "Tripti Rani",
                        City:     "Ranchi",
                        Position: "Software Engineer",
                },
                {
                        Name:     "Ravi Kant",
                        City:     "Gurgaon",
                        Position: "Software Engineer",
                },
                {
                        Name:     "Daljeet Singh",
                        City:     "Gurgaon",
                        Position: "Application Developer",
                },
        }
        var second struct {
                Name string
                City string
        }
        result, _ := pluck(first, second)
        fmt.Println("Result after pluck\n", result)
        names, _ := pluckElement(first, "Name")
        fmt.Println("Names after pluck\n", names)
}
