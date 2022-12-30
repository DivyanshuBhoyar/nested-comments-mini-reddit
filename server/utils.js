import { prisma } from "./server.js";

export const getPostList = async () =>
  await prisma.post.findMany({
    select: {
      title: true,
      id: true,
    },
  });
