import React from 'react';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import MainMenu from "../pages/MainMenu";
import SongMenu from "../pages/SongMenu";
import Search from "../pages/searchBar.jsx";
import "../pages/App.css";

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route exact path="/" element={<MainMenu />}/>
        <Route path="songmenu" element={<SongMenu />} />
        <Route path="spacebar" element={<Search />} />
      </Routes>
    </BrowserRouter>
  );
}
