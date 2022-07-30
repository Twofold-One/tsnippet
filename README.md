# tsnippet

## Simple text snippet web-app

| Method | Pattern           | Handler           | Action                                         |
| :----- | :---------------- | :---------------- | :--------------------------------------------- |
| GET    | /                 | home              | Display the home page                          |
| GET    | /snippet/view/:id | snippetView       | Display a specific snippet                     |
| GET    | /snippet/create   | snippetCreate     | Display a HTML form for creating a new snippet |
| POST   | /snippet/create   | snippetCreatePost | Create a new snippet                           |
| GET    | /static/          | http.FileServer   | Serve a specific static file                   |
