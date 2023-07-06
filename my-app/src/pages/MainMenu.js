import React from 'react';
import music from '.././music.png';
import './App.css';
import "./searchBar.js"
 
import { Outlet, Link } from "react-router-dom";

function MainMenu() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={music} className="App-logo" alt="music" />
        <p>
          Welcome to my Music Accessor App.
        </p>
        {/* <a
          className="App-link"
          //href="https://reactjs.org"
          href="localhost:3000"
          target="_blank"
          rel="noopener noreferrer"
        >
          Use This App
        </a> */}
            <Link to="/songmenu">Use This App</Link>
            <Outlet/>
      </header>

    </div>
  );
}

export default MainMenu;