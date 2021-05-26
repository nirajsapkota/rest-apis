export interface CreateBookRequest {
    title: string,
    author: string
};

export interface UpdateBookRequest {
    id: number,
    title?: string,
    author?: string
};

export interface DeleteBookRequest {
    id: number
};