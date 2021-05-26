use actix_web::{web, App, HttpServer};
use std::sync::Mutex;

mod controllers;
mod models;
mod types;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let app_state = web::Data::new(models::AppState {
        count: Mutex::new(0),
        books: Mutex::new(Vec::new())
    });
    
    HttpServer::new(move || {
        App::new()
            .app_data(app_state.clone())
            .service(
                web::scope("/api/v1")
                    .route("/book", web::post().to(controllers::create_book))
                    .route("/books", web::get().to(controllers::read_books))
                    .route("/book", web::put().to(controllers::update_book))
                    .route("/book", web::delete().to(controllers::delete_book))
            )
    })
    .bind("0.0.0.0:8080")?
    .run()
    .await
}