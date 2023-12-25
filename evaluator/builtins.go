package evaluator

import (
	"fmt"
	"os"
	"github.com/KoMaTop10/Monkey/object" 
)


var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got = %d, want = 1",len(args))
			}

			switch arg := args[0].(type) { 
			case *object.String: 
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s",args[0].Type())
			}
		},
	},

	"first": {
		Fn: func (args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got = %d. want = 1",len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s",args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},

	"last": {
		Fn: func (args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got = %d. want = 1",len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s",args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if len(arr.Elements) > 0 {
				return arr.Elements[length - 1]
			}

			return NULL
		},
	},

	"rest": {
		Fn: func (args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("rwong number of arguments.got = %d, want = 1",len(args))
			}
			
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object,length - 1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},

	"push": {
		Fn: func (args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("rwong number of arguments.got = %d, want = 2",len(args))
			}
			
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object,length+1)
				copy(newElements, arr.Elements)
				newElements[length] = args[1]
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},

	"puts": {
		Fn: func (args ...object.Object)object.Object {
			for _,arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},

	"read": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got = %d, want = 1",len(args))
			}

			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `first` must be String, got %s",args[0].Type())
			}

			f, err := os.Open(args[0].Inspect())

			data := make([]byte, 1024)
			count, err := f.Read(data)
			if err != nil {
				fmt.Println(err)
				fmt.Println("fail to read file")
				return NULL
			}

			return &object.String{Value: string(data[:count])}
		},
	},

	"write": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got = %d. want = 2",len(args))
			}

			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `first` must be String. got %s", args[0].Type())
			}
			
			if args[1].Type() != object.STRING_OBJ {
				return newError("argument to `second` must be String. got %s", args[0].Type())
			}

			f, err := os.Create(args[0].Inspect())
			
			write_data := []byte(args[1].Inspect())
			count, err := f.Write(write_data)

			if err != nil {
				fmt.Println(err)
				fmt.Println("fail to write file")
			}

			fmt.Printf("write %d bytes\n", count)

			return NULL
		},
	},
}


