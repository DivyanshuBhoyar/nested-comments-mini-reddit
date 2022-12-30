import { prisma } from "./server.js";

const COMMENT_SELECT_FIELDS = {
  id: true,
  message: true,
  parentId: true,
  createdAt: true,
  user: {
    select: {
      id: true,
      name: true,
    },
  },
};

export const getPostList = async () =>
  await prisma.post.findMany({
    select: {
      title: true,
      id: true,
    },
  });

export const getPost = async (id) =>
  await prisma.post.findUnique({
    where: { id: id },
    select: {
      title: true,
      body: true,
      comments: {
        orderBy: {
          createdAt: "desc",
        },
        select: COMMENT_SELECT_FIELDS,
      },
    },
  });
