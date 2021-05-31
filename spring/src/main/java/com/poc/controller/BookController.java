package com.poc.controller;

import java.util.ArrayList;
import java.util.List;

import com.poc.models.Book;
import com.poc.types.Tuple;

public class BookController {
    
    private static int count = 0;
    private static List<Book> books = new ArrayList<Book>();
    
    public static String CreateBook(String title, String author) {
        count++;
        
        Book book = new Book(count, title, author);
        books.add(book);
        
        return "Successfully created book with id " + count + ".";
    }
    
    public static Tuple<String, List<Book>> ReadBooks() {
        return new Tuple<String, List<Book>>("Successfully returned a list of books.", books);
    }
    
    public static String UpdateBook(int id, String title, String author) {
        Book book = books.stream().filter(existingBook -> existingBook.getId() == id).findFirst().orElse(null);
        
        if (book != null && !title.isEmpty()) book.setTitle(title);
        if (book != null && !author.isEmpty()) book.setAuthor(author);
        
        return "Successfully updated book with id " + id + ".";
    }

    public static String DeleteBook(int id) {
        books.removeIf(existingBook -> existingBook.getId() == id);        
        return "Successfully deleted book with id " + id + ".";
    }
    
}
