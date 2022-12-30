import React from "react";
import { useAsync } from "../hooks/useAsync";
import { fetchSinglePost } from "../services/posts";
import { usePost } from "../contexts/PostContext";

import CommentList from "./CommentList";

function Post() {
  const { post, rootComments } = usePost();
  return (
    <>
      <h1>{post.title}</h1>
      <article>{post.body}</article>
      <h3 className="comments-title">Comments</h3>
      <section>
        {rootComments != null && rootComments.length && (
          <div className="mt-4">
            <CommentList comments={rootComments} />
          </div>
        )}
      </section>
    </>
  );
}

export default Post;
