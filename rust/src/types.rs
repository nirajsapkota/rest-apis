use serde::{Serialize, Deserialize};
use crate::models;

#[derive(Debug, Serialize, Deserialize)]
pub struct CreateBookRequest {
    pub title: String,
    pub author: String
}

#[derive(Debug, Serialize, Deserialize)]
pub struct UpdateBookRequest {
    pub id: i64,
    pub title: Option<String>,
    pub author: Option<String>
}

#[derive(Debug, Serialize, Deserialize)]
pub struct DeleteBookRequest {
    pub id: i64
}

#[derive(Debug, Serialize, Deserialize)]
pub struct CreateBookResponse {
    pub message: String
}

#[derive(Serialize)]
pub struct ReadBooksResponse {
    pub message: String,
    pub books: Vec<models::Book>
}

#[derive(Debug, Serialize, Deserialize)]
pub struct UpdateBookResponse {
    pub message: String
}

#[derive(Debug, Serialize, Deserialize)]
pub struct DeleteBookResponse {
    pub message: String
}