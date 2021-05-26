package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/hello/models"
)

type ReadBooksResponse struct {
	Books []models.Book `json:"books" binding:"required"`
}

type CreateBookResponse struct {
	Message string `json:"message" binding:"required"`
}

type UpdateBookResponse struct {
	Message string `json:"message" binding:"required"`
}

type DeleteBookResponse struct {
	Message string `json:"message" binding:"required"`
}

func TestBookController(t *testing.T) {
	ts := httptest.NewServer(setupServer())

	resp := ReadBooks(ts)
	readResponse := ReadBooksResponse{}
	json.NewDecoder(resp.Body).Decode(&readResponse)

	if len(readResponse.Books) != 0 {
		t.Errorf(fmt.Sprintf("Initial read of books should be empty, however there were %d books.\n", len(readResponse.Books)))
		return
	}

	resp = CreateBookA(ts)
	createResponse := CreateBookResponse{}
	json.NewDecoder(resp.Body).Decode(&createResponse)

	expectedMessage := "Created Book with ID 0"
	if createResponse.Message != expectedMessage {
		t.Errorf(fmt.Sprintf("CreateBookA POST response did not match the expected message.\nExpected %s.\nGot %s.\n", expectedMessage, createResponse.Message))
		return
	}

	resp = CreateBookB(ts)
	createResponse = CreateBookResponse{}
	json.NewDecoder(resp.Body).Decode(&createResponse)

	expectedMessage = "Created Book with ID 1"
	if createResponse.Message != expectedMessage {
		t.Errorf(fmt.Sprintf("CreateBookA POST response did not match the expected message.\nExpected %s.\nGot %s.\n", expectedMessage, createResponse.Message))
		return
	}

	resp = ReadBooks(ts)
	readResponse = ReadBooksResponse{}
	json.NewDecoder(resp.Body).Decode(&readResponse)

	if len(readResponse.Books) != 2 {
		t.Errorf(fmt.Sprintf("Expected 2 books, but received %d.\n", len(readResponse.Books)))
		return
	}

	responseBookA := readResponse.Books[0]
	responseBookB := readResponse.Books[1]

	if responseBookA.ID != 0 || responseBookA.Title != bookA.Title || responseBookA.Author != bookA.Author {
		t.Errorf(fmt.Sprintf("ResponseBookA and BookA Mismatch.\n"))
		return
	}

	if responseBookB.ID != 1 || responseBookB.Title != bookB.Title || responseBookB.Author != bookB.Author {
		t.Errorf(fmt.Sprintf("ResponseBookB and BookB Mismatch.\n"))
		return
	}

	resp = UpdateBookA(ts)
	updateResponse := UpdateBookResponse{}
	json.NewDecoder(resp.Body).Decode(&updateResponse)

	expectedMessage = "Updated Book with ID 0"
	if updateResponse.Message != expectedMessage {
		t.Errorf(fmt.Sprintf("UpdateBookA PUT response did not match the expected message.\nExpected %s.\nGot %s.\n", expectedMessage, updateResponse.Message))
		return
	}

	resp = ReadBooks(ts)
	readResponse = ReadBooksResponse{}
	json.NewDecoder(resp.Body).Decode(&readResponse)

	if len(readResponse.Books) != 2 {
		t.Errorf(fmt.Sprintf("Expected 2 books, but received %d.\n", len(readResponse.Books)))
		return
	}

	responseBookA = readResponse.Books[0]
	responseBookB = readResponse.Books[1]

	if responseBookA.ID != 0 || responseBookA.Title != bookAUpdate.Title || responseBookA.Author != bookA.Author {
		t.Errorf(fmt.Sprintf("ResponseBookA and BookA Mismatch.\n"))
		return
	}

	resp = UpdateBookB(ts)
	updateResponse = UpdateBookResponse{}
	json.NewDecoder(resp.Body).Decode(&updateResponse)

	expectedMessage = "Updated Book with ID 1"
	if updateResponse.Message != expectedMessage {
		t.Errorf(fmt.Sprintf("UpdateBookB PUT response did not match the expected message.\nExpected %s.\nGot %s.\n", expectedMessage, updateResponse.Message))
		return
	}

	resp = ReadBooks(ts)
	readResponse = ReadBooksResponse{}
	json.NewDecoder(resp.Body).Decode(&readResponse)

	if len(readResponse.Books) != 2 {
		t.Errorf(fmt.Sprintf("Expected 2 books, but received %d.\n", len(readResponse.Books)))
		return
	}

	responseBookA = readResponse.Books[0]
	responseBookB = readResponse.Books[1]

	if responseBookA.ID != 0 || responseBookA.Title != bookAUpdate.Title || responseBookA.Author != bookA.Author {
		t.Errorf(fmt.Sprintf("ResponseBookA and BookA Mismatch.\n"))
		return
	}

	if responseBookB.ID != 1 || responseBookB.Title != bookB.Title || responseBookB.Author != bookBUpdate.Author {
		t.Errorf(fmt.Sprintf("ResponseBookB and BookB Mismatch.\n"))
		return
	}

	resp = DeleteBookA(ts)
	deleteResponse := DeleteBookResponse{}
	json.NewDecoder(resp.Body).Decode(&deleteResponse)

	expectedMessage = "Deleted Book with ID 0"
	if deleteResponse.Message != expectedMessage {
		t.Errorf(fmt.Sprintf("DeleteBookA DELETE response did not match the expected message.\nExpected %s.\nGot %s.\n", expectedMessage, deleteResponse.Message))
		return
	}

	resp = ReadBooks(ts)
	readResponse = ReadBooksResponse{}
	json.NewDecoder(resp.Body).Decode(&readResponse)

	if len(readResponse.Books) != 1 {
		t.Errorf(fmt.Sprintf("Expected 1 book, but received %d.\n", len(readResponse.Books)))
		return
	}

	responseBookB = readResponse.Books[0]
	if responseBookB.ID != 1 || responseBookB.Title != bookB.Title || responseBookB.Author != bookBUpdate.Author {
		t.Errorf(fmt.Sprintf("ResponseBookB and BookB Mismatch.\n"))
		return
	}

	resp = CreateBookC(ts)
	createResponse = CreateBookResponse{}
	json.NewDecoder(resp.Body).Decode(&createResponse)

	expectedMessage = "Created Book with ID 2"
	if createResponse.Message != expectedMessage {
		t.Errorf(fmt.Sprintf("CreateBookA POST response did not match the expected message.\nExpected %s.\nGot %s.\n", expectedMessage, createResponse.Message))
		return
	}

	resp = ReadBooks(ts)
	readResponse = ReadBooksResponse{}
	json.NewDecoder(resp.Body).Decode(&readResponse)

	if len(readResponse.Books) != 2 {
		t.Errorf(fmt.Sprintf("Expected 2 books, but received %d.\n", len(readResponse.Books)))
		return
	}

	responseBookB = readResponse.Books[0]
	responseBookC := readResponse.Books[1]

	if responseBookB.ID != 1 || responseBookB.Title != bookB.Title || responseBookB.Author != bookBUpdate.Author {
		t.Errorf(fmt.Sprintf("ResponseBookB and BookB Mismatch.\n"))
		return
	}

	if responseBookC.ID != 2 || responseBookC.Title != bookC.Title || responseBookC.Author != bookC.Author {
		t.Errorf(fmt.Sprintf("ResponseBookC and BookC Mismatch.\n"))
		return
	}
}

