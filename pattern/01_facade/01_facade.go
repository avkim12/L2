package main

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type TextEditor struct {
}

func (te TextEditor) CreateCode() {
	fmt.Println("Написание кода")
}

func (te TextEditor) Save() {
	fmt.Println("Сохранение")
}

type Compiler struct {
}

func (c Compiler) Compile() {
	fmt.Println("Компиляция")
}

type CLR struct {
}

func (clr CLR) Execute() {
	fmt.Println("Выполнение")
}

func (clr CLR) Finish() {
	fmt.Println("Завершение работы")
}

type VisualStudioFacade struct {
	textEditor TextEditor
	compiler   Compiler
	clr        CLR
}

func NewVisualStudioFacade(textEditor TextEditor, compiler Compiler, clr CLR) *VisualStudioFacade {
	return &VisualStudioFacade{
		textEditor: textEditor,
		compiler:   compiler,
		clr:        clr,
	}
}

func (vsf *VisualStudioFacade) Start() {
	vsf.textEditor.CreateCode()
	vsf.textEditor.Save()
	vsf.compiler.Compile()
	vsf.clr.Execute()
}

func (vsf *VisualStudioFacade) Stop() {
	vsf.clr.Finish()
}

func main() {

	te := TextEditor{}
	c := Compiler{}
	clr := CLR{}

	vsf := NewVisualStudioFacade(te, c, clr)
	vsf.Start()
	vsf.Stop()
}
