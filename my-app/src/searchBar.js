import React, {useState} from 'react'
const searchBar = () => {}
const [searchInput, setSearchInput] = useState("")
//Implementation here
function App(){
    const [inputText, setInputText] = useState;
    let inputHandler = (e) => {
        //convert input text to lower case
        var lowerCase = e.target.value.toLowerCase();
        setInputText(lowerCase);
    };
    return (
        <><div className="search">
            <TextField
                id="outlined-basic"
                onChange={inputHandler}
                variant="outlined"
                fullWidth
                label="Search" />
        </div><List input={inputText} /></>
    );
}
export default App;