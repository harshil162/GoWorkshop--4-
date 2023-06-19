import {React, useState} from "react";
import TextField from "@mui/material/TextField";
import "./App.css"
/*import Package from "../package-lock.json"
const searchBar = () => {}
const [searchInput, setSearchInput] = useState("")*/
//Implementation here

function App(){
    const [inputText, setInputText] = useState;
    let inputHandler = (e) => {
        //convert input text to lower case
        var lowerCase = e.target.value.toLowerCase();
        setInputText(lowerCase);
    };
    return (
        <div className="main">
            <div className="search">
            <TextField
                id="outlined-basic"
                onChange={inputHandler}
                variant="outlined"
                fullWidth
                label="Search"/>
            </div>
            input={inputText}
        </div>
    );
}
export default App;
//<List input={inputText}></List>