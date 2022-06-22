package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fmt.Print(
		countAsyncMethods("C:\\Users\\pnedo\\GolandProjects\\Golang_Level_2\\Lesson_7\\someFunc.go", "Kind"))
}

func getCallExprLiteral(c *ast.CallExpr) string {
	q, ok := c.Fun.(*ast.Ident)
	if ok {
		return q.Name
	}
	s, ok := c.Fun.(*ast.SelectorExpr)
	if ok {
		return s.Sel.Name
	}
	return ""
}

func countAsyncMethods(fileName string, funcName string) (int, error) {
	count := 0
	fset := token.NewFileSet()
	packs, err := parser.ParseFile(fset, fileName, nil, 0)
	if err != nil {
		fmt.Println("Failed to parse package:", err)
		return 0, err
	}

	var funcs []*ast.FuncDecl
	for _, d := range packs.Decls {
		if fn, isFn := d.(*ast.FuncDecl); isFn {
			funcs = append(funcs, fn)
		}
	}
	for _, fun := range funcs {
		ast.Inspect(fun, func(node ast.Node) bool {
			switch n := node.(type) {
			case *ast.CallExpr:
				if getCallExprLiteral(n) == funcName {
					count++
				}
			}
			return true
		})
	}

	return count, err
}
