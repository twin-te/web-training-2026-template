import { eq } from "drizzle-orm";
import { Hono } from "hono";
import { db } from "../../db/client.js";
import { messages } from "../../db/schema.js";
import { MessageRequest } from "../models/message.js";

export const messageRoutes = new Hono();

// ここに追記