/**
 * DATA FOR TESTING
 */
var bookA models.CreateBookRequest = models.CreateBookRequest{
	Title:  "Harry Potter and The Philosopher's Stone",
	Author: "J.K. Rowling",
}

var bookB models.CreateBookRequest = models.CreateBookRequest{
	Title:  "Harry Potter and The Chamber of Secrets",
	Author: "J.K. Rowling",
}

var bookC models.CreateBookRequest = models.CreateBookRequest{
	Title:  "Harry Potter and The Prisoner of Azkaban",
	Author: "J.K. Rowling",
}

var bookD models.CreateBookRequest = models.CreateBookRequest{
	Title:  "Harry Potter and The Goblet of Fire",
	Author: "J.K. Rowling",
}

var zero int = 0
var one int = 1

var zeroPtr *int = &zero
var onePtr *int = &one

var bookAUpdate models.UpdateBookRequest = models.UpdateBookRequest{
	ID:    zeroPtr,
	Title: "Harry Potter and The Sorcerer's Stone",
}

var bookBUpdate models.UpdateBookRequest = models.UpdateBookRequest{
	ID:     onePtr,
	Author: "Joanne Kathleen Rowling",
}

var bookADelete models.DeleteBookRequest = models.DeleteBookRequest{
	ID: zeroPtr,
}

/**
 * UTILITY FUNCTIONS FOR TESTING.
 */

func ReadBooks(ts *httptest.Server) *http.Response {
	resp, _ := http.Get(fmt.Sprintf("%s/books", ts.URL))
	return resp
}

func CreateBook(ts *httptest.Server, body []byte) *http.Response {
	resp, _ := http.Post(fmt.Sprintf("%s/book", ts.URL), "application/json", bytes.NewBuffer(body))
	return resp
}

func UpdateBook(ts *httptest.Server, body []byte) *http.Response {
	client := &http.Client{}

	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/book", ts.URL), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, _ := client.Do(req)
	return resp
}

func DeleteBook(ts *httptest.Server, body []byte) *http.Response {
	client := &http.Client{}

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/book", ts.URL), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, _ := client.Do(req)
	return resp
}

func CreateBookA(ts *httptest.Server) *http.Response {
	body, _ := json.Marshal(bookA)
	return CreateBook(ts, body)
}

func CreateBookB(ts *httptest.Server) *http.Response {
	body, _ := json.Marshal(bookB)
	return CreateBook(ts, body)
}

func CreateBookC(ts *httptest.Server) *http.Response {
	body, _ := json.Marshal(bookC)
	return CreateBook(ts, body)
}

func CreateBookD(ts *httptest.Server) *http.Response {
	body, _ := json.Marshal(bookD)
	return CreateBook(ts, body)
}

func UpdateBookA(ts *httptest.Server) *http.Response {
	body, _ := json.Marshal(bookAUpdate)
	return UpdateBook(ts, body)
}

func UpdateBookB(ts *httptest.Server) *http.Response {
	body, _ := json.Marshal(bookBUpdate)
	return UpdateBook(ts, body)
}

func DeleteBookA(ts *httptest.Server) *http.Response {
	body, _ := json.Marshal(bookADelete)
	return DeleteBook(ts, body)
}
