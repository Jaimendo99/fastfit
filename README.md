<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://i.imgur.com/6wj0hh6.jpg" alt="Project logo"></a>
</p>

<h3 align="center">FastFit</h3>


---

<p align="center"> A Web Development Project for UDLA's Web Development Course
    <br> 
</p>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Deployment](#deployment)
- [Usage](#usage)
- [Built Using](#built_using)
- [TODO](../TODO.md)
- [Contributing](../CONTRIBUTING.md)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)

## 🧐 About <a name = "about"></a>

FastFit is in a nutshell a fasting tracker and a calorie counter. It helps you keep track of your fasting hours and your daily calorie intake.

## 🏁 Getting Started <a name = "getting_started"></a>

This project uses the [GOTTH Stack](https://github.com/arejula27/goth-stack), which is:
- [Go Lang](https://golang.org/) as the backend
- [Echo](https://echo.labstack.com/) as the web framework
- [Templ](https://templ.guide/) as the template engine
- [TailwindCSS](https://tailwindcss.com/) for staying with tailwindUI
- [HTMX](https://htmx.org/) for the frontend magic

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

### Prerequisites

You will need to have 
[Go Lang](https://golang.org/) installed on your machine, follow the instructions on the website.
#### [Air:](https://github.com/cosmtrek/air) hot reload for Go

Air must be in path, you can install it by running the following command:
```bash
go install github.com/cosmtrek/air@latest
```
#### [Templ:](https://templ.guide/) template engine for Go
Templ must also be in PATH, you can install it by running the following command:
```bash
go install github.com/a-h/templ/cmd/templ@latest
```
#### [Echo](https://echo.labstack.com/) web framework for Go
Echo is a go module, you can install it by running the following commands:
```bash
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
go get github.com/labstack/echo-jwt/v4
```


## Auth System
The auth system is based on JWT, you will need to create a .env file in the root of the project with the following variables:
```bash
JWT_SECRET
```
#### Diagram

<img src = "./static/img/jwtFlow.png"/>

### Installing

A step by step series of examples that tell you how to get a development env running.

Say what the step will be

```
Give the example
```

And repeat

```
until finished
```

End with an example of getting some data out of the system or using it for a little demo.



## 🎈 Usage <a name="usage"></a>

Add notes about how to use the system.

## 🚀 Deployment <a name = "deployment"></a>

Add additional notes about how to deploy this on a live system.

## ⛏️ Built Using <a name = "built_using"></a>

- [MongoDB](https://www.mongodb.com/) - Database
- [Express](https://expressjs.com/) - Server Framework
- [VueJs](https://vuejs.org/) - Web Framework
- [NodeJs](https://nodejs.org/en/) - Server Environment

## ✍️ Authors <a name = "authors"></a>

- [@kylelobo](https://github.com/kylelobo) - Idea & Initial work

See also the list of [contributors](https://github.com/kylelobo/The-Documentation-Compendium/contributors) who participated in this project.

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- Hat tip to anyone whose code was used
- Inspiration
- References
