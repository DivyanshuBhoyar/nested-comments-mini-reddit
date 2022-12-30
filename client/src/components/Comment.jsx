import React from "react";
import { FaHeart, FaEdit, FaReply, FaTrash } from "react-icons/fa";

import { IconBtn } from "./IconButton";
import { usePost } from "../contexts/PostContext";
import CommentList from "./CommentList";
import { useState } from "react";

const dateFormatter = new Intl.DateTimeFormat(undefined, {
  dateStyle: "medium",
  timeStyle: "short",
});

export default Comment = ({ id, message, user, createdAt }) => {
  const { getReplies } = usePost();
  const childComments = getReplies(id);
  const [areChildrenHidden, setAreChildrenHidden] = useState(false);

  return (
    <>
      <div className="comment">
        <div className="header">
          <span className="name">{user.name}</span>
          <span className="date">
            {dateFormatter.format(Date.parse(createdAt))}
          </span>
        </div>
        <div className="message">{message}</div>
        <div className="footer">
          <IconBtn Icon={FaHeart} aria-label="like">
            2
          </IconBtn>
          <IconBtn Icon={FaReply} aria-label="reply" />
          <IconBtn Icon={FaEdit} aria-label="edit" />
          <IconBtn Icon={FaTrash} aria-label="delete" color="danger" />
        </div>
      </div>
      {childComments?.length && (
        <>
          <div
            className={`nested-comments-stack ${
              areChildrenHidden ? "hide" : ""
            }`}
          >
            <button
              className="collapse-line"
              aria-label="Hide replies"
              onClick={() => setAreChildrenHidden(true)}
            ></button>
            <div className="nested-comments">
              <CommentList comments={childComments} />
            </div>
          </div>
          <button
            className={`btn mt-1 ${!areChildrenHidden ? "hide" : ""}`}
            onClick={() => setAreChildrenHidden(false)}
          >
            Show Replies
          </button>
        </>
      )}
    </>
  );
};
