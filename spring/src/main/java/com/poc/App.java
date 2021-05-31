package com.poc;

import java.util.List;

import javax.validation.Valid;

import com.poc.controller.BookController;
import com.poc.models.Book;
import com.poc.types.CreateBookRequest;
import com.poc.types.CreateBookResponse;
import com.poc.types.DeleteBookRequest;
import com.poc.types.DeleteBookResponse;
import com.poc.types.ReadBooksResponse;
import com.poc.types.Tuple;
import com.poc.types.UpdateBookRequest;
import com.poc.types.UpdateBookResponse;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@SpringBootApplication
@RestController
public class App 
{

    public static void main(String[] args)
    {
        SpringApplication.run(App.class, args);
    }
    
    @PostMapping("/api/v1/book")
    public CreateBookResponse Create(@Valid @RequestBody CreateBookRequest request)
    {
        String message = BookController.CreateBook(request.title, request.author);
        return new CreateBookResponse(message);
    }
    
    @GetMapping("/api/v1/books")
    public ReadBooksResponse Read()
    {
        Tuple<String, List<Book>> response = BookController.ReadBooks();
        return new ReadBooksResponse(response.getKey(), response.getValue());
    }
    
    @PutMapping("/api/v1/book")
    public UpdateBookResponse Update(@Valid @RequestBody UpdateBookRequest request)
    {
        String message = BookController.UpdateBook(request.id, request.title, request.author);
        return new UpdateBookResponse(message);
    }
    
    @DeleteMapping("/api/v1/book")
    public DeleteBookResponse Delete(@Valid @RequestBody DeleteBookRequest request)
    {
        String message = BookController.DeleteBook(request.id);
        return new DeleteBookResponse(message);
    }
    
}
