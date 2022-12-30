import fastify from "fastify";
import dotenv from "dotenv";
import { PrismaClient } from "@prisma/client";
import sensible from "@fastify/sensible";
import cors from "@fastify/cors";

import { getPostList, getPost } from "./utils.js";
dotenv.config();

const app = fastify({ logger: false });
app.register(sensible);
app.register(cors, {
  credentials: true,
  origin: process.env.CLIENT_URL,
});

export const prisma = new PrismaClient();

app.get("/posts", async (req, res) => {
  try {
    const posts = await getPostList();
    res.send(posts);
  } catch (e) {
    return app.httpErrors.internalServerError(e.message);
  }
});

app.get("/posts/:id", async (req, res) => {
  try {
    const post = await getPost(req.params.id);
    console.log(post);
    res.send(post);
  } catch (e) {
    return app.httpErrors.internalServerError(e.message);
  }
});

const start = async () => {
  try {
    await app.listen({ port: process.env.PORT });
  } catch (err) {
    app.log.error(err);
    process.exit(1);
  }
};
start();
