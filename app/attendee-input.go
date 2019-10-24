package app

var InputForm = `<!DOCTYPE html>
<html lang="en">
<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <title>OCOC - Pair Up</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
</head>
<body class="bg-primary">
<div class="container mt-5">
    <div class="row justify-content-center">
        <h2 class="justify-content-center text-light px-5 py-2" style="font-weight: bold;">Morning Schmooze Pair Up</h2>
        <div class="card col-10  bg-light">
            <div class="card-body">
                <div class="mb-5">
                    <a href="/attendees" target="_blank" class="btn btn-primary col-md-4 offset-md-1">Attendee List</a>
                    <a href="/reset-attendees" target="_blank" class="btn btn-primary col-md-4 offset-md-2">Reset</a>
                </div>
                <form method="POST" action="/">
                    <div class="form-group">
                        <label for="name">Name</label>
                        <input type="text" class="form-control form-control-lg {{if .KeyErr}} is-invalid{{else}}{{end}}" name="name" id="name" placeholder="Enter Name (First and Last)" value="{{.Name}}">
                    </div>
                    <div class="form-group">
                        <label for="business">Business Name</label>
                        <input type="text" class="form-control form-control-lg {{if .KeyErr}} is-invalid{{else}}{{end}}" name="business" id="business" placeholder="Enter Name of Business" value="{{.BusinessName}}">
                    </div>
{{/*                    <div class="form-group {{if .MsgErr}} visible {{else}} invisible {{end}}">*/}}
{{/*                        <label class="text-danger">{{.Err}}</label>*/}}
{{/*                    </div>*/}}
                    <div class="form-group">
                        <label for="industry">Industry</label>
                        <select class="form-control form-control-lg" name="industry" id="industry">
                            <option>Select</option>
                            {{ range $value := .Industries }}
                            <option value="{{ $value }}">{{ $value }}</option>
                            {{ end }}
                        </select>
{{/*                        <input type="text" class="form-control form-control-lg {{if .KeyErr}} is-invalid{{else}}{{end}}" name="key" id="key" placeholder="optional custom key" value="{{.Key}}">*/}}
                    </div>
                    <div class="form-group {{if .KeyErr}} visible {{else}} invisible {{end}}">
{{/*                        <label class="text-danger">{{.Err}}</label>*/}}
                    </div>
                    <div class="row justify-content-around form-group mt-5">
                        <input type="submit" class="form-control form-control-lg col-4 btn btn-secondary" name="add" id="add" value="Submit">
                    </div>
                </form>

            </div>


        </div>
    </div>
</div>
</body>
</html>`