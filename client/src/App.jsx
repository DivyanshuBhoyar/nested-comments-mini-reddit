import { Routes, Route } from "react-router-dom";
import reactLogo from "./assets/react.svg";
import "./App.css";
import { PostList } from "./components/PostList";

function App() {
  return (
    <div className="container">
      <Routes>
        <Route path="/" element={<PostList />} />
        <Route path="/posts/:id" element={null} />
      </Routes>
    </div>
  );
}

export default App;
