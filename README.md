# Bookcrawler 

This application is a webapp for crawl the books from multiple sources
![Architecture of bookcrawler](doc/image.png)
- For details, please check the [#bookcrawler.drawio](#bookcrawler.drawio)

### Usecase this application working for
![Alt text](doc/image-1.png)

### APIs:

| Method | API Url |
|--------| --------|
| GET | /health | 
| POST | /auth | 
| GET | /book/queue | 
| POST | /book/queue | 
| DELETE | /book/queue/:id | 
| POST | /book/download/:id | 

### Todo:

- [ ] Create database 
- [ ] Create webserver
- [ ] Create bot crawler
- [ ] Create userscript for add book
