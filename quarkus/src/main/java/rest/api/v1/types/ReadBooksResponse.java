package rest.api.v1.types;

import java.util.List;

import rest.api.v1.models.Book;

public class ReadBooksResponse {
    private String message;
    private List<Book> books;
    
    public ReadBooksResponse(String message, List<Book> books) {
        this.message = message;
        this.books = books;
    }
    
    public String getMessage() {
        return message;
    }
    
    public void setMessage(String message) {
        this.message = message;
    }
    
    public List<Book> getBooks() {
        return this.books;
    }
    
    public void setBooks(List<Book> books) {
        this.books = books;
    }
}
