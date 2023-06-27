import React from 'react';
import './App.css';
import {App} from ".././searchBar.js";
const SongMenu = () => {
    return (
        <div className="App" >
        Song Menu
        </div>
    )
    const searchbar = App();
    console.log(searchbar);
}

export default SongMenu;