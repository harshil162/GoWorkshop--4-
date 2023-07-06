import React from 'react';
import './App.css';
import Search from './searchBar.js'
//import {Search} from "./pages/searchBar.jsx"
import Data from './MusicSheet.json'
const Menu = () => {
    return (
        <form className="new-song-form">
            <div className="App" >
                Song Menu
                <Search placeholder="Enter a Song Name" data={Data}/>
            </div>
            
        </form>
        
    );
};
export default Menu;