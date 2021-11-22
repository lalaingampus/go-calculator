
# Calculator

## Table of contents 👀
* [General info](#general-info)
* [Technologies](#technologies)
* [Setup](#setup)


### General info
 Football-API is a Football REST API made to sho some love for the football.

#### The Football-API Object 🍵
| Properties | Description | Type  |
|:----------- |:---------------|:--------|
|num1| for store variabel first | Float64| 
|num2| for store variabel first | Float64 |
|add| add var first with var two | String | 
|sub| substraction var first with var two |String | 
|mul| multiplication var first with var two | String|
|div| divide var first with var two | String|  


#### Routes ⚡
| Routes | HTTP Methods| Description
|:------- |:---------------|:--------------
| /calc     | GET                  | Display result from calculation
| /calc      | POST               | Creates a calculation flow

	
### Technologies
Project is created with:

* Golang
* Vue

### Setup Frontend
To run this project locally, clone repo and add an file in the root:

Then execute in command prompt:
```
$ cd frontend
$ npm install
$ npm run serve
```

### Setup Backend
To run this project locally, clone repo and add an file in the root:

Then execute in command prompt:
```
$ cd backend
$ go mod tidy
$ go run server.go
```

## API Reference

These are the endpoints available from the app

### `GET /calc`

Returns result calculation 

#### Response

<details><summary>Show example response</summary>
<p>

```json
{
  "meta": {
    "code": 200
  },
  "data": [
    {
     "add":40,
     "sub":23,
     "mul":32,
     "div":22,
    }
  ]
}
```

</p>
</details>

---


### `POST /calc`

Creates a new number 1 & number 2 for calculation

#### Request 

This request requires body payload, you can find the example below.

<details><summary>Show example payload</summary>
<p>

```json
{
 "nums1":2,
 "nums2":3,
}
```
</p>
</details>

