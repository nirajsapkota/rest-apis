# Golang Proof of Concept

This project is a simple REST API built using Golang and two modules (Gin, Gorm). It defines a simple CRUD application that mocks a library.

## API Documentation

### Create

**Endpoint**
```
/book
```

**Request Body**
```
{
	"title": string,
	"author": string
}
```

**Response Body**
```
{
    "message": string
}
```

### Read

**Endpoint**
```
/books
```

**Response Body**
```
{
    "books": Book[]
}
```

### Update

**Endpoint**
```
/book
```
  
**Request Body**
```
{
    "id": int,
    "title"?: string,
    "author"?: string
}
```

**Response Body**
```
{
    "message": string
}
```

### Delete

**Endpoint**
```
/book
```

**Request Body**
```
{
    "id": int
}
```

**Response Body**
```
{
    "message": string
}
```

### Models

**Book**
```
{
    "id": int,
    "title": string,
    "author": string
}
```

## Conclusion