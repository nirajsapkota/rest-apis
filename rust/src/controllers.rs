use actix_web::{web, HttpResponse, Result};
use crate::models;
use crate::types;

pub async fn create_book(req: web::Json<types::CreateBookRequest>, app_state: web::Data<models::AppState>) -> Result<HttpResponse> {
    let mut count = app_state.count.lock().unwrap();
    let mut books = app_state.books.lock().unwrap();
    *count += 1;
    
    let book = models::Book{
        id: *count,
        title: req.title.to_string(),
        author: req.author.to_string()
    };
    
    books.push(book);
    Ok(HttpResponse::Ok().json(types::CreateBookResponse {
        message: format!("Successfully created book with id {}.", count)
    }))
}

pub async fn read_books(app_state: web::Data<models::AppState>) -> Result<HttpResponse> {
    let books = app_state.books.lock().unwrap();

    Ok(HttpResponse::Ok().json(types::ReadBooksResponse {
        message: format!("Successfully returned list of books."),
        books: books.to_vec()
    }))
}

pub async fn update_book(req: web::Json<types::UpdateBookRequest>, app_state: web::Data<models::AppState>) -> Result<HttpResponse> {
    let mut books = app_state.books.lock().unwrap();
    
    for book in books.iter_mut() {
        if book.id == req.id {
            match &req.title {
                Some(title) => book.title = title.to_string(),
                None => ()
            }
            
            match &req.author {
                Some(author) => book.author = author.to_string(),
                None => ()
            }
        }
    }
    
    Ok(HttpResponse::Ok().json(types::UpdateBookResponse {
        message: format!("Successfully updated book with id {}.", req.id)
    }))
}

pub async fn delete_book(req: web::Json<types::DeleteBookRequest>, app_state: web::Data<models::AppState>) -> Result<HttpResponse> {
    let mut books = app_state.books.lock().unwrap();
    books.retain(|x| x.id != req.id);
    
    Ok(HttpResponse::Ok().json(types::DeleteBookResponse {
        message: format!("Successfully deleted book with id {}.", req.id)
    }))
}