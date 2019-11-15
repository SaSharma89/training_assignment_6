package main

// Datastruct to contain primery data
type Datastruct struct {
	PageTitle string
	InURL string
	OutURL string
}

// URLregister to contain input and output URL
type URLregister struct {
	InURL string
	OutURL string
}

//URLregisterList to contain list of URLregister
var  URLregisterList []URLregister

// OutURLBase converted base URL
const (OutURLBase = "http://sandeep-test.com")

// Firsttmpl for /hello page
const ( Firsttmpl = `<h1>{{.PageTitle}}</h1>
<form method="POST" action="/convert">
	<lable>Input URL</lable><input  type="text" name="test" value=""/>
	<input  type="submit" value="submit" />
</form>`)

// Secondtmpl for /convert page
const ( Secondtmpl = `<h1>{{.PageTitle}}</h1>
	<lable>Input URL</lable><input  type="text" name="test1" value={{.InURL}} />
	<lable>Output URL</lable><input  type="text" name="test2" value={{.OutURL}} />
	<form method="GET" action="/list">
		<input  type="submit" value="URLs list" />
	</form>
	<form method="GET" action="/hello">
		<input  type="submit" value="Convert next" />
	</form>`
)

// Thirdtmpl for /list page
const ( Thirdtmpl = `<h1>URL Conversion List</h1>
					<style>
					table, td, th {
					border: 1px solid black;
					}

					table {
					border-collapse: collapse;
					width: 100%;
					}

					th {
					height: 50px;
					}
					</style>
					<table>
						<tr>
						<td>Input URL</td>
						<td>Converted URL</td>
						</tr>
						{{range .}}
						<tr>
						<td>{{.InURL}}</td>
						<td>{{.OutURL}}</td>
						</tr>
						{{end}}
					</table>`)
