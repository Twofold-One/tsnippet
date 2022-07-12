# tsnippet

## Simple text snippet web-app

| Method | Pattern            | Handler         | Action                       |
| :----- | :----------------- | :-------------- | :--------------------------- |
| ANY    | /                  | home            | Display the home page        |
| ANY    | /snippet/view?id=1 | snippetView     | Display a specific snippet   |
| POST   | /snippet/create    | snippetCreate   | Create a new snippet         |
| ANY    | /static/           | http.FileServer | Serve a specific static file |
