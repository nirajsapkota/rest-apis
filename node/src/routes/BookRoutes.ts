import express, {
    Request,
    Response
} from "express";

import { 
    CreateBook,
    ReadBooks,
    UpdateBook,
    DeleteBook
} from "../controllers/BookController";

const router = express.Router();

router.post("/book", (req: Request, res: Response): void => {
    
    const [status, message] = CreateBook(req.body);
    res.status(status).json({ "message": message });

});

router.get("/books", (req: Request, res: Response): void => {

    const [status, message, books] = ReadBooks();
    res.status(status).json({ "message": message, "books": books });

});

router.put("/book", (req: Request, res: Response): void => {

    const [status, message] = UpdateBook(req.body);
    res.status(status).json({ "message": message });

});

router.delete("/book", (req: Request, res: Response): void => {

    const [status, message] = DeleteBook(req.body);
    res.status(status).json({ "message": message });

});

export default router;