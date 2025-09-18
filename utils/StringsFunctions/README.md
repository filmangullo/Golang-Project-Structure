# `/utils/StringsFunctions`

### Introduction
Go includes a variety of functions for manipulating string values. 
Many of these functions are used by the Go standard library itself;
however, you are free to use them in your own applications if you find them convenient.


- The **`StringsFunctions.After`** method returns everything after the given value in a string. The entire string will be returned if the value does not exist within the string:
- The **`StringsFunctions.Int8CoalescePositive`** method ensure the minimum value is 1 (or def result) is int8:
- The **`StringsFunctions.Int16CoalescePositive`** method ensure the minimum value is 1 (or def result) is int16:
- The **`StringsFunctions.Int32CoalescePositive`** method ensure the minimum value is 1 (or def result) is int32:
- The **`StringsFunctions.Int64CoalescePositive`** method ensure the minimum value is 1 (or def result) is int64:
- The **`StringsFunctions.IntCoalescePositive`** method ensure the minimum value is 1 (or def result) is int:
- The **`StringsFunctions.IsInteger`** method whether the given value is of an integer type such as int, int8 - int64, uint, etc:
- The **`StringsFunctions.IsSlug`** method determines whether the given string is a valid Slug:
- The **`StringsFunctions.Lower`** method converts the given string to lowercase:
- The **`StringsFunctions.Slug`** method method generates a URL friendly "slug" from the given string:
- The **`StringsFunctions.StringPointer`** method converts the given string into a *string:
- The **`StringsFunctions.ToInt8`** method converts the given string to int8:
- The **`StringsFunctions.ToInt16`** method converts the given string to int16:
- The **`StringsFunctions.ToInt32`** method converts the given string to int32:
- The **`StringsFunctions.ToInt64`** method converts the given string to int64:
- The **`StringsFunctions.ToInt`** method converts the given string to int: