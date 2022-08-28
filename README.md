# tsnippet

## Simple text snippet fullstack web-app written on Go
![tsnippet](https://user-images.githubusercontent.com/87330747/187065612-27663f19-1bd2-4c3a-95fd-358397b73966.gif)

App created with help of and following along with [ALEX EDWARDS](https://www.alexedwards.net/) ([Github](https://github.com/alexedwards)) book [“Let's Go!” (2nd edition)](https://lets-go.alexedwards.net/)—one of the greatest practical Go books I've read at the moment, which excellent implementation of the approach:

> _“learn by doing”_

hugely helped me grasp the creation of Go web apps.

Frontend part implemented using go templates (html) and CSS.

Description of available functionality:

| Method | Pattern                  | Handler                   | Action                                         |
| :----- | :----------------------- | :------------------------ | :--------------------------------------------- |
| GET    | /                        | home                      | Display the home page                          |
| GET    | /snippet/view/:id        | snippetView               | Display a specific snippet                     |
| GET    | /snippet/create          | snippetCreate             | Display a HTML form for creating a new snippet |
| POST   | /snippet/create          | snippetCreatePost         | Create a new snippet                           |
| GET    | /static/                 | http.FileServer           | Serve a specific static file                   |
| GET    | /user/signup             | userSignup                | Display a HTML form for signing up a new user  |
| POST   | /user/signup             | userSignupPost            | Create a new user                              |
| GET    | /user/login              | userLogin                 | Display a HTML form for logging in a user      |
| POST   | /user/login              | userLoginPost             | Authenticate and login the user                |
| POST   | /user/logout             | userLogoutPost            | Logout the user                                |
| GET    | /about                   | about                     | Display about page                             |
| GET    | /account/view            | accountView               | Display current user account information       |
| GET    | /account/password/update | accountPasswordUpdate     | Display a HTML form for updating the password  |
| POST   | /account/password/update | accountPasswordUpdatePost | Update password                                |

## Installation & Usage

`$ git clone https://github.com/Twofold-One/tsnippet.git`

`$ go run go run ./cmd/web/ -dsn postgres://username:pass@localhost:5432/tsnippet`
