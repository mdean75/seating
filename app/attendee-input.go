package app

var InputForm = `<!DOCTYPE html>
<html lang="en">
<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <title>OCOC - Pair Up</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">

    <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js" integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>
</head>
<body class="bg-primary">
<div class="container mt-5">
    <div class="row justify-content-center">
        <h2 class="justify-content-center text-light px-5 py-2" style="font-weight: bold;">Morning Schmooze Pairing App</h2>
        <div class="card col-10  bg-light">
            <div class="card-body">
                <div class="mb-5">
                    <a href="/attendees" target="_blank" class="btn btn-primary col-md-3 mb-3">Attendee List</a>
                    <a href="/seating" class="btn btn-primary col-md-3 offset-md-1 mb-3">Build Pairings ( {{ .ListCount }} )</a>
                    <a href="/reset-attendees" class="btn btn-danger col-md-3 offset-md-1 mb-3">Reset</a>
                </div>

                <div class="alert alert-success alert-dismissible fade show {{if .SuccessMsg}} d-block {{else}} d-none {{end}}" role="alert">
                    <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                        <span aria-hidden="true">&times;</span>
                    </button>
                    {{ .SuccessMsg}}
                </div>

                <form method="POST" action="/">
                    <div class="form-group">
                        <label for="name">Name</label>
                        <input type="text" class="form-control form-control-lg {{if .KeyErr}} is-invalid{{else}}{{end}}" name="name" id="name" placeholder="Enter Name (First and Last)" value="{{.Name}}" required="required">
                    </div>
                    <div class="form-group">
                        <label for="business">Business Name</label>
                        <input type="text" class="form-control form-control-lg {{if .KeyErr}} is-invalid{{else}}{{end}}" name="business" id="business" placeholder="Enter Name of Business" value="{{.BusinessName}}" required="required">
                    </div>
{{/*                    <div class="form-group {{if .MsgErr}} visible {{else}} invisible {{end}}">*/}}
{{/*                        <label class="text-danger">{{.Err}}</label>*/}}
{{/*                    </div>*/}}
                    <div class="form-group">
                        <label for="industry">Industry</label>
                        <select class="form-control form-control-lg" name="industry" id="industry" required="required">
                            <option value="">Select</option>
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

                <div class="row form-group mt-5">
                    <input type="button" class="form-control form-control-lg col-2 btn btn-secondary" name="show" id="show" value="Add Industry" onclick="showDiv();">
                </div>

                <div class="form-group" name="test" id="test" style="display: none">

                    <input type="text" class="form-control form-control-lg" name="newIndustry" id="newIndustry" placeholder="Enter New Industry" >

                    <div class="row form-group mt-5">
                        <input type="button" class="form-control form-control-lg col-2 btn btn-secondary" name="addIndustry" id="addIndustry" value="Add To List" onclick="addIndustry();">
                    </div>
                </div>


            </div>


        </div>
    </div>
</div>
</body>
</html>

<script>
    $(document).ready(function() {
        // show the alert
        setTimeout(function() {
            $(".alert").alert('close');
        }, 3000);
    });

    function showDiv() {
        document.getElementById('test').style.display = "block";
    }

    function addIndustry() {

        let i = document.getElementById('newIndustry');
        let list = document.getElementById('industry');

        console.log("i: ", i.value);
        let add = document.createElement("option");
        add.value = i.value;
        add.text = i.value;

        console.log("add: ", add);

        list.add(add);

        i.value = "";

        document.getElementById('test').style.display = "none";
    }
</script>`
