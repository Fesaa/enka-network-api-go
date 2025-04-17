package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"os"
	"path"
	"strconv"
)

type LocalFileExtUrl struct {
	name string
	b    []byte
}

func (src *LocalFileExtUrl) LoadMethod() (stmt *ast.BlockStmt, importSpec []ast.Spec) {
	return &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						ast.NewIdent("res"),
						ast.NewIdent("err"),
					},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   ast.NewIdent("http"),
								Sel: ast.NewIdent("Get"),
							},
							Args: []ast.Expr{
								&ast.BasicLit{
									Kind:  token.STRING,
									Value: fmt.Sprintf(`"https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/%s.json"`, src.name),
								},
							},
						},
					},
				},
				&ast.IfStmt{
					Cond: &ast.BinaryExpr{
						X:  ast.NewIdent("err"),
						Op: token.NEQ,
						Y:  ast.NewIdent("nil"),
					},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ReturnStmt{
								Results: []ast.Expr{ast.NewIdent("err")},
							},
						},
					},
				},
				&ast.DeferStmt{
					Call: &ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent("res.Body"),
							Sel: ast.NewIdent("Close"),
						},
						Args: []ast.Expr{},
					},
				},
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						ast.NewIdent("data"),
						ast.NewIdent("err"),
					},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   ast.NewIdent("io"),
								Sel: ast.NewIdent("ReadAll"),
							},
							Args: []ast.Expr{
								ast.NewIdent("res.Body"),
							},
						},
					},
				},
				&ast.IfStmt{
					Cond: &ast.BinaryExpr{
						X:  ast.NewIdent("err"),
						Op: token.NEQ,
						Y:  ast.NewIdent("nil"),
					},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ReturnStmt{
								Results: []ast.Expr{ast.NewIdent("err")},
							},
						},
					},
				},
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   ast.NewIdent("json"),
								Sel: ast.NewIdent("Unmarshal"),
							},
							Args: []ast.Expr{
								ast.NewIdent("data"),
								&ast.UnaryExpr{
									Op: token.AND,
									X: &ast.SelectorExpr{
										X:   ast.NewIdent("a"),
										Sel: ast.NewIdent("_data"),
									},
								},
							},
						},
					},
				},
			},
		}, []ast.Spec{
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: strconv.Quote("net/http"),
				},
			},
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: strconv.Quote("io"),
				},
			},
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: strconv.Quote("encoding/json"),
				},
			},
		}
}

func (src *LocalFileExtUrl) Json() ([]byte, error) {
	if src.b != nil {
		return src.b, nil
	}

	var file *os.File
	var err error

	file, err = os.Open(path.Join("..", "jsons", src.name+".json"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	src.b, err = io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return src.b, nil
}

func (src *LocalFileExtUrl) Name() string {
	return src.name
}
