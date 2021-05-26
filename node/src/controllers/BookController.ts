import {
    Book
} from "../models";

import { 
    CreateBookRequest,
    UpdateBookRequest,
    DeleteBookRequest
} from "../typings/requests";

let count: number = 0;
let books: Book[] = [];

export const CreateBook = (req: CreateBookRequest): [number, string] => {
    count += 1;
    books.push({ "id": count, ...req });
    return [201, `Successfully created book with id ${count}.`];
}

export const ReadBooks = (): [number, string, Book[]] => {
    return [200, `Successfully returned list of books.`, books];
}

export const UpdateBook = (req: UpdateBookRequest): [number, string] => {
    const book: Book | undefined = books.find(book => book.id === req.id);
    if (book && req.title) book.title = req.title;
    if (book && req.author) book.author = req.author;
    
    return [200, `Successfully updated book with id ${req.id}.`];
}

export const DeleteBook = (req: DeleteBookRequest): [number, string] => {
    books = books.filter(book => book.id !== req.id);
    return [200, `Successfully deleted book with id ${req.id}.`];
}