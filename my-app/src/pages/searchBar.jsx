//import TextField from "@material-ui/core/TextField";
import React from "react";
import "./searchBar.css"
import SearchIcon from "@material-ui/icons/Search";
//import CloseIcon from "@material"
/*import Package from "../package-lock.json"
const searchBar = () => {}
const [searchInput, setSearchInput] = useState("")*/

//Implementation here
function Search({placeholder, data}){
    return (
        <div className="search">
            <div className="searchInputs"></div>
                <input type="text" placeholder={placeholder}/>
                <div className="SearchIcon">
                    <SearchIcon />
                </div>
            <div className="dataResult">
                {data.map((value, key)=> {
                    return (
                    <a className="dataItem" href={value.link} target="_blank">
                        <p>{value.song}</p>
                    </a>
                    );
                })}
            </div>
        </div>
    )
}
export default Search;