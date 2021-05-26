import express, { Application, Request, Response } from "express";
import cors from "cors";

import {
  BookRoutes
} from "./routes";

const app: Application = express();
const PORT: string = "8080" || process.env.PORT;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(cors());

app.use("/api/v1", BookRoutes);

app.get("/", (req: Request, res: Response) => {
  res.status(200).send("Health Check!");
});

app.listen(PORT, () => {
  console.log(`🚀 Server started at http://localhost:${PORT}!`);
});