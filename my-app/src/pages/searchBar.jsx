//import TextField from "@material-ui/core/TextField";
import React from "react";
import "./searchBar.css"
/*import Package from "../package-lock.json"
const searchBar = () => {}
const [searchInput, setSearchInput] = useState("")*/

//Implementation here
function Search({placeholder, data}){
    return (
        <div className="search">
            <div className="searchInputs"></div>
                <input type="text" placeholder={placeholder}/>
                <div className="searchIcon"></div>
            <div className="dataResult"></div>
        </div>
    )
}
export default Search;