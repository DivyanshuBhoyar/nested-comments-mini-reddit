import { createContext, useContext, useMemo } from "react";
import { useParams } from "react-router-dom";

import { useAsync } from "../hooks/useAsync.js";
import { fetchSinglePost } from "../services/posts.js";

const PostCtx = createContext();

export const usePost = () => useContext(PostCtx);

export function PostProvider({ children }) {
  const { id } = useParams();
  const {
    loading,
    error,
    value: post,
  } = useAsync(() => fetchSinglePost(id), [id]);

  const commentByParentId = useMemo(() => {
    if (post?.comments == null) return [];
    const group = {};
    post.comments.forEach((comment) => {
      group[comment.parentId] ||= [];
      group[comment.parentId].push(comment);
    });
    return group;
  }, [post?.comments]);

  function getReplies(parentId) {
    return commentByParentId[parentId];
  }

  return (
    <PostCtx.Provider
      value={{
        post: { ...post, id },
        rootComments: commentByParentId[null],
        getReplies,
      }}
    >
      {loading ? (
        <h3>Loading... </h3>
      ) : error ? (
        <h2 className="error-msg">{error}</h2>
      ) : (
        children
      )}
    </PostCtx.Provider>
  );
}
