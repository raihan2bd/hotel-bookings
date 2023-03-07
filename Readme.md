# Hotel Bookings - using Go (golang)
<p>Hotel Bookings project built using Go and postgresql is likely a web application that allows users to search for available hotels, view room details, and make reservations. It may include features such as user authentication, and integration with third-party APIs for displaying hotel information and availability.</p>

### Tech Stack <a name="tech-stack"></a>

I used Go (golang), PostgreSQL, Bootstrap, Html, Javascript and css to build this project.
  <summary>Full Stack</summary>
  <ul>
    <li>Go</li>
    <li>PostgreSQL</li>
    <li>Bootstrap</li>
    <li>JAVASCRIPT</li>
    <li>Html</li>
    <li>CSS</li>
  </ul>

 <summary>Dependencies</summary>
  <ul>
    <li><a href="https://github.com/go-chi/chi">Chi router</a></li>
    <li><a href="https://github.com/alexedwards/scs/v2">Alex edwards SCS </a> Session Manager</li>
    <li><a href="https://github.com/justinas/nosurf">Nosurf</a> for CSRFToken </li>
    <li><a href="https://github.com/jackc/pgx/v4">pgx</a> Database Driver</li>
    <li><a href="https://github.com/xhit/go-simple-mail/v2">simple mail</a> To create a simple server</li>
    <li><a href="https://github.com/asaskevich/govalidator">Go validator</a> server side form validator</li>
  </ul>

## Demo
![Capture](https://user-images.githubusercontent.com/35267447/223337938-c8ab34f0-20c6-4d8e-b47a-10558d3c7beb.PNG)
![Capture2](https://user-images.githubusercontent.com/35267447/223337986-d66dc7e1-8cc0-4f83-a980-6a9989a48d30.PNG)
![Capture3](https://user-images.githubusercontent.com/35267447/223338009-dc633d11-d320-4579-a8c9-e0e087a415ae.PNG)
![Capture4](https://user-images.githubusercontent.com/35267447/223338037-c9291368-8963-4b16-8704-632aca0e7dbb.PNG)


## 💻 Getting Started
- To get star with this package first of all you have to clone the project ⬇️
``` bash
https://github.com/raihan2bd/hotel-bookings.git
```
- Then Make sure you have install [Go (golang)](https://go.dev/dl/) version 1.8.0 or latest stable version.
- Then make sure you have install [PostgreSQL](https://www.postgresql.org/) on your local mechine if you want to use this project as localy.
- To install all the Go packages navigate the folder address on your terminal and enter the below command ⬇️
``` bash
go mod tidy
```
- After downloading the packages you should rename example.database.yml file name to database.yml file and edit database credentials to your own database information.
![Capture5](https://user-images.githubusercontent.com/35267447/223344475-c64994c5-8a73-44d7-a571-5d3247c8db74.PNG)
- To setup database tables and columns by onClick install [soda cli database migration tool](https://gobuffalo.io/documentation/database/soda/) Then run below command ⬇️
```sh
soda migrate
```


# Usages
> *Note: Before enter the below command make sure you are in the right directory.*

- After finishing the avove instructions you can see the project in your local mechine by entering the below command ⬇️
```bash
./run.bat
```
or 
```sh
go run cmd/web/main.go cmd/web/routes.go cmd/web/middleware.go cmd/web/send-mail.go
```

- Then you can see this project live on your browser by this link http://localhost:8080 or your given the port nuber you set for the project.


## 👥 Author

👤 **Abu Raihan**

- GitHub: [@githubhandle](https://github.com/raihan2bd)
- Twitter: [@twitterhandle](https://twitter.com/raihan2bd)
- LinkedIn: [LinkedIn](https://linkedin.com/in/raihan2bd)

## 🙏 Acknowledgments <a name="acknowledgements"></a>

I would like to thatnks [Trevor Sawler](https://github.com/tsawler) Who help me a lot learn go with this project. 

## ⭐️ Show your support <a name="support"></a>

> Thanks for visiting my repository. Give a ⭐️ if you like this project!

## 📝 License <a name="license"></a>

This project is [MIT](./LICENSE) licensed.

## Contribution
*Your suggestions will be more than appreciated. If you want to suggest anything for this project feel free to do that. :slightly_smiling_face:*
