# Library Management

Develop a library management system backend in Go that manages details of members, details of books, and borrowing transactions

- System allowing create new member, add new book in library, borrow book from library, etc
- System uses a external databse (postgres here).


1. ## Data Structures

members Table:

- member_id: Unique identifier for members.
- name: Name of the member.
- age: Age of the member.
- add_date: Registration date of the member.
- aadhar: aadhar card of the member

books Table:

- book_id: Unique identifier for books.
- name: name of the book.
- author: Author of the book.
- count: Number of copies available.

booktransaction Table:

- borrow_id: Unique identifier for each borrowing transaction.
- member_id: ID of the member borrowing the book.
- book_name: Name of the book being borrowed.
- borrow_date: Date when the book was borrowed.
- return_date: Date of book return to library.

2. ## Data Storage

The postgres database:
DB Name : newlibrary

- members
- books
- booktransaction

## Service Design Pattern

The functionalities are divided into separate services:

Member Services:

- /addMember
- /memberDetails/:member

Book Services:

- /addnewBook
- /booksPresent

Book Transaction Services:

- /bookBorrow
- /bookReturn
# APIs

### 1. API Endpoint: POST - /addMember

#### Purpose

Create a new member in the system.

#### Request Body (JSON)

```json
{
    "name":"sonali",
    "age":15,
    "aadhar": "1129456000887652"
}
```

#### Response Body (JSON)

```json
{
  "Status": 1,
  "User_ID": "f9b014cb-3846-41af-a7a1-98d731b5cc8a",
  "msg": "You got the admission, Remember your user_ID for future reference"
}
```

- responds with the newly created member details including a unique ID.

### 2. API Endpoint: GET - /memberDetails/:member

#### Purpose

Retrieve details of a specific member by their name.

#### Request

- Endpoint: `/memberDetails/:member_id`
- Method: GET

#### Response Body (JSON)

```json
{
  "Member Details": {
    "name": "hari",
    "age": 15,
    "aadhar": "1129456000800982"
  }
}
```

### 3. API Endpoint: POST - /addnewBook

#### Purpose

Create a new book in the library

#### Request Body (JSON)

```json
{
    "name":"The Alchemist",
    "author":"Paulo Coelho",
    "count": 6
}
```

#### Response Body (JSON)

```json
{
  "Status": 1,
  "msg": "new book arrived to the library"
}

```


### 4. API Endpoint: GET  - /booksPresent

#### Purpose

Retrieve a list of books present

#### Request

- Endpoint: `/booksPresent`
- Method: GET


#### Response Body (JSON)

```json
{
  "All Bools Present": [
    "sherlock holmes",
    "sherlock holmes 2",
    "starwar",
    "starwar 3",
    "3 idiot"
  ]
}

```


### 5. API Endpoint: POST - /bookBorrow

#### Purpose

To Borrow Book

#### Request Body (JSON)

```json
{
    "member_id":"6d1361bc-1493-4b02-bb37-0eee762bf918",
    "name":"The Hobbit"
}
```

#### Response Body (JSON)

```json
{
  "Book borrow ID": "cr01iu2cifct4e25ovc0",
  "status": "You got the book, return it within 10 days"
}
```

### 7. API Endpoint: POST - /bookReturn

#### Purpose

To Return the borrowed Book

#### Request

- Endpoint: `/bookReturn`

#### Request Body (JSON)

```json
{
  "borrow_id":"cr01iu2cifct4e25ovc0"
}
```

#### Response Body (JSON)

```json
{
  "status": "You returned the book"
}
```