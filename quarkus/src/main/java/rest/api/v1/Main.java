package rest.api.v1;

import java.util.List;

import javax.validation.Valid;
import javax.ws.rs.Consumes;
import javax.ws.rs.DELETE;
import javax.ws.rs.GET;
import javax.ws.rs.POST;
import javax.ws.rs.PUT;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;

import rest.api.v1.controller.BookController;
import rest.api.v1.models.Book;
import rest.api.v1.types.CreateBookRequest;
import rest.api.v1.types.CreateBookResponse;
import rest.api.v1.types.DeleteBookRequest;
import rest.api.v1.types.DeleteBookResponse;
import rest.api.v1.types.ReadBooksResponse;
import rest.api.v1.types.Tuple;
import rest.api.v1.types.UpdateBookRequest;
import rest.api.v1.types.UpdateBookResponse;

@Path("/api/v1")
@Produces(MediaType.APPLICATION_JSON)
@Consumes(MediaType.APPLICATION_JSON)
public class Main {

    @POST
    @Path("/book")
    public CreateBookResponse Create(@Valid CreateBookRequest request) {
        String message = BookController.CreateBook(request.title, request.author);
        return new CreateBookResponse(message);
    }

    @GET
    @Path("/books")
    public ReadBooksResponse Read() {
        Tuple<String, List<Book>> response = BookController.ReadBooks();
        return new ReadBooksResponse(response.getKey(), response.getValue());
    }
    
    @PUT
    @Path("/book")
    public UpdateBookResponse Update(@Valid UpdateBookRequest request) {
        String message = BookController.UpdateBook(request.id, request.title, request.author);
        return new UpdateBookResponse(message);
    }
    
    @DELETE
    @Path("/book")
    public DeleteBookResponse Delete(@Valid DeleteBookRequest request) {
        String message = BookController.DeleteBook(request.id);
        return new DeleteBookResponse(message);
    }
    
}