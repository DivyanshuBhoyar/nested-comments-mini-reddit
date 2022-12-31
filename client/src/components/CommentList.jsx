import React from "react";
import { useAutoAnimate } from "@formkit/auto-animate/react";
import { Comment } from "./Comment";

export const CommentList = ({ comments }) => {
  const [parent] = useAutoAnimate();
  return comments.map((comment) => (
    <div ref={parent} key={comment.id} className="comment-stack">
      <Comment {...comment} />
    </div>
  ));
};
