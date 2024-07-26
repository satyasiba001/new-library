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

books Table:

- book_id: Unique identifier for books.
- name: name of the book.
- author: Author of the book.
- count: Number of copies available.

booktransactions Table:

- borrow_id: Unique identifier for each borrowing transaction.
- member_id: ID of the member borrowing the book.
- book_id: ID of the book being borrowed.
- borrow_date: Date when the book was borrowed.

2. ## Data Storage

The postgres database:
DB Name : newlibrary

- members: Slice of Book structs.
- books: Slice of Member structs.
- booktransactions : Slice of Transaction structs.

## Service Design Pattern

The functionalities are divided into separate services:

Member Services:

- /addMember
- /memberDetails/:member

Book Services:

- /addnewBook
- /booksPresent
- /checkBookAvailability

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
    "member_id":34245,
    "age":89,
    "name":"satya",
    "add_date":"23april"
}
```

#### Response Body (JSON)

```json
{
  "Status": 1,
  "msg": "new member created"
}
```

- responds with the newly created member details including a unique ID.

### 2. API Endpoint: GET - /memberDetails/:member

#### Purpose

Retrieve details of a specific member by their name.

#### Request

- Endpoint: `/memberDetails/:member`
- Method: GET

#### Response Body (JSON)

```json
{
  "Member Details": {
    "member_id": 34245,
    "name": "satya",
    "age": 89,
    "add_date": "23april"
  }
}
```

### 3. API Endpoint: POST - /addnewBook

#### Purpose

Create a new book in the library

#### Request Body (JSON)

```json
{
    "book_id":6906,
    "name":"starwar 4",
    "author":"albert",
    "count": 0
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

### 5. API Endpoint: POST - /checkBookAvailability

#### Purpose

To check if the books is available or not 

#### Request Body (JSON)

```json
{
  
}
```

#### Response Body (JSON)

```json
{

}
```

### 6. API Endpoint: POST - /bookBorrow

#### Purpose

To Borrow Book

#### Request Body (JSON)

```json
{
  
}
```

#### Response Body (JSON)

```json
{

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
  
}
```

#### Response Body (JSON)

```json
{

}
```