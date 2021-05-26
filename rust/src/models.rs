use std::sync::Mutex;
use serde::{Serialize};

#[derive(Serialize, Clone)]
pub struct Book {
    pub id: i64,
    pub title: String,
    pub author: String
}

pub struct AppState {
    pub count: Mutex<i64>,
    pub books: Mutex<Vec<Book>>
}
