import React from "react";

import { useAsync, useAsyncFnc } from "../hooks/useAsync";
import { fetchSinglePost } from "../services/posts";
import { createComment } from "../services/comment";
import { usePost } from "../contexts/PostContext";
import { CommentList } from "./CommentList";
import { CommentForm } from "./CommentForm";

export function Post() {
  const { post, rootComments } = usePost();
  const { loading, error, exec: createCommentFn } = useAsyncFnc(createComment);

  const onCommentCreate = (message) => {
    return createCommentFn({
      postId: post.id,
      message,
    }).then((comment) => console.log(comment));
  };

  return (
    <>
      <h1>{post.title}</h1>
      <article>{post.body}</article>
      <h3 className="comments-title">Comments</h3>
      <section>
        <CommentForm
          loading={loading}
          error={error}
          onSubmit={onCommentCreate}
          initialValue=""
          autoFocus={false}
        />
        {rootComments != null && rootComments.length && (
          <div className="mt-4">
            <CommentList comments={rootComments} />
          </div>
        )}
      </section>
    </>
  );
}
